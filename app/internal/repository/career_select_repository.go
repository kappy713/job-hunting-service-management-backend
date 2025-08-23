package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type CareerSelectRepository interface {
	GetCareerSelectByUserID(c *gin.Context, userID uuid.UUID) (*entity.CareerSelect, error)
	CreateOrUpdateCareerSelect(c *gin.Context, careerSelect *entity.CareerSelect) (*entity.CareerSelect, error)
}

type careerSelectRepository struct {
	db *gorm.DB
}

func NewCareerSelectRepository(db *gorm.DB) CareerSelectRepository {
	return &careerSelectRepository{db: db}
}

func (r *careerSelectRepository) GetCareerSelectByUserID(c *gin.Context, userID uuid.UUID) (*entity.CareerSelect, error) {
	var careerSelect entity.CareerSelect
	result := r.db.Where("id = ?", userID).First(&careerSelect)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // レコードが見つからない場合はnilを返す
		}
		return nil, result.Error
	}
	return &careerSelect, nil
}

func (r *careerSelectRepository) CreateOrUpdateCareerSelect(c *gin.Context, careerSelect *entity.CareerSelect) (*entity.CareerSelect, error) {
	// idで既存レコードを検索してupsert
	result := r.db.Save(careerSelect)
	if result.Error != nil {
		return nil, result.Error
	}
	return careerSelect, nil
}
