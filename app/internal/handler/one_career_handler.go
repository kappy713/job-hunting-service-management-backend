package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type OneCareerHandler interface {
	GetOneCareerByID(c *gin.Context)
	CreateOrUpdateOneCareer(c *gin.Context)
}

type oneCareerHandler struct {
	ocu usecase.OneCareerUsecase
}

func NewOneCareerHandler(u usecase.OneCareerUsecase) OneCareerHandler {
	return &oneCareerHandler{ocu: u}
}

func (h *oneCareerHandler) GetOneCareerByID(c *gin.Context) {
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

	oneCareer, err := h.ocu.GetOneCareerByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if oneCareer == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "OneCareer record not found"})
		return
	}

	c.JSON(http.StatusOK, oneCareer)
}

func (h *oneCareerHandler) CreateOrUpdateOneCareer(c *gin.Context) {
	var req entity.CreateOneCareerRequest
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
	oneCareer, err := h.ocu.CreateOrUpdateOneCareer(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, oneCareer)
}
