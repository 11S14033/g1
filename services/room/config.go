package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/gorilla/mux"

	roompb "github.com/11s14033/g1/services/room/commons/pb"
	dbMigrate "github.com/11s14033/g1/services/room/database"
	rRPCService "github.com/11s14033/g1/services/room/delivery/grpc"
	rRestService "github.com/11s14033/g1/services/room/delivery/rest"
	utils "github.com/11s14033/g1/services/room/delivery/utils"
	repository "github.com/11s14033/g1/services/room/repository/sql"
	"github.com/11s14033/g1/services/room/usecase"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type RoomService struct {
	DB     *gorm.DB
	Router *mux.Router
}

//Init viper to set and read config from config.yml
func init() {
	viper.SetConfigFile(`services/room/config.yml`)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

}

func initDB() (DBDriver string, DBURL string) {
	//Get DB config from config.yml
	DBHost := viper.GetString(`database.DB_HOST`)
	DBPort := viper.GetString(`database.DB_PORT`)
	DBUser := viper.GetString(`database.DB_USER`)
	DBName := viper.GetString(`database.DB_NAME`)
	DBPassword := viper.GetString(`database.DB_PASSWORD`)
	DBDriver = viper.GetString(`database.DB_DRIVER`)
	DBURL = fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

	return DBDriver, DBURL

}

func (rService *RoomService) ConnectDB() (err error) {
	//Get config from config.yml
	dbURL, dbDriver := initDB()

	db, err := gorm.Open(dbURL, dbDriver)

	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", dbDriver)
		log.Fatalf("Error casuse %v", err)
	} else {
		fmt.Printf("Connected to database %s\n", dbDriver)
	}

	rService.DB = db

	return nil
}

//Configuration GRPC server
func initGRPC(db *gorm.DB, roomUC usecase.RoomUseCase, address string) error {
	roomServer := rRPCService.NewRoomRPCService(roomUC)
	opts := []grpc.ServerOption{}
	svr := grpc.NewServer(opts...)

	roompb.RegisterRoomServiceServer(svr, roomServer)

	//Register reflection for access service using CLI
	reflection.Register(svr)

	//Starting GRPC Server
	fmt.Println("Room GRPC Server started, using CLI mode, on port", address)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen %v ", err)
	}

	return svr.Serve(lis)
}

//Configuration GRPC using GRPC-EcoSystem gateway delivery
func initGRPCGateway(db *gorm.DB, roomUC usecase.RoomUseCase, addressMux string, addressRPCServer string) error {

	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	mux := runtime.NewServeMux()
	dialOptions := []grpc.DialOption{grpc.WithInsecure()}

	err := roompb.RegisterRoomServiceHandlerFromEndpoint(ctx, mux, addressRPCServer, dialOptions)
	if err != nil {
		fmt.Println("Error when register Room GRPC Gateway, cause: ", err)
		return err
	}

	//Starting GRPC Gateway
	fmt.Println("Room GRPC Server started, using GRPC Gateway mode, on :", addressMux)
	lis, err := net.Listen("tcp", addressMux)

	if err != nil {
		log.Fatalf("Failed to listen %v ", err)
	}

	return http.Serve(lis, mux)

}

//Configuration rest delivery
func initRest(roomUC usecase.RoomUseCase, addressMux string, r *mux.Router) error {

	restService := rRestService.NewRoomRPCService(roomUC)

	//room router
	r.HandleFunc("/v1/rooms", utils.SetMiddlewareJSON(restService.GetRooms)).Methods("GET")
	r.HandleFunc("/v1/room/{id}", utils.SetMiddlewareJSON(restService.GetRoomByID)).Methods("GET")
	r.HandleFunc("/v1/room", utils.SetMiddlewareJSON(restService.SaveRoom)).Methods("POST")
	r.HandleFunc("/v1/room", utils.SetMiddlewareJSON(restService.UpdateRoom)).Methods("PUT")
	r.HandleFunc("/v1/room/{id}", utils.SetMiddlewareJSON(restService.DeleteRoom)).Methods("DELETE")

	fmt.Println("Room rest Server started, using rest mode, on :", addressMux)

	err := http.ListenAndServe(addressMux, r)
	if err != nil {
		log.Fatalf("Failed to listen %v ", err)
	}

	return err

}

func (rService *RoomService) StartService() error {
	var (
		err error

		// gRPC server endpoint
		grpcServerEndpoint = flag.Int("grpcAddress", 0, "gRPC server endpoint")

		// Port usage
		port = flag.Int("port", 0, "port usage")

		// Chose delivery mode. There are three mode: cli, grpc gateway, and rest
		mode = flag.String("mode", "", "mode usage")
	)

	flag.Parse()

	//Connect to database
	rService.ConnectDB()

	//Database migration
	dbMigrate.Load(rService.DB)

	//Integrate to repository service
	rRepo := repository.NewGormRepository(rService.DB)

	//Integrate to usecase service
	rUseCase := usecase.NewRoomUseCase(rRepo)

	//initiae mux
	rService.Router = mux.NewRouter()

	//Handling Delivery
	if *mode == "cli" {
		//define grpc server
		address := fmt.Sprint("0.0.0.0", ":", *port)
		err = initGRPC(rService.DB, rUseCase, address)
		if err != nil {
			return err
		}
	}
	if *mode == "grpc_gateway" {
		//define grpc server gateway
		addressMux := fmt.Sprint("0.0.0.0", ":", *port)
		addressRPCServer := fmt.Sprint("0.0.0.0", ":", *grpcServerEndpoint)
		err = initGRPCGateway(rService.DB, rUseCase, addressMux, addressRPCServer)
		if err != nil {
			return err
		}
	}
	if *mode == "rest" {
		//define rest server
		addressMux := fmt.Sprint("0.0.0.0", ":", *port)
		err = initRest(rUseCase, addressMux, rService.Router)
		if err != nil {
			return err
		}
	}

	return nil

}
