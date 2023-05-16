package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"go-auth/pkg/api"
	"go-auth/pkg/cache"
	"go-auth/pkg/repository/user"
	"go-auth/pkg/service"
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
	userAPI := InitUserAPI(app)

	app.Router.HandleFunc("/users", userAPI.FindAllUsers()).Methods("GET")
	app.Router.HandleFunc("/users", userAPI.CreateUser()).Methods("POST")
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.FindByID()).Methods("GET")
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PUT")
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.DeleteUser()).Methods("DELETE")
}

// InitUserAPI ..
func InitUserAPI(a *App) api.UserAPI {
	userRepository := user.NewRepository(a.DB)
	userService := service.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService, a.Cache)
	userAPI.Migrate()

	return userAPI
}
