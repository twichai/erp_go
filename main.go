package main

import (
	"erp/handler"
	"erp/models"
	"erp/repository"
	"erp/service"
	"fmt"
	"log"
	"os"
	"time"

	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

const (
	host     = "localhost"  // or the Docker service name if running in another container
	port     = 5432         // default PostgreSQL port
	user     = "myuser"     // as defined in docker-compose.yml
	password = "mypassword" // as defined in docker-compose.yml
	dbname   = "erp"        // as defined in docker-compose.yml
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold: time.Second, // Slow SQL threshold
			LogLevel:      logger.Info, // Log level
			Colorful:      true,        // Enable color
		},
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
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

	app.Post("/login", userHandler.LoginHandler)

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("SECRET_KEY_JWT"))},
	}))

	app.Get("/me", userHandler.GetMe)

	app.Get("/users/:id", userHandler.GetUserHandler)
	app.Put("/users/:id", userHandler.UpdateUserHandler)
	log.Fatal(app.Listen(":8000"))
}
