package models

import (
	"github.com/A7med-Mido/educationalPlatformApi/src/models/generics"
)

type User struct {
	generics.Model
	Email    string    `gorm:"unique" json:"email"`
	Password string    `json:"-"`
	Role     string    `json:"role"` // "teacher" or "student"
}
