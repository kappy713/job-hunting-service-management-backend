package usecase

import (
	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type SampleUserUsecase interface {
	GetAllSampleUsers(c *gin.Context) (*[]entity.SampleUser, error)
}

type sampleUserUsecase struct {
	sur repository.SampleUserRepository
}

func NewSampleUserUsecase(r repository.SampleUserRepository) SampleUserUsecase {
	return &sampleUserUsecase{sur: r}
}

func (u *sampleUserUsecase) GetAllSampleUsers(c *gin.Context) (*[]entity.SampleUser, error) {
	users, err := u.sur.GetAllSampleUsers(c)
	if err != nil {
		return nil, err
	}

	return &users, nil
}
