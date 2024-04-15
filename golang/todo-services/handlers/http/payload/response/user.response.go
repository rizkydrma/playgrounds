package response

import "todo-services/models"


type UserResponse struct {
	Name				string		`json:"name"`
	Gender			string		`json:"gender" sql:"type:ENUM('M','F')" gorm:"column:gender"`
	Email				string		`json:"email"`
	Address			string		`json:"address"`
	models.BaseModel
}