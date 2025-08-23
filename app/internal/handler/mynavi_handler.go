package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type MynaviHandler interface {
	GetMynaviByID(c *gin.Context)
	CreateOrUpdateMynavi(c *gin.Context)
}

type mynaviHandler struct {
	mu usecase.MynaviUsecase
}

func NewMynaviHandler(u usecase.MynaviUsecase) MynaviHandler {
	return &mynaviHandler{mu: u}
}

func (h *mynaviHandler) GetMynaviByID(c *gin.Context) {
	// URLパラメータからIDを取得
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID parameter is required"})
		return
	}

	// UUIDのパース
	userID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	mynavi, err := h.mu.GetMynaviByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if mynavi == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Mynavi record not found"})
		return
	}

	c.JSON(http.StatusOK, mynavi)
}

func (h *mynaviHandler) CreateOrUpdateMynavi(c *gin.Context) {
	var req entity.CreateMynaviRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// UUIDのパース
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// Usecaseに渡すのはdataの部分のみ
	mynavi, err := h.mu.CreateOrUpdateMynavi(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, mynavi)
}
