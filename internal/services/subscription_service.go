package services

import (
	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/internal/models"
	"github.com/A7med-Mido/educationalPlatformApi/internal/repositories"
)

type SubscriptionService interface {
	Subscribe(studentID, courseID uuid.UUID, isFree bool) (*models.Subscription, error)
	GetEnrollmentCount(courseID uuid.UUID) (int, error)
}

type subscriptionService struct {
	repo repositories.SubscriptionRepo
}

func NewSubscriptionService() SubscriptionService {
	return &subscriptionService{repo: repositories.NewSubscriptionRepo()}
}

func (s *subscriptionService) Subscribe(studentID, courseID uuid.UUID, isFree bool) (*models.Subscription, error) {
	if !isFree {
		// Simulate payment failure or integrate real payment
		return nil, nil // Return nil to indicate payment needed
	}
	sub := &models.Subscription{
		ID:       uuid.New(),
		UserID:   studentID,
		CourseID: courseID,
		Plan:     "free",
		Status:   "active",
	}
	err := s.repo.Create(sub)
	return sub, err
}

func (s *subscriptionService) GetEnrollmentCount(courseID uuid.UUID) (int, error) {
	count, err := s.repo.CountByCourse(courseID)
	return int(count), err
}