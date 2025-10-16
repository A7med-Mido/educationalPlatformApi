package services

import (
	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/src/models"
	"github.com/A7med-Mido/educationalPlatformApi/src/repositories"
)

type CourseService interface {
	Create(course *models.Course, teacherID uuid.UUID) error
	GetAll(teacherID uuid.UUID) ([]models.Course, error)
	Update(id uuid.UUID, updates *models.Course) error
	Delete(id uuid.UUID) error
	// Add more like validation if needed
}

type courseService struct {
	repo repositories.CourseRepo
}

func NewCourseService() CourseService {
	return &courseService{repo: repositories.NewCourseRepo()}
}

func (s *courseService) Create(course *models.Course, teacherID uuid.UUID) error {
	course.ID = uuid.New()
	course.TeacherID = teacherID
	return s.repo.Create(course)
}

func (s *courseService) GetAll(teacherID uuid.UUID) ([]models.Course, error) {
	return s.repo.FindAllByTeacher(teacherID)
}

func (s *courseService) Update(id uuid.UUID, updates *models.Course) error {
	course, err := s.repo.FindByID(id)
	if err != nil {
		return err
	}
	// Apply updates (simple merge; improve with reflection if needed)
	if updates.Title != "" {
		course.Title = updates.Title
	}
	if updates.Description != "" {
		course.Description = updates.Description
	}
	course.IsFree = updates.IsFree
	course.Price = updates.Price
	if updates.Content != "" {
		course.Content = updates.Content
	}
	return s.repo.Update(course)
}

func (s *courseService) Delete(id uuid.UUID) error {
	return s.repo.Delete(id)
}