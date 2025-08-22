package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/usecase"
)

type UserHandler interface {
	UpdateUserServices(c *gin.Context)
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