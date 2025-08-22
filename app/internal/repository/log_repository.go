package repository

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	"job-hunting-service-management-backend/app/internal/entity"
)

var (
	JST = time.FixedZone("Asia/Tokyo", 9*60*60)
)

type LogRepository interface {
	GetLogsByUserID(c *gin.Context, userID uuid.UUID) ([]entity.Log, error)
	UpsertLog(c *gin.Context, userID uuid.UUID, targetTable, fieldName string) error
}

type logRepository struct {
	db *gorm.DB
}

func NewLogRepository(d *gorm.DB) LogRepository {
	return &logRepository{db: d}
}

func (r *logRepository) GetLogsByUserID(c *gin.Context, userID uuid.UUID) ([]entity.Log, error) {
	var logs []entity.Log
	result := r.db.Where("user_id = ?", userID).Find(&logs)
	if result.Error != nil {
		return nil, result.Error
	}
	return logs, nil
}

func (r *logRepository) UpsertLog(c *gin.Context, userID uuid.UUID, targetTable, fieldName string) error {
	// 既存のレコードを検索（ログレベルを一時的に下げる）
	var existingLog entity.Log
	result := r.db.Session(&gorm.Session{Logger: r.db.Logger.LogMode(logger.Silent)}).
		Where("user_id = ? AND target_table = ? AND field_name = ?", userID, targetTable, fieldName).
		First(&existingLog)

	if result.Error != nil && result.Error != gorm.ErrRecordNotFound {
		return result.Error
	}

	if result.Error == gorm.ErrRecordNotFound {
		// レコードが存在しない場合は新規作成
		newLog := entity.Log{
			ID:          uuid.New(),
			UserID:      userID,
			TargetTable: targetTable,
			FieldName:   fieldName,
			UpdatedAt:   time.Now().In(JST),
		}
		return r.db.Create(&newLog).Error
	} else {
		// レコードが存在する場合はupdated_atのみ更新
		return r.db.Model(&existingLog).Update("updated_at", time.Now().In(JST)).Error
	}
}
