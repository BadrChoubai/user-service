package infrastructure

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
	"user-service/internal/misc"
	"user-service/internal/service"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/requestid"
)

const (
	IDLE_TIMEOUT = 5 * time.Second
)

// @title User Service
// @version 1.0
// @description RESTful service for managing users in DB Instance
// @contact.name TEAM_NAME
// @contact.email TEAM_EMAIL
// @license.url LICENSE_URL
// @host localhost:8080
// @BasePath /api/v1
func Run() {
	mariadb, err := DBConnection()
	if err != nil {
		log.Fatal("Database Connection Failure: $s", err)
	}

	app := fiber.New(fiber.Config{
		AppName:     "User Service",
		IdleTimeout: IDLE_TIMEOUT,
	})

	app.Use(cors.New())
	app.Use(logger.New(logger.ConfigDefault))
	app.Use(recover.New())
	app.Use(requestid.New())

	userRepository := service.NewUserRepository(mariadb)

	userService := service.NewUserService(userRepository)

	misc.NewMiscHandler(app.Group("/api"))
	service.NewUserHandler(app.Group("/api/v1/users"), userService)

	go func() {
		if err := app.Listen(":8080"); err != nil {
			log.Panic(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	fmt.Println("Running Cleanup Tasks...")

	ctx, cancel := context.WithTimeout(context.Background(), IDLE_TIMEOUT)
	defer cancel()

	if err := app.Shutdown(); err != nil {
		log.Fatal("Graceful Shutdown Failed ", err)
	}

	<-ctx.Done()
	log.Println("Timeout of 5 Seconds")
	log.Println("Service shut down successfully")
}
