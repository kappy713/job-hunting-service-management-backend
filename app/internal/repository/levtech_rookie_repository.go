package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type LevtechRookieRepository interface {
	GetLevtechRookieByUserID(c *gin.Context, userID uuid.UUID) (*entity.LevtechRookie, error)
	CreateOrUpdateLevtechRookie(c *gin.Context, levtechRookie *entity.LevtechRookie) (*entity.LevtechRookie, error)
}

type levtechRookieRepository struct {
	db *gorm.DB
}

func NewLevtechRookieRepository(db *gorm.DB) LevtechRookieRepository {
	return &levtechRookieRepository{db: db}
}

func (r *levtechRookieRepository) GetLevtechRookieByUserID(c *gin.Context, userID uuid.UUID) (*entity.LevtechRookie, error) {
	var levtechRookie entity.LevtechRookie
	result := r.db.Where("id = ?", userID).First(&levtechRookie)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil // レコードが見つからない場合はnilを返す
		}
		return nil, result.Error
	}
	return &levtechRookie, nil
}

func (r *levtechRookieRepository) CreateOrUpdateLevtechRookie(c *gin.Context, levtechRookie *entity.LevtechRookie) (*entity.LevtechRookie, error) {
	// idで既存レコードを検索してupsert
	result := r.db.Save(levtechRookie)
	if result.Error != nil {
		return nil, result.Error
	}
	return levtechRookie, nil
}
