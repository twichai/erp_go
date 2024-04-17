package main

import (
	"erp/handler"
	"erp/models"
	"erp/repository"
	"erp/service"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "erp"        // as defined in docker-compose.yml
)

func main() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Connect database successful")

	if err = db.AutoMigrate(models.User{}); err != nil {
		panic(err)
	}
	fmt.Println("Migrate successful")

	userRepo := &repository.UserRepositoryGrom{DB: *db}
	userService := &service.UserService{UserRepo: userRepo}
	userHandler := &handler.UserHandler{UserService: *userService}

	app := fiber.New()

	app.Post("/users", userHandler.CreateUserHandler)

	log.Fatal(app.Listen(":8000"))
}
