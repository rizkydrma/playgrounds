package request


type UserCreateRequest struct {
	Name			string		`json:"name" validate:"required,min=3"`
	Gender		string		`json:"gender" validate:"required,oneof='M' 'F'"`
	Email			string		`json:"email" validate:"required,email"`
	Address		string		`json:"address" validate:"required"`
	Password	string		`json:"password" validate:"required"`
}

type UserUpdateRequest struct {
	Name			string		`json:"name" validate:"required,min=3"`
	Gender		string		`json:"gender" validate:"required,oneof='M' 'F'"`
	Address		string		`json:"address" validate:"required"`
}

type UserUpdateEmailRequest struct {
	Email			string		`json:"email" validate:"required,email"`
}