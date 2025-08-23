package repository

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

type UserRepository interface {
	GetUserByID(c *gin.Context, userID string) (*entity.User, error)
	//ユーザーIDとサービスリストを基に既存ユーザーの情報を更新
	UpdateUserServices(c *gin.Context, userID string, services []string) error
	//新しいユーザー情報をデータベースに保存
	CreateUser(c *gin.Context, user *entity.User) error
	UpdateUser(c *gin.Context, userID string, updateData map[string]interface{}) (*entity.User, error)
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

func (r *userRepository) UpdateUser(c *gin.Context, userID string, updateData map[string]interface{}) (*entity.User, error) {
	// user_idで既存レコードを検索
	var existingUser entity.User
	result := r.db.Where("user_id = ?", userID).First(&existingUser)
	if result.Error != nil {
		// レコードが見つからない場合はエラーを返す
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("user not found: %w", result.Error)
		}
		return nil, result.Error
	}

	// 更新日時を追加
	updateData["updated_at"] = time.Now()

	// レコードが存在する場合のみ更新を実行
	if err := r.db.Model(&existingUser).Updates(updateData).Error; err != nil {
		return nil, err
	}

	// 更新されたレコードを再度取得して返す
	var updatedUser entity.User
	if err := r.db.Where("user_id = ?", userID).First(&updatedUser).Error; err != nil {
		return nil, err
	}

	return &updatedUser, nil
}

func (r *userRepository) GetUserByID(c *gin.Context, userID string) (*entity.User, error) {
	var user entity.User
	result := r.db.First(&user, "user_id = ?", userID)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
