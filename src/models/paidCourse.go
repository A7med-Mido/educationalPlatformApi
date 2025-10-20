package models

import (
	"time"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type PaidCourses struct {
	gorm.Model
	ID            uuid.UUID `gorm:"type:uuid;primary_key" json:"id"`
	UserID        uuid.UUID `json:"user_id"` // Student
	CourseID      uuid.UUID `json:"course_id"`
	Amount         string    `json:"amount"` // e.g., "one-time", "monthly"
	PaymentDate   time.Time     `json:"paymentData"` // "active", "expired"
}