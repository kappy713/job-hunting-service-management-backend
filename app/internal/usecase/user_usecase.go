package usecase

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type UserUsecase interface {
	UpdateUserServices(c *gin.Context, userID string, services []string) error
	UpdateUser(c *gin.Context, userID uuid.UUID, req entity.UserData) (*entity.User, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{ur: r}
}

func (u *userUsecase) UpdateUserServices(c *gin.Context, userID string, services []string) error {
	// ここでバリデーションやビジネスロジックを追加
	return u.ur.UpdateUserServices(c, userID, services)
}

func (u *userUsecase) UpdateUser(c *gin.Context, userID uuid.UUID, req entity.UserData) (*entity.User, error) {
	// リクエストから entity に変換
	now := time.Now()
	user := &entity.User{
		UserID:        userID,
		LastName:      req.LastName,
		FirstName:     req.FirstName,
		BirthDate:     req.BirthDate,
		Age:           req.Age,
		University:    req.University,
		Category:      req.Category,
		Faculty:       req.Faculty,
		TargetJobType: req.TargetJobType,
		Services:      pq.StringArray(req.Services),
		UpdatedAt:     now,
	}

	// Gradeが指定されている場合のみ設定
	if req.Grade != nil {
		user.Grade = *req.Grade
	}

	return u.ur.UpdateUser(c, user)
}
