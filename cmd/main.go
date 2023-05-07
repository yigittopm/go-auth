package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"log"
	"net/http"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
}

func main() {
	app := App{}

	app.initialize(
		"localhost",
		"5432",
		"postgres",
		"postgres",
		"analytic",
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

	// Router
	app.Router = mux.NewRouter()
}

func (app *App) run(addr string) {
	fmt.Printf("Running on port %s", addr)
	log.Fatal(http.ListenAndServe(addr, app.Router))
}

func (a *App) routes() {
	userAPI := InitUserAPI(a.DB)
	a.Router.HandleFunc("/users", userAPI.FindAllUsers()).Methods("GET")
	a.Router.HandleFunc("/users", userAPI.CreateUser()).Methods("POST")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.FindByID()).Methods("GET")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods("PUT")
	a.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.DeleteUser()).Methods("DELETE")
}

// InitUserAPI ..
func InitUserAPI(db *gorm.DB) api.UserAPI {
	userRepository := user.NewRepository(db)
	userService := service.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService)
	userAPI.Migrate()

	return userAPI
}
