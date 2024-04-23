package services

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/models"
)

type AuthService interface {
	Login(loginReq request.LoginRequest)(response.BaseResponse, error)
	Register()(response.BaseResponse, error)
}

type AuthServiceImplement struct {}

func NewAuthService() AuthService{
	return &AuthServiceImplement{}
}

func (s *AuthServiceImplement) Login(loginReq request.LoginRequest)(response.BaseResponse, error){
	var user models.User

	database.DB.DB.First(&user, "email = ? AND deleted_at = 0", loginReq.Email)

	return response.BaseResponse{
		Data: "",
		Message: "",
		Code: "",
	}, nil

}

func (s *AuthServiceImplement) Register()(response.BaseResponse, error){
	
	return response.BaseResponse{
		Data: "",
		Message: "",
		Code: "",
	}, nil
}