package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/usecase"
)

type LogHandler interface {
	GetLogsByUserID(c *gin.Context)
}

type logHandler struct {
	lu usecase.LogUsecase
}

func NewLogHandler(u usecase.LogUsecase) LogHandler {
	return &logHandler{lu: u}
}

func (h *logHandler) GetLogsByUserID(c *gin.Context) {
	// パスパラメータからuser_idを取得
	userIDStr := c.Param("id")
	if userIDStr == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User ID is required"})
		return
	}

	// UUIDのパース
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// ログ情報を取得
	logs, err := h.lu.GetLogsByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, logs)
}
