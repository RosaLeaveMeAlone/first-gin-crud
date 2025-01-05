package database

import (
	"log"

	"github.com/RosaLeaveMeAlone/gin-crud/src/database/models"
	"github.com/RosaLeaveMeAlone/gin-crud/src/plugins"
)

func init() {
	plugins.LoadEnv()
	plugins.DBConection()
}

func RunMigrations() {
	err := plugins.DB.AutoMigrate(&models.Post{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}
}
