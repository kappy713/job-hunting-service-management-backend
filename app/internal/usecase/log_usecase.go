package usecase

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type LogUsecase interface {
	GetLogsByUserID(c *gin.Context, userID uuid.UUID) (*entity.LogResponse, error)
	UpsertLog(c *gin.Context, userID uuid.UUID, targetTable, fieldName string) error
}

type logUsecase struct {
	lr repository.LogRepository
}

func NewLogUsecase(r repository.LogRepository) LogUsecase {
	return &logUsecase{lr: r}
}

func (u *logUsecase) GetLogsByUserID(c *gin.Context, userID uuid.UUID) (*entity.LogResponse, error) {
	logs, err := u.lr.GetLogsByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	// テーブル名ごとにフィールド名と更新日時をグループ化
	response := make(entity.LogResponse)

	for _, log := range logs {
		if _, exists := response[log.TargetTable]; !exists {
			response[log.TargetTable] = make(map[string]time.Time)
		}
		response[log.TargetTable][log.FieldName] = log.UpdatedAt
	}

	return &response, nil
}

func (u *logUsecase) UpsertLog(c *gin.Context, userID uuid.UUID, targetTable, fieldName string) error {
	return u.lr.UpsertLog(c, userID, targetTable, fieldName)
}
