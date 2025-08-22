package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type MynaviUsecase interface {
	GetMynaviByUserID(c *gin.Context, userID uuid.UUID) (*entity.Mynavi, error)
	CreateOrUpdateMynavi(c *gin.Context, userID uuid.UUID, req entity.MynaviData) (*entity.Mynavi, error)
}

type mynaviUsecase struct {
	mr repository.MynaviRepository
}

func NewMynaviUsecase(r repository.MynaviRepository) MynaviUsecase {
	return &mynaviUsecase{mr: r}
}

func (u *mynaviUsecase) GetMynaviByUserID(c *gin.Context, userID uuid.UUID) (*entity.Mynavi, error) {
	mynavi, err := u.mr.GetMynaviByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return mynavi, nil
}

func (u *mynaviUsecase) CreateOrUpdateMynavi(c *gin.Context, userID uuid.UUID, req entity.MynaviData) (*entity.Mynavi, error) {
	mynavi := &entity.Mynavi{
		ID:            userID,
		SelfPromotion: req.SelfPromotion,
		FuturePlan:    req.FuturePlan,
	}

	return u.mr.CreateOrUpdateMynavi(c, mynavi)
}
