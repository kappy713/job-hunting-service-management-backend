package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type SampleUserHandler interface {
	GetAllSampleUsers(c *gin.Context)
}

type sampleUserHandler struct {
	sur usecase.SampleUserUsecase
}

func NewSampleUserHandler(u usecase.SampleUserUsecase) SampleUserHandler {
	return &sampleUserHandler{sur: u}
}

func (h *sampleUserHandler) GetAllSampleUsers(c *gin.Context) {
	var users *[]entity.SampleUser
	users, err := h.sur.GetAllSampleUsers(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, users)
}
