package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type SupporterzHandler interface {
	GetSupporterzByID(c *gin.Context)
	CreateOrUpdateSupporterz(c *gin.Context)
}

type supporterzHandler struct {
	su usecase.SupporterzUsecase
}

func NewSupporterzHandler(u usecase.SupporterzUsecase) SupporterzHandler {
	return &supporterzHandler{su: u}
}

func (h *supporterzHandler) GetSupporterzByID(c *gin.Context) {
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

	supporterz, err := h.su.GetSupporterzByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if supporterz == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Supporterz record not found"})
		return
	}

	c.JSON(http.StatusOK, supporterz)
}

func (h *supporterzHandler) CreateOrUpdateSupporterz(c *gin.Context) {
	var req entity.CreateSupporterzRequest
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
	supporterz, err := h.su.CreateOrUpdateSupporterz(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, supporterz)
}
