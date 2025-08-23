package repository

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type UserRepository interface {
	GetUserByID(c *gin.Context, userID string) (*entity.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(c *gin.Context, userID string) (*entity.User, error) {
	var user entity.User
	result := r.db.First(&user, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
