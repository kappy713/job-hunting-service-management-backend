package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/usecase"
)

type UserHandler interface {
	GetUserByID(c *gin.Context)
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{uu: u}
}

func (h *userHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	user, err := h.uu.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
