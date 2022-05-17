package controllers

import (
	"clickcash_backend/logs"
	"clickcash_backend/services/user_service"
	"github.com/gofiber/fiber/v2"
)

type UserController interface {
	CreateUserCtrl(ctx *fiber.Ctx) error
	UserLoginCtrl(ctx *fiber.Ctx) error
	GetUserCtrl(ctx *fiber.Ctx) error
}
type userController struct {
	userServices user_service.UserServices
}
func (u userController) GetUserCtrl(ctx *fiber.Ctx) error {
	paramsInt, err := ctx.ParamsInt("id")
	if err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	getUserByIDSvc, err := u.userServices.GetUserByIDSvc(uint(paramsInt))
	if err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, getUserByIDSvc)
}
func (u userController) UserLoginCtrl(ctx *fiber.Ctx) error {
	userRequest := user_service.UserLoginRequest{}
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	userLoginResponse, err := u.userServices.UserLoginSvc(&userRequest)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewSuccessResponse(ctx, userLoginResponse)

}
func (u userController) CreateUserCtrl(ctx *fiber.Ctx) error {
	userRequest := user_service.UserRequest{}
	err := ctx.BodyParser(&userRequest)
	if err != nil {
		logs.Error(err)
		return NewErrorResponses(ctx, err)
	}
	response, err := u.userServices.CreateUserSvc(&userRequest)
	if err != nil {
		return NewErrorResponses(ctx, err)
	}
	return NewCreateSuccessResponse(ctx, response)
}

func NewUserController(userServices user_service.UserServices) UserController {
	return &userController{
		userServices: userServices,
	}
}
