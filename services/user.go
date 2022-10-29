package services

import (
	"mygram/helpers"
	"mygram/models"
	"mygram/params"
	"mygram/repositories"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(userRegisterRequest params.UserRegisterRequest) (*params.UserRegisterResponse, error)
	Login(userLoginRequest params.UserLoginRequest) (*params.ResponseLogin, error)
	Update(userId uint, userUpdateRequest params.UserUpdateRequest) (*params.UserUpdateResponse, error)
	Delete(userId uint) error
}

type userService struct {
	repository repositories.UserRepository
}

func NewUserService(repository repositories.UserRepository) *userService {
	return &userService{repository: repository}
}

func (u *userService) Register(userRegisterRequest params.UserRegisterRequest) (*params.UserRegisterResponse, error) {
	pwHash, err := bcrypt.GenerateFromPassword([]byte(userRegisterRequest.Password), bcrypt.DefaultCost)
	if err != nil {
		return &params.UserRegisterResponse{}, err
	}

	newUser := models.User{
		Username: userRegisterRequest.Username,
		Email:    userRegisterRequest.Email,
		Password: string(pwHash),
		Age:      userRegisterRequest.Age,
	}

	_, err = u.repository.CreateUser(&newUser)

	if err != nil {
		return &params.UserRegisterResponse{}, err
	}
	resp := params.ParseToCreateUserResponse(&newUser)

	return &resp, nil	
}

func (u *userService) Login(userLoginRequest params.UserLoginRequest) (*params.ResponseLogin, error) {
	userFound, err := u.repository.FindUserByEmail(userLoginRequest.Email)
	if err != nil {
		return &params.ResponseLogin{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(userFound.Password), []byte(userLoginRequest.Password))
	if err != nil {
		return &params.ResponseLogin{}, err
	}

	token := helpers.GenerateToken(userFound.ID, userFound.Email)

	resp := params.ResponseLogin{}
	resp.Token = token

	return &resp, nil
}

func (u *userService) Update(userId uint, userUpdateRequest params.UserUpdateRequest) (*params.UserUpdateResponse, error) {
	userModel, err := u.repository.FindUserByID(userId)
	if err != nil {
		return &params.UserUpdateResponse{}, err
	}

	userModel.Email = userUpdateRequest.Email
	userModel.Username = userUpdateRequest.Username
	userModel.UpdatedAt = time.Now()
	user, err := u.repository.UpdateUser(userId, userModel)
	if err != nil {
		return &params.UserUpdateResponse{}, err
	}

	resp := params.ParseToUpdateUserResponse(user)
	return &resp, err
}

func (u *userService) Delete(userId uint) error {
	return u.repository.DeleteUser(userId)
}