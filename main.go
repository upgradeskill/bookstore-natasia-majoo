package main

import (
	"bookstore/database"
	"bookstore/server"
	"bookstore/util"
)

func main() {
	// load config from config.env
	util.LoadConfig()
	// setup model and create seeder
	database.SetupModels()
	// call endpoint routing (serve on localhost:8000)
	server.RegisterRoutes()
}
