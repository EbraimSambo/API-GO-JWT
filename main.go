package main

import (
	"api/initializers"
	"api/routes"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectDatabase()
	initializers.SyncDatabase()
}

func main() {
	app := routes.SetupRouter()
	app.Run()
}
