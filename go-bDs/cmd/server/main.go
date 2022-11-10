package main

import (
	"github.com/bootcamp-go/go-bDs/cmd/server/routes"
	"github.com/bootcamp-go/go-bDs/pkg/db"
	"github.com/joho/godotenv"
)

func main() {
	loadEnv()

	webEngine, db := db.ConnectDatabase()
	router := routes.NewRouter(webEngine, db)
	router.MapRoutes()

	webEngine.Run(":8080")
}

func loadEnv() {
	if err := godotenv.Load(); err != nil {
		panic(err)
	}
}
