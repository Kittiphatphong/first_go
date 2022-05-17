package main

import (
	"clickcash_backend/config"
	"clickcash_backend/controllers"
	"clickcash_backend/repositories/user_repo"
	"clickcash_backend/routes"
	"clickcash_backend/services/user_service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)
func main() {

	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	postgresConnection, err := config.PostgresConnection()
	if err != nil {
		return
	}
	newUserRepo := user_repo.NewUserRepo(postgresConnection)
	newUserServices := user_service.NewUserServices(newUserRepo)
	newUserController := controllers.NewUserController(newUserServices)
	newUserRoutes := routes.NewWebApiRoutes(newUserController)
	newUserRoutes.Install(app)

	APP_PORT := config.GetEnv("app.port", "8080")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", APP_PORT)))

}
