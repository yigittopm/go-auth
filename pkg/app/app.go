package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-auth/pkg/api/user"
	"go-auth/pkg/cache"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Cache  *cache.Client
}

func New() *App {
	return &App{}
}

func (app *App) Initialize(host, port, username, password, dbname string) {
	var err error

	// Database
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, username, password, dbname)
	app.DB, err = gorm.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Successfully connect to the database.")
	// Router
	app.Router = mux.NewRouter()
}

func (app *App) Run(addr string) {
	fmt.Printf("Running on port %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (app *App) SetupRoutes() {
	user.InitUserAPI(app)

}
