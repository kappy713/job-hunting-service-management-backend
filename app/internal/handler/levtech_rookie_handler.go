package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type LevtechRookieHandler interface {
	GetLevtechRookieByID(c *gin.Context)
	CreateOrUpdateLevtechRookie(c *gin.Context)
}

type levtechRookieHandler struct {
	lru usecase.LevtechRookieUsecase
}

func NewLevtechRookieHandler(u usecase.LevtechRookieUsecase) LevtechRookieHandler {
	return &levtechRookieHandler{lru: u}
}

func (h *levtechRookieHandler) GetLevtechRookieByID(c *gin.Context) {
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

	levtechRookie, err := h.lru.GetLevtechRookieByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if levtechRookie == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "LevtechRookie record not found"})
		return
	}

	c.JSON(http.StatusOK, levtechRookie)
}

func (h *levtechRookieHandler) CreateOrUpdateLevtechRookie(c *gin.Context) {
	var req entity.CreateLevtechRookieRequest
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
	levtechRookie, err := h.lru.CreateOrUpdateLevtechRookie(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, levtechRookie)
}
