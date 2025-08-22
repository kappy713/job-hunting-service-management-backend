package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type UserHandler interface {
	UpdateUserServices(c *gin.Context)
	UpdateUser(c *gin.Context)
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{uu: u}
}

type updateServicesRequest struct {
	UserID   string   `json:"user_id" binding:"required"`
	Services []string `json:"services"`
}

func (h *userHandler) UpdateUserServices(c *gin.Context) {
	var req updateServicesRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.uu.UpdateUserServices(c, req.UserID, req.Services); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Services updated successfully"})
}

func (h *userHandler) UpdateUser(c *gin.Context) {
	var req entity.UpdateUserRequest
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
	user, err := h.uu.UpdateUser(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}
