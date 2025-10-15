package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Course struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Thumbnail   string    `json:"thumbnail"` // Path to uploaded thumbnail
	IsFree      bool      `json:"is_free"`
	Price       float64   `json:"price"` // If not free
	Content     string    `json:"content"` // JSON or path to course files
	TeacherID   uuid.UUID `json:"teacher_id"`
}