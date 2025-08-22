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
	lu LogUsecase
}

func NewMynaviUsecase(r repository.MynaviRepository, l LogUsecase) MynaviUsecase {
	return &mynaviUsecase{mr: r, lu: l}
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

	result, err := u.mr.CreateOrUpdateMynavi(c, mynavi)
	if err != nil {
		return nil, err
	}

	// 更新されたフィールドのログを記録
	u.logFieldUpdates(c, userID, req)

	return result, nil
}

// フィールド更新のログを記録
func (u *mynaviUsecase) logFieldUpdates(c *gin.Context, userID uuid.UUID, req entity.MynaviData) {
	targetTable := "mynavi"

	// 各フィールドが空でなければログを記録
	if req.SelfPromotion != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "self_promotion")
	}
	if req.FuturePlan != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "future_plan")
	}
}
