package repository

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type UserRepository interface {
	//ユーザーIDとサービスリストを基に既存ユーザーの情報を更新
	UpdateUserServices(c *gin.Context, userID string, services []string) error
	//新しいユーザー情報をデータベースに保存
	CreateUser(c *gin.Context, user *entity.User) error
	UpdateUser(c *gin.Context, user *entity.User) (*entity.User, error)
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

	// servicesフィールドを更新
	user.Services = services

	// ユーザーレコード全体を保存
	result := r.db.Save(&user)
	return result.Error
}

func (r *userRepository) CreateUser(c *gin.Context, user *entity.User) error {
	// GORMのCreateメソッドを使用してユーザーをデータベースに保存
	result := r.db.WithContext(c).Create(user)
	return result.Error
}

func (r *userRepository) UpdateUser(c *gin.Context, user *entity.User) (*entity.User, error) {
	// user_idで既存レコードを更新
	// Gradeが0の場合は除外して更新
	var result *gorm.DB
	if user.Grade == 0 {
		result = r.db.Omit("grade").Save(user)
	} else {
		result = r.db.Save(user)
	}

	if result.Error != nil {
		return nil, result.Error
	}
	return user, nil
}
