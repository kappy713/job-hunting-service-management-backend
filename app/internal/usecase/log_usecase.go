package usecase

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

var JST = time.FixedZone("JST", 9*60*60)

type LogUsecase interface {
	GetLogsByUserID(c *gin.Context, userID uuid.UUID) (entity.LogResponse, error)
	UpsertLog(userID uuid.UUID, targetTable, fieldName string) error
	LogFieldUpdateWithErrorHandling(userID uuid.UUID, targetTable, fieldName string)
}

type logUsecase struct {
	lr repository.LogRepository
}

func NewLogUsecase(r repository.LogRepository) LogUsecase {
	return &logUsecase{lr: r}
}

func (u *logUsecase) GetLogsByUserID(c *gin.Context, userID uuid.UUID) (entity.LogResponse, error) {
	logs, err := u.lr.GetLogsByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	logMap := make(map[string]map[string]time.Time)
	for _, logEntry := range logs {
		if logMap[logEntry.TargetTable] == nil {
			logMap[logEntry.TargetTable] = make(map[string]time.Time)
		}
		jstTime := logEntry.UpdatedAt.In(JST)
		logMap[logEntry.TargetTable][logEntry.FieldName] = jstTime
	}

	return logMap, nil
}

func (u *logUsecase) UpsertLog(userID uuid.UUID, targetTable, fieldName string) error {
	return u.lr.UpsertLog(nil, userID, targetTable, fieldName)
}

// LogFieldUpdateWithErrorHandling logs field update with error handling
func (u *logUsecase) LogFieldUpdateWithErrorHandling(userID uuid.UUID, targetTable, fieldName string) {
	err := u.lr.UpsertLog(nil, userID, targetTable, fieldName)
	if err != nil {
		log.Printf("Failed to log field update for table %s, field %s: %v", targetTable, fieldName, err)
	}
}
