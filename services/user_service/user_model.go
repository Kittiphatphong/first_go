package user_service

import (
	"time"
)

type UserLoginResponse struct {
	ID          uint      `json:"id"`
	FullName    string    `json:"full_name"`
	Email       string    `json:"email"`
	Status      bool      `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	AccessToken string    `json:"access_token"`
}
type UserLoginRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}


type UserResponse struct {
	ID        uint      `json:"id"`
	FullName  string    `json:"full_name"`
	Email     string    `json:"email"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type UserRequest struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

