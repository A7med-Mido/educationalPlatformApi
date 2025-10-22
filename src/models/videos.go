package models

import (
	"github.com/A7med-Mido/educationalPlatformApi/src/models/generics"
)

type Videos struct {
	generics.Model
	CourseID     uint      `json:"course_id"`
	VidLink      string    `json:"vidLink"` // e.g., "one-time", "monthly"
	Title        string    `json:"title"` // "active", "expired"
	Body         string    `json:"body"` // "active", "expired"
}