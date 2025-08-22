package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type CareerSelectHandler interface {
	GetCareerSelectByID(c *gin.Context)
	CreateOrUpdateCareerSelect(c *gin.Context)
}

type careerSelectHandler struct {
	csu usecase.CareerSelectUsecase
}

func NewCareerSelectHandler(u usecase.CareerSelectUsecase) CareerSelectHandler {
	return &careerSelectHandler{csu: u}
}

func (h *careerSelectHandler) GetCareerSelectByID(c *gin.Context) {
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

	careerSelect, err := h.csu.GetCareerSelectByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if careerSelect == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "CareerSelect record not found"})
		return
	}

	c.JSON(http.StatusOK, careerSelect)
}

func (h *careerSelectHandler) CreateOrUpdateCareerSelect(c *gin.Context) {
	var req entity.CreateCareerSelectRequest
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
	careerSelect, err := h.csu.CreateOrUpdateCareerSelect(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, careerSelect)
}
