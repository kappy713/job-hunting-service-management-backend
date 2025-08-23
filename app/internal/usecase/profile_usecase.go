package usecase

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

// ProfileUsecase はプロフィール関連のビジネスロジックを定義するインターフェースです。
type ProfileUsecase interface {
	GetProfileByUserID(c *gin.Context, userID uuid.UUID) (*entity.Profile, error)
	CreateOrUpdateProfile(c *gin.Context, userID uuid.UUID, req entity.ProfileData) (*entity.Profile, error)
}

type profileUsecase struct {
	pr repository.ProfileRepository
	lu LogUsecase
}

// NewProfileUsecase は新しいProfileUsecaseのインスタンスを作成します。
func NewProfileUsecase(r repository.ProfileRepository, l LogUsecase) ProfileUsecase {
	return &profileUsecase{pr: r, lu: l}
}

// GetProfileByUserID はユーザーIDに基づいてプロフィール情報を取得します。
func (u *profileUsecase) GetProfileByUserID(c *gin.Context, userID uuid.UUID) (*entity.Profile, error) {
	profile, err := u.pr.GetProfileByUserID(c, userID)
	if err != nil {
		return nil, err
	}
	return profile, nil
}

// CreateOrUpdateProfile はプロフィール情報をDBに登録または更新します。
// 空のフィールドが含まれている場合は、既存の値を保持し、空でないフィールドのみを更新します。
func (u *profileUsecase) CreateOrUpdateProfile(c *gin.Context, userID uuid.UUID, req entity.ProfileData) (*entity.Profile, error) {
	// リクエストデータの基本検証
	if err := validateProfileData(req); err != nil {
		return nil, fmt.Errorf("validation failed: %w", err)
	}

	// 既存のプロフィールを取得（存在しない場合は新規作成）
	existingProfile, err := u.pr.GetProfileByUserID(c, userID)
	if err != nil {
		// プロフィールが存在しない場合は新規作成として処理
		existingProfile = &entity.Profile{
			ID: userID,
		}
	}

	// 部分更新ロジック：空でないフィールドのみを更新
	profile := u.mergeProfileData(existingProfile, req)

	// プロフィールを保存
	result, err := u.pr.CreateOrUpdateProfile(c, profile)
	if err != nil {
		return nil, fmt.Errorf("failed to create or update profile: %w", err)
	}

	// 更新されたフィールドのログを記録（エラーは無視）
	u.logFieldUpdates(userID, req)

	return result, nil
}

// validateProfileData はプロフィールデータの基本的なバリデーションを行います
func validateProfileData(req entity.ProfileData) error {
	// 必須フィールドのチェック（必要に応じて追加）
	// 現在は基本的なチェックのみ実装

	// 文字列の長さチェック
	if len(req.CareerVision) > 2000 {
		return fmt.Errorf("career_vision exceeds maximum length of 2000 characters")
	}
	if len(req.SelfPromotion) > 5000 {
		return fmt.Errorf("self_promotion exceeds maximum length of 5000 characters")
	}
	if len(req.StudentExperience) > 5000 {
		return fmt.Errorf("student_experience exceeds maximum length of 5000 characters")
	}
	if len(req.Research) > 2000 {
		return fmt.Errorf("research exceeds maximum length of 2000 characters")
	}
	if len(req.Organization) > 2000 {
		return fmt.Errorf("organization exceeds maximum length of 2000 characters")
	}
	if len(req.DesiredJobType) > 2000 {
		return fmt.Errorf("desired_job_type exceeds maximum length of 2000 characters")
	}
	if len(req.CompanySelectionCriteria) > 2000 {
		return fmt.Errorf("company_selection_criteria exceeds maximum length of 2000 characters")
	}
	if len(req.EngineerAspiration) > 2000 {
		return fmt.Errorf("engineer_aspiration exceeds maximum length of 2000 characters")
	}

	return nil
}

// mergeProfileData は既存のプロフィールデータと新しいリクエストデータをマージします。
// 空でないフィールドのみを更新し、空のフィールドは既存の値を保持します。
func (u *profileUsecase) mergeProfileData(existing *entity.Profile, req entity.ProfileData) *entity.Profile {
	// 既存のプロフィールをベースにコピー
	result := &entity.Profile{
		ID:                        existing.ID,
		CareerVision:              existing.CareerVision,
		SelfPromotion:             existing.SelfPromotion,
		StudentExperience:         existing.StudentExperience,
		Research:                  existing.Research,
		Products:                  existing.Products,
		ProductDescriptions:       existing.ProductDescriptions,
		Skills:                    existing.Skills,
		SkillDescriptions:         existing.SkillDescriptions,
		Interns:                   existing.Interns,
		InternDescriptions:        existing.InternDescriptions,
		Organization:              existing.Organization,
		Certifications:            existing.Certifications,
		CertificationDescriptions: existing.CertificationDescriptions,
		DesiredJobType:            existing.DesiredJobType,
		CompanySelectionCriteria:  existing.CompanySelectionCriteria,
		EngineerAspiration:        existing.EngineerAspiration,
	}

	// 文字列フィールドの部分更新
	if req.CareerVision != "" {
		result.CareerVision = req.CareerVision
	}
	if req.SelfPromotion != "" {
		result.SelfPromotion = req.SelfPromotion
	}
	if req.StudentExperience != "" {
		result.StudentExperience = req.StudentExperience
	}
	if req.Research != "" {
		result.Research = req.Research
	}
	if req.Organization != "" {
		result.Organization = req.Organization
	}
	if req.DesiredJobType != "" {
		result.DesiredJobType = req.DesiredJobType
	}
	if req.CompanySelectionCriteria != "" {
		result.CompanySelectionCriteria = req.CompanySelectionCriteria
	}
	if req.EngineerAspiration != "" {
		result.EngineerAspiration = req.EngineerAspiration
	}

	// スライスフィールドの部分更新（空でない場合のみ）
	if len(req.Products) > 0 {
		result.Products = pq.StringArray(req.Products)
	}
	if len(req.ProductDescriptions) > 0 {
		result.ProductDescriptions = pq.StringArray(req.ProductDescriptions)
	}
	if len(req.Skills) > 0 {
		result.Skills = pq.StringArray(req.Skills)
	}
	if len(req.SkillDescriptions) > 0 {
		result.SkillDescriptions = pq.StringArray(req.SkillDescriptions)
	}
	if len(req.Interns) > 0 {
		result.Interns = pq.StringArray(req.Interns)
	}
	if len(req.InternDescriptions) > 0 {
		result.InternDescriptions = pq.StringArray(req.InternDescriptions)
	}
	if len(req.Certifications) > 0 {
		result.Certifications = pq.StringArray(req.Certifications)
	}
	if len(req.CertificationDescriptions) > 0 {
		result.CertificationDescriptions = pq.StringArray(req.CertificationDescriptions)
	}

	return result
}

// logFieldUpdates は更新されたフィールドのログを記録します。
func (u *profileUsecase) logFieldUpdates(userID uuid.UUID, req entity.ProfileData) {
	targetTable := "profiles"

	// 各フィールドが空でなければログを記録
	if req.CareerVision != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "career_vision")
	}
	if req.SelfPromotion != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "self_promotion")
	}
	if req.StudentExperience != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "student_experience")
	}
	if req.Research != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "research")
	}
	if len(req.Products) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "products")
	}
	if len(req.ProductDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "product_descriptions")
	}
	if len(req.Skills) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "skills")
	}
	if len(req.SkillDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "skill_descriptions")
	}
	if len(req.Interns) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "interns")
	}
	if len(req.InternDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "intern_descriptions")
	}
	if req.Organization != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "organization")
	}
	if len(req.Certifications) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "certifications")
	}
	if len(req.CertificationDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "certification_descriptions")
	}
	if req.DesiredJobType != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "desired_job_type")
	}
	if req.CompanySelectionCriteria != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "company_selection_criteria")
	}
	if req.EngineerAspiration != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "engineer_aspiration")
	}
}
