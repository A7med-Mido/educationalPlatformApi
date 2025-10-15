package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Report struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	Type      string    `json:"type"` // e.g., "enrollments", "revenue"
	Data      string    `json:"data"` // JSON serialized report data
	TeacherID uuid.UUID `json:"teacher_id"`
}