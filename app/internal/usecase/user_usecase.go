package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type UserUsecase interface {
	UpdateUserServices(c *gin.Context, userID string, services []string) error
	CreateUser(c *gin.Context, userID uuid.UUID, req entity.CreateUserData) (*entity.User, error)
	UpdateUser(c *gin.Context, userID uuid.UUID, req entity.UserData) (*entity.User, error)
	GetUserByID(c *gin.Context, userID string) (*entity.User, error)
	GetUserServices(c *gin.Context, userID string) ([]string, error)
	GetUserServiceDetails(c *gin.Context, userID string) (map[string]interface{}, error)
}

type userUsecase struct {
	ur        repository.UserRepository
	aiUsecase AIGenerationUsecase
}

func NewUserUsecase(r repository.UserRepository, aiUsecase AIGenerationUsecase) UserUsecase {
	return &userUsecase{
		ur:        r,
		aiUsecase: aiUsecase,
	}
}

func (u *userUsecase) UpdateUserServices(c *gin.Context, userID string, services []string) error {
	// サービス情報を更新
	if err := u.ur.UpdateUserServices(c, userID, services); err != nil {
		return err
	}

	// サービスが設定されている場合、AI生成を実行
	if len(services) > 0 {
		// UUIDに変換
		userUUID, err := uuid.Parse(userID)
		if err != nil {
			// UUID変換エラーの場合はログに記録するが、処理は継続
			// TODO: ログ出力を追加
			return nil
		}

		// 日本語サービス名を英語名に変換
		convertedServices := make([]string, 0, len(services))
		for _, service := range services {
			if englishName, exists := entity.ConvertServiceName(service); exists {
				convertedServices = append(convertedServices, englishName)
			}
			// 変換できないサービス名はスキップ
		}

		// 変換されたサービスがある場合のみAI生成を実行
		if len(convertedServices) > 0 {
			// AI生成を実行（エラーが発生してもサービス更新は成功とする）
			if _, err := u.aiUsecase.GenerateServiceProfiles(c, userUUID, convertedServices); err != nil {
				// AI生成エラーはログに記録するが、処理は継続
				fmt.Printf("AI generation failed for user %s: %v\n", userUUID, err)
			}
		}
	}

	return nil
}

func (u *userUsecase) CreateUser(c *gin.Context, userID uuid.UUID, req entity.CreateUserData) (*entity.User, error) {
	// 全ての項目が入ってくる想定なので、直接マップに設定
	updateData := map[string]interface{}{
		"last_name":       req.LastName,
		"first_name":      req.FirstName,
		"age":             req.Age,
		"university":      req.University,
		"category":        req.Category,
		"faculty":         req.Faculty,
		"grade":           req.Grade,
		"target_job_type": req.TargetJobType,
	}

	// birth_dateは nilの可能性があるので条件付きで追加
	if req.BirthDate != nil {
		updateData["birth_date"] = req.BirthDate
	}

	// リポジトリに更新用データマップを渡す
	return u.ur.UpdateUser(c, userID.String(), updateData)
}

func (u *userUsecase) UpdateUser(c *gin.Context, userID uuid.UUID, req entity.UserData) (*entity.User, error) {
	// 更新するフィールドをマップに格納
	updateData := make(map[string]interface{})

	if req.LastName != "" {
		updateData["last_name"] = req.LastName
	}
	if req.FirstName != "" {
		updateData["first_name"] = req.FirstName
	}
	if req.BirthDate != nil {
		updateData["birth_date"] = req.BirthDate
	}
	if req.Age != 0 {
		updateData["age"] = req.Age
	}
	if req.University != "" {
		updateData["university"] = req.University
	}
	if req.Category != "" {
		updateData["category"] = req.Category
	}
	if req.Faculty != "" {
		updateData["faculty"] = req.Faculty
	}
	if req.TargetJobType != "" {
		updateData["target_job_type"] = req.TargetJobType
	}
	if req.Services != nil {
		updateData["services"] = pq.StringArray(req.Services)
	}
	// Gradeが指定されている場合のみ更新対象に追加
	if req.Grade != nil {
		updateData["grade"] = *req.Grade
	}

	// リポジトリに更新用データマップを渡す
	return u.ur.UpdateUser(c, userID.String(), updateData)
}

func (u *userUsecase) GetUserByID(c *gin.Context, userID string) (*entity.User, error) {
	user, err := u.ur.GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}

	// ここでビジネスロジックを追加（例：データ変換、追加のチェックなど）

	return user, nil
}

func (u *userUsecase) GetUserServices(c *gin.Context, userID string) ([]string, error) {
	user, err := u.ur.GetUserByID(c, userID)
	if err != nil {
		return nil, err
	}

	// user.Servicesがpq.StringArrayの場合、[]stringに変換
	services := make([]string, len(user.Services))
	copy(services, user.Services)

	return services, nil
}

func (u *userUsecase) GetUserServiceDetails(c *gin.Context, userID string) (map[string]interface{}, error) {
	// ユーザーのサービス一覧を取得
	services, err := u.GetUserServices(c, userID)
	if err != nil {
		return nil, err
	}

	// サービスリストとAPIエンドポイントのマッピングを返す
	serviceEndpoints := make(map[string]string)
	for _, serviceName := range services {
		switch serviceName {
		case "サポーターズ":
			serviceEndpoints["supporterz"] = "/api/supporterz/" + userID
		case "マイナビ":
			serviceEndpoints["mynavi"] = "/api/mynavi/" + userID
		case "レバテックルーキー":
			serviceEndpoints["levtech_rookie"] = "/api/levtech-rookie/" + userID
		case "ワンキャリア":
			serviceEndpoints["one_career"] = "/api/one-career/" + userID
		case "キャリアセレクト":
			serviceEndpoints["career_select"] = "/api/career-select/" + userID
		}
	}

	result := map[string]interface{}{
		"user_id":           userID,
		"services":          services,
		"service_endpoints": serviceEndpoints,
	}

	return result, nil
}
