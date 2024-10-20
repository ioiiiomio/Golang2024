package models

type Task struct {
	ID          uint   `gorm:"primaryKey" json:"id"`
	Title       string `json:"title" binding:"required"`
	Description string `json:"description,omitempty"`
	Status      string `json:"status" binding:"required"`
}
