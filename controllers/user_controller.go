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

// GetUserCtrl func
// @Summary get one user by id
// @Description get one user by id
// @Tags USER
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Security ApiKeyAuth
// @Security Authorization
// @Success 200 {object} user_service.UserResponse
// @Failure 402 {object} controllers.ErrorResponse
// @Router /api/v1/admin/oauth/user/get-one/{id} [get]
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

// UserLoginCtrl UserLogin func
// @Summary login user
// @Description login user with email and password
// @Tags USER
// @Accept json
// @Produce json
// @Param UserLoginRequest body user_service.UserLoginRequest{} true "UserLoginRequest"
// @Failure 402 {object} controllers.ErrorResponse
// @Router /api/v1/admin/user/login [post]
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

// CreateUserCtrl CreateUser func
// @Summary create user
// @Description create user with fullname, email, password
// @Tags USER
// @Accept json
// @Produce json
// @Param UserRequest body user_service.UserRequest true "UserRequest"
// @Success 201 {object} user_service.UserResponse
// @Failure 402 {object} controllers.ErrorResponse
// @Router /api/v1/admin/user/create [post]
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
