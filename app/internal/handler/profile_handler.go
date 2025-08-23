package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type ProfileHandler interface {
	GetProfileByUserID(c *gin.Context)
	CreateOrUpdateProfile(c *gin.Context)
}

type profileHandler struct {
	pu usecase.ProfileUsecase
}

func NewProfileHandler(u usecase.ProfileUsecase) ProfileHandler {
	return &profileHandler{pu: u}
}

func (h *profileHandler) GetProfileByUserID(c *gin.Context) {
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

	profile, err := h.pu.GetProfileByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Profile record not found"})
		return
	}

	c.JSON(http.StatusOK, profile)
}

func (h *profileHandler) CreateOrUpdateProfile(c *gin.Context) {
	var req entity.CreateProfileRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	profile, err := h.pu.CreateOrUpdateProfile(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}
