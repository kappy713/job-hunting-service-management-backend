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
	CreateUser(c *gin.Context)
	UpdateUser(c *gin.Context)
	GetUserByID(c *gin.Context)
	GetUserServices(c *gin.Context)
	GetUserServiceDetails(c *gin.Context)
}

type userHandler struct {
	uu  usecase.UserUsecase
	su  usecase.SupporterzUsecase
	csu usecase.CareerSelectUsecase
	lru usecase.LevtechRookieUsecase
	mu  usecase.MynaviUsecase
	ocu usecase.OneCareerUsecase
	pu  usecase.ProfileUsecase
}

func NewUserHandler(u usecase.UserUsecase, su usecase.SupporterzUsecase, csu usecase.CareerSelectUsecase, lru usecase.LevtechRookieUsecase, mu usecase.MynaviUsecase, ocu usecase.OneCareerUsecase, pu usecase.ProfileUsecase) UserHandler {
	return &userHandler{
		uu:  u,
		su:  su,
		csu: csu,
		lru: lru,
		mu:  mu,
		ocu: ocu,
		pu:  pu,
	}
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
	var req entity.CreateUserRequest
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

	// Usecaseに渡すのはuser_idとdataの部分
	user, err := h.uu.CreateUser(c, userID, req.Data)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
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

func (h *userHandler) GetUserByID(c *gin.Context) {
	userID := c.Param("userID")
	user, err := h.uu.GetUserByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func (h *userHandler) GetUserServices(c *gin.Context) {
	userID := c.Param("userID")
	services, err := h.uu.GetUserServices(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"services": services})
}

func (h *userHandler) GetUserServiceDetails(c *gin.Context) {
	userID := c.Param("userID")

	// UUIDに変換
	userUUID, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user ID format"})
		return
	}

	// ユーザーのサービス一覧を取得
	services, err := h.uu.GetUserServices(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	serviceData := make(map[string]interface{})

	// 各サービスの詳細データを取得
	for _, serviceName := range services {
		switch serviceName {
		case "サポーターズ":
			data, err := h.su.GetSupporterzByUserID(c, userUUID)
			if err == nil {
				serviceData["supporterz"] = data
			} else {
				serviceData["supporterz"] = gin.H{"error": err.Error()}
			}

		case "マイナビ":
			data, err := h.mu.GetMynaviByUserID(c, userUUID)
			if err == nil {
				serviceData["mynavi"] = data
			} else {
				serviceData["mynavi"] = gin.H{"error": err.Error()}
			}

		case "レバテックルーキー":
			data, err := h.lru.GetLevtechRookieByUserID(c, userUUID)
			if err == nil {
				serviceData["levtech_rookie"] = data
			} else {
				serviceData["levtech_rookie"] = gin.H{"error": err.Error()}
			}

		case "ワンキャリア":
			data, err := h.ocu.GetOneCareerByUserID(c, userUUID)
			if err == nil {
				serviceData["one_career"] = data
			} else {
				serviceData["one_career"] = gin.H{"error": err.Error()}
			}

		case "キャリアセレクト":
			data, err := h.csu.GetCareerSelectByUserID(c, userUUID)
			if err == nil {
				serviceData["career_select"] = data
			} else {
				serviceData["career_select"] = gin.H{"error": err.Error()}
			}
		}
	}

	// プロフィール情報を取得
	profile, err := h.pu.GetProfileByUserID(c, userUUID)
	if err != nil {
		// プロフィールが存在しない場合はnullを設定
		serviceData["profile"] = nil
	} else {
		serviceData["profile"] = profile
	}

	c.JSON(http.StatusOK, gin.H{
		"services": services,
		"data":     serviceData,
	})
}
