package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type LogRepository interface {
	GetLogsByUserID(c *gin.Context, userID uuid.UUID) ([]entity.Log, error)
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(d *gorm.DB) LogRepository {
	return &logRepository{db: d}
}

func (r *logRepository) GetLogsByUserID(c *gin.Context, userID uuid.UUID) ([]entity.Log, error) {
	var logs []entity.Log
	result := r.db.Where("user_id = ?", userID).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}
