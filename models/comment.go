package models

import "time"

type Comment struct {
	ID        uint `json:"id" gorm:"primaryKey;type:integer"`
	UserID    uint `json:"user_id"`
	PhotoID   uint `json:"photo_id"`
	Photo     Photo
	Message   string    `json:"message" binding:"required" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
