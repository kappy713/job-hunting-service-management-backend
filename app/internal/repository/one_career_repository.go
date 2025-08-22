package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type OneCareerRepository interface {
	GetOneCareerByUserID(c *gin.Context, userID uuid.UUID) (*entity.OneCareer, error)
	CreateOrUpdateOneCareer(c *gin.Context, oneCareer *entity.OneCareer) (*entity.OneCareer, error)
}

type oneCareerRepository struct {
	db *gorm.DB
}

func NewOneCareerRepository(db *gorm.DB) OneCareerRepository {
	return &oneCareerRepository{db: db}
}

func (r *oneCareerRepository) GetOneCareerByUserID(c *gin.Context, userID uuid.UUID) (*entity.OneCareer, error) {
	var oneCareer entity.OneCareer
	result := r.db.Where("id = ?", userID).First(&oneCareer)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // レコードが見つからない場合はnilを返す
		}
		return nil, result.Error
	}
	return &oneCareer, nil
}

func (r *oneCareerRepository) CreateOrUpdateOneCareer(c *gin.Context, oneCareer *entity.OneCareer) (*entity.OneCareer, error) {
	// idで既存レコードを検索してupsert
	result := r.db.Save(oneCareer)
	if result.Error != nil {
		return nil, result.Error
	}
	return oneCareer, nil
}
