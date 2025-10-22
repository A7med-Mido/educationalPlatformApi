package models

import (
	"github.com/A7med-Mido/educationalPlatformApi/src/models/generics"
)

type Course struct {
	generics.Model
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Thumbnail      string    `json:"thumbnail"` // Path to uploaded thumbnail
	IsFree         bool      `json:"is_free"`
	Price          float64   `json:"price"` // If not free
	Content        []*PaidCourses   `gorm:"one2many:course_content;" json:"content"` // JSON or path to course files
	TeacherID      uint      `json:"teacher_id"`
}
