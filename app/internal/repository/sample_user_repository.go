package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type SampleUserRepository interface {
	GetAllSampleUsers(c *gin.Context) ([]entity.SampleUser, error)
}

type sampleUserRepository struct {
	db *gorm.DB
}

func NewSampleUserRepository(db *gorm.DB) SampleUserRepository {
	return &sampleUserRepository{db: db}
}

func (r *sampleUserRepository) GetAllSampleUsers(c *gin.Context) ([]entity.SampleUser, error) {
	var users []entity.SampleUser
	result := r.db.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	return users, nil
}
