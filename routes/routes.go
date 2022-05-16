package routes

import (
	"clickcash_backend/controllers"
	"github.com/gofiber/fiber/v2"
)


type Routes interface {
	Install(app *fiber.App)
}
type routes struct {
	rootController controllers.RootController
}


func (r routes) Install(app *fiber.App) {
	app.Get("/:id", r.rootController.GetDataByID)
	app.Get("/", r.rootController.Index)
	app.Post("/", r.rootController.PostData)
	app.Post("/profile/upload", r.rootController.PostFile)
}

func NewRoutes(rootController controllers.RootController) Routes {
	return &routes{rootController: rootController}
}

