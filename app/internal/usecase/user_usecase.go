package usecase

import (
	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/repository"
)

type UserUsecase interface {
	UpdateUserServices(c *gin.Context, userID string, services []string) error
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
