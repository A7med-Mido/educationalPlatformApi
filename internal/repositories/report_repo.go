package repositories

import (
	"github.com/A7med-Mido/educationalPlatformApi/internal/database"
	"github.com/A7med-Mido/educationalPlatformApi/internal/models"
	"gorm.io/gorm"
)

type ReportRepo interface {
	Create(report *models.Report) error
}

type reportRepo struct {
	db *gorm.DB
}

func NewReportRepo() ReportRepo {
	return &reportRepo{db: database.DB}
}

func (r *reportRepo) Create(report *models.Report) error {
	return r.db.Create(report).Error
}