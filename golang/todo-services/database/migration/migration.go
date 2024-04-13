package migration

import (
	"fmt"
	"log"
	"todo-services/database"
	"todo-services/models"
)

func RunMigration() {
	err := database.DB.DB.AutoMigrate(&models.Todo{})
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Database Migrated")
}