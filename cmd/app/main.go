package main

import (
	"fmt"
	"github.com/labstack/echo"
	"go-auth/pkg-v2/config"
	"go-auth/pkg-v2/infrastructure/datastore"
	"go-auth/pkg-v2/infrastructure/router"
	"go-auth/pkg-v2/registry"
	"log"
)

func main() {
	config.ReadConfig()

	db := datastore.NewDB()
	db.LogMode(true)
	defer db.Close()

	r := registry.NewRegistry(db)

	e := echo.New()
	e = router.NewRouter(e, r.NewAppController())

	fmt.Println("Server listen at http://localhost" + ":" + config.C.Server.Address)
	if err := e.Start(":" + config.C.Server.Address); err != nil {
		log.Fatalln(err)
	}
}
