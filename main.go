package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/joho/godotenv"
	"go-auth/pkg/api"
	"go-auth/pkg/cache"
	"go-auth/pkg/model"
	"go-auth/pkg/repository"
	"go-auth/pkg/service"
	"log"
	"net/http"
	"os"
)

type App struct {
	Router *mux.Router
	DB     *gorm.DB
	Cache  *cache.Client
}

func init() {
	// Load values in env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	server := New()

	// Init cache
	server.Cache = cache.New()

	server.Initialize(
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	server.SetupRoutes()
	server.DB.AutoMigrate(&model.Post{}, &model.User{})

	server.Run(":3001")
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
	postAPI := InitPostAPI(app)

	app.Router.HandleFunc("/posts", postAPI.FindAll()).Methods(http.MethodGet)
	app.Router.HandleFunc("/posts", postAPI.CreatePost()).Methods(http.MethodPost)

	app.Router.HandleFunc("/users", userAPI.FindAll()).Methods(http.MethodGet)
	app.Router.HandleFunc("/users", userAPI.CreateUser()).Methods(http.MethodPost)
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.FindByID()).Methods(http.MethodGet)
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.UpdateUser()).Methods(http.MethodPut)
	app.Router.HandleFunc("/users/{id:[0-9]+}", userAPI.DeleteUser()).Methods(http.MethodDelete)
}

func InitPostAPI(app *App) api.PostAPI {
	postRepository := repository.NewPostRepository(app.DB)
	postService := service.NewPostService(postRepository)
	postAPI := api.NewPostAPI(postService, app.Cache)
	//postAPI.Migrate()

	return postAPI
}

func InitUserAPI(app *App) api.UserAPI {
	userRepository := repository.NewUserRepository(app.DB)
	userService := service.NewUserService(userRepository)
	userAPI := api.NewUserAPI(userService, app.Cache)
	//userAPI.Migrate()

	return userAPI
}
