package repositories

import (
	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/src/database"
	"github.com/A7med-Mido/educationalPlatformApi/src/models"
	"gorm.io/gorm"
)

type CourseRepo interface {
	Create(course *models.Course) error
	FindAllByTeacher(teacherID uuid.UUID) ([]models.Course, error)
	FindAll() ([]models.Course, error) // New: Fetch all courses (for students)
	FindByID(id uuid.UUID) (*models.Course, error)
	Update(course *models.Course) error
	Delete(id uuid.UUID) error
}

type courseRepo struct {
	db *gorm.DB
}

func NewCourseRepo() CourseRepo {
	return &courseRepo{db: database.DB}
}

func (r *courseRepo) Create(course *models.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepo) FindAllByTeacher(teacherID uuid.UUID) ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Where("teacher_id = ?", teacherID).Find(&courses).Error
	return courses, err
}

func (r *courseRepo) FindAll() ([]models.Course, error) {
	var courses []models.Course
	err := r.db.Find(&courses).Error
	return courses, err
}

func (r *courseRepo) FindByID(id uuid.UUID) (*models.Course, error) {
	var course models.Course
	err := r.db.First(&course, "id = ?", id).Error
	return &course, err
}

func (r *courseRepo) Update(course *models.Course) error {
	return r.db.Save(course).Error
}

func (r *courseRepo) Delete(id uuid.UUID) error {
	return r.db.Delete(&models.Course{}, "id = ?", id).Error
}