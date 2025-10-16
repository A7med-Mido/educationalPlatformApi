package repositories

import (
	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/src/database"
	"github.com/A7med-Mido/educationalPlatformApi/src/models"
	"gorm.io/gorm"
)

type SubscriptionRepo interface {
	Create(sub *models.Subscription) error
	CountByCourse(courseID uuid.UUID) (int64, error)
}

type subscriptionRepo struct {
	db *gorm.DB
}

func NewSubscriptionRepo() SubscriptionRepo {
	return &subscriptionRepo{db: database.DB}
}

func (r *subscriptionRepo) Create(sub *models.Subscription) error {
	return r.db.Create(sub).Error
}

func (r *subscriptionRepo) CountByCourse(courseID uuid.UUID) (int64, error) {
	var count int64
	err := r.db.Model(&models.Subscription{}).Where("course_id = ?", courseID).Count(&count).Error
	return count, err
}