package handlers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" //postgres database driver
)

type App struct {
	DB     *gorm.DB
	Router *mux.Router
}

func (app *App) Initialize(DBDriver, DBUser, DBPassword, DBPort, DBHost, DBName string) {
	var err error

	//connect to DB
	DBURL := fmt.Sprintf("host=%s port=%s user=%s dbname=%s sslmode=disable password=%s", DBHost, DBPort, DBUser, DBName, DBPassword)

	app.DB, err = gorm.Open(DBDriver, DBURL)

	if err != nil {
		fmt.Printf("Cannot connect to %s database\n", DBDriver)
		log.Fatalf("Error casuse %v", err)
	} else {
		fmt.Printf("Connected to database %s\n", DBDriver)
	}

	//app.DB.Debug().AutoMigrate(&models.Room{}) //to do : database migration

	app.Router = mux.NewRouter()

}

func (app *App) Run(addr string) {
	fmt.Println("Listen and serve on: ", addr)
	err := http.ListenAndServe(addr, app.Router)
	if err != nil {
		log.Fatalf("Error listen and serve on port %s  \ncause: %v", addr, err)
	}

}
