package usecase

import (
	"context"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/infrastructure/client"
	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type AIGenerationUsecase interface {
	GenerateServiceProfiles(c *gin.Context, userID uuid.UUID, services []string) (*entity.AIGenerationResponse, error)
}

type aiGenerationUsecase struct {
	repo         repository.AIGenerationRepository
	geminiClient *client.GeminiClient
}

func NewAIGenerationUsecase(repo repository.AIGenerationRepository, geminiClient *client.GeminiClient) AIGenerationUsecase {
	return &aiGenerationUsecase{
		repo:         repo,
		geminiClient: geminiClient,
	}
}

func (u *aiGenerationUsecase) GenerateServiceProfiles(c *gin.Context, userID uuid.UUID, services []string) (*entity.AIGenerationResponse, error) {
	// ユーザー情報を取得
	user, err := u.repo.GetUserByID(c.Request.Context(), userID)
	if err != nil {
		return &entity.AIGenerationResponse{
			UserID:  userID,
			Status:  "error",
			Message: fmt.Sprintf("Failed to get user information: %v", err),
		}, err
	}

	results := make(map[string]interface{})
	successCount := 0
	errorMessages := []string{}

	// 各サービスに対してコンテンツを生成
	for _, serviceName := range services {
		log.Printf("Generating content for service: %s", serviceName)

		// レスポンス用に日本語サービス名を取得
		japaneseServiceName, exists := entity.ConvertServiceNameToJapanese(serviceName)
		if !exists {
			japaneseServiceName = serviceName // 変換できない場合は元の名前を使用
		}

		generatedData, err := u.geminiClient.GenerateServiceContent(c.Request.Context(), serviceName, user)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to generate content for %s: %v", japaneseServiceName, err)
			log.Printf("Error: %s", errorMsg)
			errorMessages = append(errorMessages, errorMsg)
			results[japaneseServiceName] = map[string]interface{}{
				"status": "error",
				"error":  errorMsg,
			}
			continue
		}

		// 生成されたデータを対応するテーブルに保存
		err = u.saveServiceData(c.Request.Context(), serviceName, userID, generatedData)
		if err != nil {
			errorMsg := fmt.Sprintf("Failed to save data for %s: %v", japaneseServiceName, err)
			log.Printf("Error: %s", errorMsg)
			errorMessages = append(errorMessages, errorMsg)
			results[japaneseServiceName] = map[string]interface{}{
				"status": "error",
				"error":  errorMsg,
			}
			continue
		}

		results[japaneseServiceName] = map[string]interface{}{
			"status": "success",
			"data":   generatedData,
		}
		successCount++
		log.Printf("Successfully generated and saved content for service: %s", japaneseServiceName)
	} // レスポンスの構築
	status := "success"
	message := fmt.Sprintf("Successfully generated profiles for %d out of %d services", successCount, len(services))

	if successCount == 0 {
		status = "error"
		message = "Failed to generate profiles for all services"
	} else if successCount < len(services) {
		status = "partial_success"
		message = fmt.Sprintf("Generated profiles for %d out of %d services. Errors: %v", successCount, len(services), errorMessages)
	}

	return &entity.AIGenerationResponse{
		UserID:  userID,
		Results: results,
		Status:  status,
		Message: message,
	}, nil
}

func (u *aiGenerationUsecase) saveServiceData(ctx context.Context, serviceName string, userID uuid.UUID, data map[string]interface{}) error {
	switch serviceName {
	case "supporterz":
		return u.repo.SaveSupporterzData(ctx, userID, data)
	case "career_select":
		return u.repo.SaveCareerSelectData(ctx, userID, data)
	case "one_career":
		return u.repo.SaveOneCareerData(ctx, userID, data)
	case "mynavi":
		return u.repo.SaveMynaviData(ctx, userID, data)
	case "levtech_rookie":
		return u.repo.SaveLevtechRookieData(ctx, userID, data)
	default:
		return fmt.Errorf("unsupported service: %s", serviceName)
	}
}
