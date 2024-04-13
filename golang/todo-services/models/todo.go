package models

import (
	"time"
)


type Todo struct {
	ID					uint						`json:"id"`
	Title				string					`json:"title"`
	Description	string					`json:"description"`
	Status			bool						`json:"status" gorm:"type:bool;default:0;not null"`
	CreatedAt 	*time.Time			`json:"created_at"`
	UpdatedAt		*time.Time			`json:"updated_at"`
	DeletedAt		int8						`json:"-" gorm:"column:deleted_at"`	
}