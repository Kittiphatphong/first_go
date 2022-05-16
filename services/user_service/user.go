package user_service

import (
	"clickcash_backend/logs"
	"clickcash_backend/middleware"
	"clickcash_backend/repositories/user_repo"
	"clickcash_backend/security"
	"errors"
	"regexp"
	"time"
)

type UserServices interface {
	CreateUserSvc(user *UserRequest) (*UserResponse, error)
	UserLoginSvc(user *UserLoginRequest) (*UserLoginResponse, error)
	GetUserByIDSvc(ID uint)(*UserResponse, error)
}
type userServices struct {
	userRepo user_repo.UserRepo
}

func (u userServices) GetUserByIDSvc(ID uint) (*UserResponse, error) {
	getUserByID, err := u.userRepo.GetUserByID(ID)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	if getUserByID.ID==0 {
		return nil,errors.New("USER_NOT_FOUND")
	}
	userResponse := UserResponse{
		ID:        getUserByID.ID,
		FullName:  getUserByID.FullName,
		Email:     getUserByID.Email,
		Status:    getUserByID.Status,
		CreatedAt: getUserByID.CreatedAt,
		UpdatedAt: getUserByID.UpdatedAt,
	}
	return &userResponse, nil
}


func (u userServices) UserLoginSvc(user *UserLoginRequest) (*UserLoginResponse, error) {
	Rex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !Rex.MatchString(user.Email) {
		logs.Error("INVALID_YOUR_EMAIL")
		return nil, errors.New("INVALID_YOUR_EMAIL")
	}
	if user.Email == "" {
		logs.Error("INVALID_YOUR_EMAIL_OR_PASSWORD")
		return nil, errors.New("INVALID_YOUR_EMAIL_OR_PASSWORD")
	}
	getUserByEmail, err := u.userRepo.GetUserByEmail(user.Email)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	err = security.VerifyPassword(getUserByEmail.Password, user.Password)
	if err != nil {
		logs.Error("INVALID_YOUR_EMAIL_OR_PASSWORD")
		return nil, errors.New("INVALID_YOUR_EMAIL_OR_PASSWORD")
	}

	newGenerateAccessToken, err := middleware.NewGenerateAccessToken(getUserByEmail.Email)
	if err != nil {
		logs.Error(err)
		return nil, err
	}
	userLoginResponse := UserLoginResponse{
		ID:          getUserByEmail.ID,
		FullName:    getUserByEmail.FullName,
		Email:       getUserByEmail.Email,
		Status:      getUserByEmail.Status,
		CreatedAt:   getUserByEmail.CreatedAt,
		UpdatedAt:   getUserByEmail.UpdatedAt,
		AccessToken: newGenerateAccessToken,
	}
	return &userLoginResponse, nil
}

func (u userServices) CreateUserSvc(user *UserRequest) (*UserResponse, error) {
	Rex := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
	if !Rex.MatchString(user.Email) {
		logs.Error("INVALID_YOUR_EMAIL")
		return nil, errors.New("INVALID_YOUR_EMAIL")
	}
	if user.Email == "" || user.FullName == "" || user.Password == "" {
		return nil, errors.New("ENTER_ALL_FIELD")
	}
	getUserByEmail, _ := u.userRepo.GetUserByEmail(user.Email)
	if getUserByEmail != nil && getUserByEmail.Email == user.Email {
		logs.Error("USER_EXITS_ALREADY")
		return nil, errors.New("USER_EXITS_ALREADY")
	}

	newEncryptPassword, err := security.NewEncryptPassword(user.Password)
	if err != nil {
		return nil, err
	}

	userRepos := user_repo.User{
		FullName:  user.FullName,
		Email:     user.Email,
		Password:  newEncryptPassword,
		Status:    true,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	createUser, err := u.userRepo.CreateUser(&userRepos)
	if err != nil {
		return nil, err
	}
	userResponse := UserResponse{
		ID:        createUser.ID,
		FullName:  createUser.FullName,
		Email:     createUser.Email,
		Status:    createUser.Status,
		CreatedAt: createUser.CreatedAt,
		UpdatedAt: createUser.UpdatedAt,
	}
	return &userResponse, nil
}

func NewUserServices(userRepo user_repo.UserRepo) UserServices {
	return &userServices{
		userRepo: userRepo,
	}
}
