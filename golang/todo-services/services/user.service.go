package services

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/models"
)

type UserService interface {
	GetAll() ([]models.User, error)
	GetById(id string) (models.User, error)
	CheckEmail(email string) (models.User, error)
	Update(userReq request.UserUpdateRequest, id string) (response.UserResponse, error)
	UpdateEmail(userReq request.UserUpdateEmailRequest, id string) (response.UserResponse, error)
	Delete(id string) (response.UserResponse, error)
}

type UserServiceImplement struct{}

func NewUserService() UserService {
	return &UserServiceImplement{}
}


func (u *UserServiceImplement) GetAll() ([]models.User, error) {
	var users []models.User
	database.DB.DB.Find(&users, "deleted_at = 0")

	return users, nil
}

func (u *UserServiceImplement) GetById(id string) (models.User, error) {
	var user models.User
	database.DB.DB.First(&user, "id = ? AND deleted_at = 0", id)

	return user, nil
}

func (u *UserServiceImplement) CheckEmail(email string) (models.User, error){
	var user models.User
	database.DB.DB.First(&user, "email = ? AND deleted_at = 0", email)

	return user, nil
}

func (u *UserServiceImplement) Update(userReq request.UserUpdateRequest, id string) (response.UserResponse, error) {
	user, _ := u.GetById(id)

	user.Name 		= userReq.Name
	user.Address	= userReq.Address
	user.Gender		= userReq.Gender

	database.DB.DB.Save(&user)

	userResponse := response.UserResponse{
		Name: user.Name,
		Gender: user.Gender,
		Email: user.Email,
		Address: user.Address,
	}

	return userResponse, nil
}

func (u *UserServiceImplement) UpdateEmail(userReq request.UserUpdateEmailRequest, id string) (response.UserResponse, error) {
	user, _ := u.GetById(id)

	user.Email = userReq.Email

	database.DB.DB.Save(&user)

	userResponse := response.UserResponse{
		Name: user.Name,
		Gender: user.Gender,
		Email: user.Email,
		Address: user.Address,
	}

	return userResponse, nil
}

func (u *UserServiceImplement) Delete(id string) (response.UserResponse, error) {
	user, _ := u.GetById(id)

	user.DeletedAt = 1

	database.DB.DB.Save(&user)

	userResponse := response.UserResponse{
		Name: user.Name,
		Gender: user.Gender,
		Email: user.Email,
		Address: user.Address,
	}

	return userResponse, nil
}