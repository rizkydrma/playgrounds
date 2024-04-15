package models

type User struct {
	Name			string		`json:"name"`
	Gender		string		`json:"gender" gorm:"column:gender"`
	Email			string		`json:"email"`
	Address		string		`json:"address"`
	Password	string		`json:"-" gorm:"column:password"`
	BaseModel
}