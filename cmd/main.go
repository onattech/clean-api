package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/onattech/invest/config"
	"github.com/onattech/invest/routes"
)

func init() {
	log.SetFlags(log.Ltime | log.Lshortfile)
}

func main() {

	app := config.App()

	env := app.Env

	db := app.DB
	defer app.CloseDBConnection()

	timeout := time.Duration(env.ContextTimeout) * time.Second

	router := gin.Default()

	routes.RegisterRoutes(env, timeout, db, router)

	router.Run(env.ServerAddress)
}
