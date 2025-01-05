package main

import (
	"github.com/RosaLeaveMeAlone/gin-crud/src/database"
	"github.com/RosaLeaveMeAlone/gin-crud/src/plugins"
	"github.com/RosaLeaveMeAlone/gin-crud/src/presentation/server"
)

func init() {
	plugins.LoadEnv()
	plugins.DBConection()
	database.RunMigrations()
}

func main() {
	server.StartServer()
}
