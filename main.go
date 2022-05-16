package main

import (
	"clickcash_backend/config"
	"clickcash_backend/controllers"
	_ "clickcash_backend/docs"
	"clickcash_backend/repositories/user_repo"
	"clickcash_backend/routes"
	"clickcash_backend/services/user_service"
	"fmt"
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"log"
)

// @title ClickCash API
// @version 2.0
// @description This is a sample server api for ClickCash.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swaager.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @securityDefinitions.apikey Authorization
// @in header
// @name Authorization


// @schemas http
// @schemas https

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

	nwRootController := controllers.NwRootController()
	newRoutes := routes.NewRoutes(nwRootController)
	newRoutes.Install(app)

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	APP_PORT := config.GetEnv("app.port", "8080")
	log.Fatal(app.Listen(fmt.Sprintf(":%s", APP_PORT)))

}
