package main

import (
	"goingEter/infrastructure/database"
	"goingEter/infrastructure/environments"
	"goingEter/routes"
	"os"
)

func init() {
	environments.LoadEnv()
	database.ConnectDB()
}

func main() {
	r := routes.RouterSetup()
	r.Run(os.Getenv("SERVER_PORT"))
}
