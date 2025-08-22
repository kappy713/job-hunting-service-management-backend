package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type MynaviRepository interface {
	GetMynaviByUserID(c *gin.Context, userID uuid.UUID) (*entity.Mynavi, error)
	CreateOrUpdateMynavi(c *gin.Context, mynavi *entity.Mynavi) (*entity.Mynavi, error)
}

type mynaviRepository struct {
	db *gorm.DB
}

func NewMynaviRepository(db *gorm.DB) MynaviRepository {
	return &mynaviRepository{db: db}
}

func (r *mynaviRepository) GetMynaviByUserID(c *gin.Context, userID uuid.UUID) (*entity.Mynavi, error) {
	var mynavi entity.Mynavi
	result := r.db.Where("id = ?", userID).First(&mynavi)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // レコードが見つからない場合はnilを返す
		}
		return nil, result.Error
	}
	return &mynavi, nil
}

func (r *mynaviRepository) CreateOrUpdateMynavi(c *gin.Context, mynavi *entity.Mynavi) (*entity.Mynavi, error) {
	// idで既存レコードを検索してupsert
	result := r.db.Save(mynavi)
	if result.Error != nil {
		return nil, result.Error
	}
	return mynavi, nil
}
