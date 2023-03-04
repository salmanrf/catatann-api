package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/salmanfr/catatann-api/api/middlewares"
	"github.com/salmanfr/catatann-api/api/routes"
	"github.com/salmanfr/catatann-api/pkg/entities"
	"github.com/salmanfr/catatann-api/pkg/note"
	"github.com/salmanfr/catatann-api/pkg/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("error loading .env file")
		
		return
	}

	db, err := ConnectDB()

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
	}

	noteService := note.NewService(db)
	userService := user.NewService(db)

	app := fiber.New()

	app.Use(middlewares.Cors)
	
	// app.Use(cors.New(cors.Config{
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTION"},
	// 	AllowHeaders: []string{"Origin", "Content-Type", "Accept"},
	// }))

	api := app.Group("/api")

	noteRouter := api.Group("/notes")
	userRouter := api.Group("/users")

	routes.NoteRouter(noteRouter, noteService)
	routes.UserRoutes(userRouter, userService)

	app.Listen(":8080")
}

func ConnectDB() (*gorm.DB, error) {
	dsn := "host=localhost user=salmanrf password='philiasophia123' dbname=catatann port=5433 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return nil, err
	}

	log.Println("Connected to the database.")

	db.Logger = logger.Default.LogMode(logger.Info)

	db.AutoMigrate(&entities.Note{}, &entities.User{}, &entities.Token{})

	return db, nil
}
