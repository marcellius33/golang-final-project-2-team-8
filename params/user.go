package params

import (
	"mygram/models"
	"time"
)

type UserRegisterResponse struct {
	ID        uint       `json:"id"  example:"1"`
	Username  string     `json:"username"  example:"jhondoe"`
	Email     string     `json:"email" example:"test@example.com"`
	Age       uint        `json:"age" example:"23"`
}

type UserRegisterRequest struct {
	Age      uint   `json:"age" binding:"required,number,min=8"`
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
	Username string `json:"username" binding:"required"`
}

type UserLoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type UserUpdateResponse struct {
	ID        uint       `json:"id"  example:"1"`
	Username  string     `json:"username"  example:"jhondoe"`
	Email     string     `json:"email" example:"test@example.com"`
	Age       uint        `json:"age" example:"23"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserUpdateRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Username string `json:"username" binding:"required"`
}

type RequestLogin struct {
	Email    string `json:"email" example:"test@example.com"`
	Password string `json:"password" example:"password"`
}

type ResponseLogin struct {
	Token string `json:"token" example:"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJxd2Vxd2..."`
}

func ParseToCreateUserResponse(user *models.User) UserRegisterResponse {
	return UserRegisterResponse{
		ID		 : user.ID,
		Age      : user.Age,
		Username : user.Username,
		Email    : user.Email,
	}
}

func ParseToUpdateUserResponse(user *models.User) UserUpdateResponse {
	return UserUpdateResponse{
		ID		  : user.ID,
		Age       : user.Age,
		Username  : user.Username,
		Email     : user.Email,
		UpdatedAt : user.UpdatedAt,
	}
}