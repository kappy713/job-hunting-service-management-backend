package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

type AIGenerationHandler interface {
	GenerateServiceProfiles(c *gin.Context)
}

type aiGenerationHandler struct {
	aiUsecase usecase.AIGenerationUsecase
}

func NewAIGenerationHandler(aiUsecase usecase.AIGenerationUsecase) AIGenerationHandler {
	return &aiGenerationHandler{
		aiUsecase: aiUsecase,
	}
}

func (h *aiGenerationHandler) GenerateServiceProfiles(c *gin.Context) {
	var req entity.AIGenerationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request body",
			"details": err.Error(),
		})
		return
	}

	// 日本語サービス名をアルファベットに変換
	convertedServices := make([]string, 0, len(req.Services))
	for _, service := range req.Services {
		// まず英語名として有効かチェック
		validEnglishServices := map[string]bool{
			"supporterz":     true,
			"career_select":  true,
			"one_career":     true,
			"mynavi":         true,
			"levtech_rookie": true,
		}

		if validEnglishServices[service] {
			// 既に英語名の場合はそのまま使用
			convertedServices = append(convertedServices, service)
		} else {
			// 日本語名の場合は変換
			englishName, exists := entity.ConvertServiceName(service)
			if !exists {
				c.JSON(http.StatusBadRequest, gin.H{
					"error":                   "Invalid service name",
					"service":                 service,
					"valid_services_japanese": []string{"サポーターズ", "キャリアセレクト", "ワンキャリア", "レバテックルーキー", "マイナビ"},
					"valid_services_english":  []string{"supporterz", "career_select", "one_career", "mynavi", "levtech_rookie"},
				})
				return
			}
			convertedServices = append(convertedServices, englishName)
		}
	}

	if len(convertedServices) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "At least one service must be specified",
		})
		return
	}

	// AI生成処理を実行（変換されたサービス名を使用）
	response, err := h.aiUsecase.GenerateServiceProfiles(c, req.UserID, convertedServices)
	if err != nil && response.Status == "error" {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to generate service profiles",
			"details": err.Error(),
		})
		return
	}

	// レスポンスのステータスに応じてHTTPステータスコードを設定
	var statusCode int
	switch response.Status {
	case "success":
		statusCode = http.StatusOK
	case "partial_success":
		statusCode = http.StatusPartialContent
	case "error":
		statusCode = http.StatusInternalServerError
	default:
		statusCode = http.StatusOK
	}

	c.JSON(statusCode, response)
}
