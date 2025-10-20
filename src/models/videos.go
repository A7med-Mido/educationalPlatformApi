package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Videos struct {
	gorm.Model
	ID          uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	CourseID    uuid.UUID `json:"course_id"`
	VidLink      string    `json:"vidLink"` // e.g., "one-time", "monthly"
	Title        string    `json:"title"` // "active", "expired"
	Body         string    `json:"body"` // "active", "expired"
}