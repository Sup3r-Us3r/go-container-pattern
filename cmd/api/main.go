package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Sup3r-Us3r/go-container-pattern/src/handlers"
	"github.com/Sup3r-Us3r/go-container-pattern/src/repositories"
	"github.com/Sup3r-Us3r/go-container-pattern/src/services"
	"github.com/Sup3r-Us3r/go-container-pattern/src/structs"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func setupDatabase() *gorm.DB {
	dialector := postgres.Open(fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
	))

	db, err := gorm.Open(dialector)

	if err != nil {
		log.Fatalf("database not connected | error: %v", err)
	}

	db.AutoMigrate(&structs.Installment{})

	return db
}

func main() {
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("could not load environment variables | error: %v", err)
	}

	server := fiber.New()

	db := setupDatabase()

	repositoryContainer := repositories.GetRepositories(db)
	serviceContainer := services.GetServices(*repositoryContainer)
	handlers.NewHandlerContainer(server, *serviceContainer)

	server.Get("/health", func(c *fiber.Ctx) error {
		return c.Status(http.StatusOK).JSON(
			map[string]string{"message": "pong :)"},
		)
	})

	err = server.Listen(":3333")

	if err != nil {
		log.Fatalf("error on listen server | error: %v", err.Error())
	}
}
