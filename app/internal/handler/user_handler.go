package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"job-hunting-service-management-backend/app/internal/usecase"
	"job-hunting-service-management-backend/app/internal/entity" // entityパッケージをインポート
	"github.com/google/uuid" // uuidを扱うためにインポート
)

type UserHandler interface {
	UpdateUserServices(c *gin.Context)
	CreateUser(c *gin.Context)
}

type userHandler struct {
	uu usecase.UserUsecase
}

func NewUserHandler(u usecase.UserUsecase) UserHandler {
	return &userHandler{uu: u}
}

// 新しいユーザー作成のリクエストボディを定義
type createRequest struct {
	UserID    string   `json:"user_id" binding:"required"`
	LastName  string   `json:"last_name" binding:"required"`
	FirstName string   `json:"first_name" binding:"required"`
	BirthDate string   `json:"birth_date"`
	Age       int      `json:"age" binding:"required"`
	University string  `json:"university"`
	Category string    `json:"category"`
	Faculty string     `json:"faculty"`
	Grade int          `json:"grade"`
	TargetJobType string `json:"target_job_type" binding:"required"`
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

// 新しいユーザーを作成するAPIの実装
func (h *userHandler) CreateUser(c *gin.Context) {
	var req createRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// リクエストデータをエンティティにマッピング
	user := &entity.User{
		UserID: uuid.MustParse(req.UserID),
		LastName: req.LastName,
		FirstName: req.FirstName,
		Age: req.Age,
		University: req.University,
		Category: req.Category,
		Faculty: req.Faculty,
		Grade: req.Grade,
		TargetJobType: req.TargetJobType,
	}

	if err := h.uu.CreateUser(c, user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}