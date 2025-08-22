package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type SupporterzRepository interface {
	GetSupporterzByUserID(c *gin.Context, userID uuid.UUID) (*entity.Supporterz, error)
	CreateOrUpdateSupporterz(c *gin.Context, supporterz *entity.Supporterz) (*entity.Supporterz, error)
}

type supporterzRepository struct {
	db *gorm.DB
}

func NewSupporterzRepository(db *gorm.DB) SupporterzRepository {
	return &supporterzRepository{db: db}
}

func (r *supporterzRepository) GetSupporterzByUserID(c *gin.Context, userID uuid.UUID) (*entity.Supporterz, error) {
	var supporterz entity.Supporterz
	result := r.db.Where("id = ?", userID).First(&supporterz)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // レコードが見つからない場合はnilを返す
		}
		return nil, result.Error
	}
	return &supporterz, nil
}

func (r *supporterzRepository) CreateOrUpdateSupporterz(c *gin.Context, supporterz *entity.Supporterz) (*entity.Supporterz, error) {
	// idで既存レコードを検索してupsert
	result := r.db.Save(supporterz)
	if result.Error != nil {
		return nil, result.Error
	}
	return supporterz, nil
}
