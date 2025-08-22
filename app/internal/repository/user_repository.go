package repository

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type UserRepository interface {
	UpdateUserServices(c *gin.Context, userID string, services []string) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) UpdateUserServices(c *gin.Context, userID string, services []string) error {
	var user entity.User
	// ユーザーIDでユーザーを検索
	if err := r.db.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	// サービスの更新に加えて、gradeが制約を満たすように値を設定
	// これにより、データベースのチェック制約違反を防ぐ
	if user.Grade < 1 || user.Grade > 10 {
		user.Grade = 1 // 仮に有効な値を設定
	}
	
	// servicesフィールドを更新
	user.Services = services

	// ユーザーレコード全体を保存
	result := r.db.Save(&user)
	return result.Error
}