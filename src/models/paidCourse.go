package models

import (
	"github.com/A7med-Mido/educationalPlatformApi/src/models/generics"

)

type PaidCourses struct {
	generics.Model
	ID            uint `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uint `json:"user_id"` // Student
	CourseID      uint `json:"course_id"`
	Amount         string    `json:"amount"` // e.g., "one-time", "monthly"
}