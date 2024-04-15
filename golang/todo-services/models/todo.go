package models

type Todo struct {
	BaseModel
	Title				string					`json:"title"`
	Description	string					`json:"description"`
	Status			bool						`json:"status" gorm:"type:bool;default:0;not null"`
}

// func (u *Todo) BeforeCreate(tx *gorm.DB) (err error) {
// 	u.UpdatedAt = nil
//   return
// }