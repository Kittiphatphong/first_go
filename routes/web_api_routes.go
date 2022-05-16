package routes

import (
	"clickcash_backend/controllers"
	"clickcash_backend/middleware"
	"github.com/gofiber/fiber/v2"
)

type webApiRoutes struct {
	userController controllers.UserController
}

func (u webApiRoutes) Install(app *fiber.App) {
	routerAdmin := app.Group("api/v1/admin", func(ctx *fiber.Ctx) error {
		return ctx.Next()
	})
	routerAdmin.Post("/user/create",u.userController.CreateUserCtrl)
	routerAdmin.Post("/user/login",u.userController.UserLoginCtrl)
	routerOauthAdmin := routerAdmin.Group("/oauth/user/get-one",middleware.NewAuthentication, func(ctx *fiber.Ctx) error {
		return  ctx.Next()
	})
	routerOauthAdmin.Get("/:id",u.userController.GetUserCtrl)

}

func NewWebApiRoutes(userController controllers.UserController)  Routes{
	return &webApiRoutes{userController: userController}
}