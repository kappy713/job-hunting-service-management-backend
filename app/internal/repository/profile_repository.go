package repository

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type ProfileRepository interface {
	GetProfileByUserID(c *gin.Context, userID uuid.UUID) (*entity.Profile, error)
	CreateOrUpdateProfile(c *gin.Context, profile *entity.Profile) (*entity.Profile, error)
}

type profileRepository struct {
	db *gorm.DB
}

// 新しいProfileRepositoryのインスタンスを作成
func NewProfileRepository(db *gorm.DB) ProfileRepository {
	return &profileRepository{db: db}
}

// 指定されたユーザーIDに紐づくProfileを取得
func (r *profileRepository) GetProfileByUserID(c *gin.Context, userID uuid.UUID) (*entity.Profile, error) {
	var profile entity.Profile
	result := r.db.Where("id = ?", userID).First(&profile)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, result.Error
	}
	return &profile, nil
}

// Profileレコードを作成または更新
func (r *profileRepository) CreateOrUpdateProfile(c *gin.Context, profile *entity.Profile) (*entity.Profile, error) {
	result := r.db.Save(profile)
	if result.Error != nil {
		return nil, result.Error
	}
	return profile, nil
}
