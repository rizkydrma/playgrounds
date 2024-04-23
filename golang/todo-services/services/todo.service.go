package services

import (
	"todo-services/database"
	"todo-services/handlers/http/payload/request"
	"todo-services/handlers/http/payload/response"
	"todo-services/models"
)

type TodoService interface {
	GetAll(userId uint) ([]models.Todo, error)
	GetById(id string, userId uint) (models.Todo, error)
	Create(todoReq request.TodoRequest, userId uint) (response.TodoResponse, error)
	Update(todoUpdateReq request.TodoUpdateRequest, id string, userId uint) (response.TodoResponse, error)
	ToggleStatus(todoToggleStatus request.TodoToggleRequest, id string, userId uint) (response.TodoResponse, error)
	Delete(id string, userId uint) (response.TodoResponse, error)
}

type TodoServiceImplement struct {
}

func NewTodoService() TodoService {
	return &TodoServiceImplement{}
}

func (t *TodoServiceImplement) GetAll(userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	database.DB.DB.Find(&todos, "user_id = ? AND deleted_at = 0", userId)

	return todos, nil
}

func (t *TodoServiceImplement) GetById(id string, userId uint) (models.Todo, error){
	var todo models.Todo
	database.DB.DB.First(&todo, "id = ? AND user_id = ? AND deleted_at = 0", id, userId)

	return todo, nil
}

func (t *TodoServiceImplement) Create(todoReq request.TodoRequest, userId uint) (response.TodoResponse, error) {
	var todo models.Todo

	todo.Title = todoReq.Title
	todo.Description = todoReq.Description
	todo.UserId = userId

	database.DB.DB.Create(&todo)

	todoResponse := response.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Status: todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return todoResponse, nil
}

func (t *TodoServiceImplement) Update(todoUpdateReq request.TodoUpdateRequest, id string, userId uint) (response.TodoResponse, error){

	todo, _ :=	t.GetById(id, userId)

	todo.Title = todoUpdateReq.Title
	todo.Description = todoUpdateReq.Description

	database.DB.DB.Save(&todo)

	todoResponse := response.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Status: todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return todoResponse, nil
}

func (t *TodoServiceImplement) ToggleStatus(todoToggleReq request.TodoToggleRequest, id string, userId uint) (response.TodoResponse, error){
	todo, _ := t.GetById(id, userId)

	todo.Status = todoToggleReq.Status

	database.DB.DB.Save(&todo)

	todoResponse := response.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Status: todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return todoResponse, nil
}

func (t *TodoServiceImplement) Delete(id string, userId uint) (response.TodoResponse, error){
	todo, _ := t.GetById(id, userId)

	todo.DeletedAt = 1

	database.DB.DB.Save(&todo)

	todoResponse := response.TodoResponse{
		ID: todo.ID,
		Title: todo.Title,
		Description: todo.Description,
		Status: todo.Status,
		CreatedAt: todo.CreatedAt,
		UpdatedAt: todo.UpdatedAt,
	}

	return todoResponse, nil
}