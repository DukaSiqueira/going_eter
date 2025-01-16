package main

import (
	"goingEter/config"
	"goingEter/routes"
	"os"
)

func init() {
	config.LoadEnvs()
	config.ConnectDB()
}

func main() {
	r := routes.RouterSetup()
	r.Run(os.Getenv("SERVER_PORT"))
}
