package database

import (
	"log"
	"os"

	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/src/models"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	dbPath := os.Getenv("DB_PATH")
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to DB")
	}
	DB = db

	// Auto-migrate models
	DB.AutoMigrate(&models.User{}, &models.Course{})

	// Seed teacher if not exists
	seedTeacher()
}

func seedTeacher() {
	var teacher models.User
	err := DB.Where("role = ?", "teacher").First(&teacher).Error
	if err == gorm.ErrRecordNotFound {
		hashedPass, _ := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
		teacher = models.User{
			ID:       uuid.New(),
			Email:    "teacher@example.com",
			Password: string(hashedPass),
			Role:     "teacher",
		}
		if err := DB.Create(&teacher).Error; err != nil {
			log.Println("Failed to seed teacher")
		} else {
			log.Println("Teacher seeded: email=teacher@example.com, password=password")
		}
	}
}