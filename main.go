package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/salmanfr/catatann-api/api/routes"
	"github.com/salmanfr/catatann-api/pkg/entities"
	"github.com/salmanfr/catatann-api/pkg/note"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	db, err := ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	noteService := note.NewService(db)

	app := fiber.New()

	api := app.Group("/api")

	noteRoutes := api.Group("/notes")

	routes.NoteRouter(noteRoutes, noteService)
	
	app.Listen(":8080")
}

func ConnectDB() (*gorm.DB, error) {
	dsn := ""

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	log.Println("Connected.")

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&entities.Note{})

	return db, nil
}