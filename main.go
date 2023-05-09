package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"go-auth/pkg/api"
	"go-auth/pkg/cache"
	"go-auth/pkg/repository/user"
	"go-auth/pkg/service"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	cache  *cache.Client
}

func init() {
	// Load values in env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	app := App{}

	// Init cache
	app.cache = cache.New()

	app.initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	app.routes()

	app.run(":3001")
}

func (app *App) initialize(host, port, username, password, dbname string) {
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

func (app *App) run(addr string) {
	fmt.Printf("Running on port %s \n", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (a *App) routes() {
	userAPI := InitUserAPI(a)
	a.Router.HandleFunc("/users", userAPI.FindAllUsers()).Methods("GET")
	a.Router.HandleFunc("/users", userAPI.CreateUser()).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.FindByID()).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PUT")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.DeleteUser()).Methods("DELETE")
}

// InitUserAPI ..
func InitUserAPI(a *App) api.UserAPI {
	userRepository := user.NewRepository(a.DB)
	userService := service.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService, a.cache)
	userAPI.Migrate()

	return userAPI
}
