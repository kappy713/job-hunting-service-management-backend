package usecase

import (
	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type UserUsecase interface {
	GetUserByID(c *gin.Context, userID string) (*entity.User, error)
}

type userUsecase struct {
	ur repository.UserRepository
}

func NewUserUsecase(r repository.UserRepository) UserUsecase {
	return &userUsecase{ur: r}
}

func (u *userUsecase) GetUserByID(c *gin.Context, userID string) (*entity.User, error) {
	user, err := u.ur.GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}

	// ここでビジネスロジックを追加（例：データ変換、追加のチェックなど）

	return user, nil
}
