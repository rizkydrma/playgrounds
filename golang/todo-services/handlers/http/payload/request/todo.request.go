package request


type TodoRequest struct {
	Title				string					`json:"title" validate:"required"`
	Description	string					`json:"description" validate:"required"`
	Status			bool						`json:"status" gorm:"type:bool;default:0;not null"`
}

type TodoUpdateRequest struct {
	Title				string					`json:"title" validate:"required"`
	Description	string					`json:"description" validate:"required"`
}

type TodoToggleRequest struct {
	Status			bool						`json:"status" validate:"required"`
}