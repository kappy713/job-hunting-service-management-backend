package usecase

import (
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
				// TODO: ログ出力を追加
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
