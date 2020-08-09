package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	roompb "github.com/11s14033/g1/services/room/commons/pb"
	rServiceRPC "github.com/11s14033/g1/services/room/delivery/grpc"
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
	DB *gorm.DB
}

//Init viper to set and read config from config.yml
func init() {
	viper.SetConfigFile(`services/room/config.yml`)
	if err := viper.ReadInConfig(); err != nil {
		fmt.Printf("Error reading config file, %s", err)
	}

}

//Configuration GRPC server
func initGRPC(db *gorm.DB, roomUC usecase.RoomUseCase, address string) error {
	roomServer := rServiceRPC.NewRoomRPCService(roomUC)
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

//Configuration GRPC using GRPC-EcoSystem gateway
func initGRPCGatewayRest(db *gorm.DB, roomUC usecase.RoomUseCase, addressMux string, addressRPCServer string) error {

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

func InitializeDB() (DBDriver string, DBURL string) {
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
	dbURL, dbDriver := InitializeDB()

	db, err := gorm.Open(dbURL, dbDriver)

	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", dbDriver)
		log.Fatalf("Error casuse %v", err)
	} else {
		fmt.Printf("Connected to database %s\n", dbDriver)
	}

	rService.DB = db

	//app.DB.Debug().AutoMigrate(&models.Room{}) //to do : database migration

	return nil
}

func (rService *RoomService) StartService() error {
	var (
		err error

		// gRPC server endpoint
		grpcServerEndpoint = flag.Int("grpcAddress", 0, "gRPC server endpoint")

		// Port usage
		port = flag.Int("port", 0, "port usage")

		// Chose delivery mode. There are three mode: cli, grpc gateway, rest using http
		mode = flag.String("mode", "", "mode usage")
	)

	flag.Parse()

	//Connect to database
	rService.ConnectDB()

	//Integrate to repository service
	rRepo := repository.NewGormRepository(rService.DB)

	//Integrate to usecase service
	rUseCase := usecase.NewRoomUseCase(rRepo)

	//Handling Delivery
	if *mode == "cli" {

		address := fmt.Sprint("0.0.0.0", ":", *port)

		//define rpc server
		err = initGRPC(rService.DB, rUseCase, address)
		if err != nil {
			return err
		}
	}
	if *mode == "grpc_gateway" {

		addressMux := fmt.Sprint("0.0.0.0", ":", *port)
		addressRPCServer := fmt.Sprint("0.0.0.0", ":", *grpcServerEndpoint)
		err = initGRPCGatewayRest(rService.DB, rUseCase, addressMux, addressRPCServer)
		if err != nil {
			return err
		}
	}
	if *mode == "rest" {
		fmt.Println("To do handle rest http")
	}

	return nil

}
