package services

import (
	"encoding/json"
	"github.com/google/uuid"
	"github.com/A7med-Mido/educationalPlatformApi/internal/models"
	"github.com/A7med-Mido/educationalPlatformApi/internal/repositories"
)

type ReportService interface {
	GenerateEnrollments(teacherID uuid.UUID) (*models.Report, error)
}

type reportService struct {
	courseRepo       repositories.CourseRepo
	subscriptionRepo repositories.SubscriptionRepo
	reportRepo       repositories.ReportRepo
}

func NewReportService() ReportService {
	return &reportService{
		courseRepo:       repositories.NewCourseRepo(),
		subscriptionRepo: repositories.NewSubscriptionRepo(),
		reportRepo:       repositories.NewReportRepo(),
	}
}

func (s *reportService) GenerateEnrollments(teacherID uuid.UUID) (*models.Report, error) {
	courses, err := s.courseRepo.FindAllByTeacher(teacherID)
	if err != nil {
		return nil, err
	}

	reportData := make(map[string]int)
	for _, course := range courses {
		count, err := s.subscriptionRepo.CountByCourse(course.ID)
		if err == nil {
			reportData[course.Title] = int(count)
		}
	}

	dataJSON, _ := json.Marshal(reportData)
	report := &models.Report{
		ID:        uuid.New(),
		Type:      "enrollments",
		Data:      string(dataJSON),
		TeacherID: teacherID,
	}
	err = s.reportRepo.Create(report)
	return report, err
}