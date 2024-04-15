package models

import "time"


type BaseModel struct {
	ID					uint						`json:"id"`
	CreatedAt 	*time.Time			`json:"created_at"`
	UpdatedAt		*time.Time			`json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	DeletedAt		int8						`json:"-" gorm:"column:deleted_at"`	
}