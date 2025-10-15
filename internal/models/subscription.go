package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Subscription struct {
	gorm.Model
	ID        uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID    uuid.UUID `json:"user_id"` // Student
	CourseID  uuid.UUID `json:"course_id"`
	Plan      string    `json:"plan"` // e.g., "one-time", "monthly"
	Status    string    `json:"status"` // "active", "expired"
}