package usecase

import (
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
func (u *profileUsecase) CreateOrUpdateProfile(c *gin.Context, userID uuid.UUID, req entity.ProfileData) (*entity.Profile, error) {
	profile := &entity.Profile{
		ID:                        userID,
		CareerVision:              req.CareerVision,
		SelfPromotion:             req.SelfPromotion,
		StudentExperience:         req.StudentExperience,
		Research:                  req.Research,
		Products:                  pq.StringArray(req.Products),
		ProductDescriptions:       pq.StringArray(req.ProductDescriptions),
		Skills:                    pq.StringArray(req.Skills),
		SkillDescriptions:         pq.StringArray(req.SkillDescriptions),
		Interns:                   pq.StringArray(req.Interns),
		InternDescriptions:        pq.StringArray(req.InternDescriptions),
		Organization:              req.Organization,
		Certifications:            pq.StringArray(req.Certifications),
		CertificationDescriptions: pq.StringArray(req.CertificationDescriptions),
		DesiredJobType:            req.DesiredJobType,
		CompanySelectionCriteria:  req.CompanySelectionCriteria,
		EngineerAspiration:        req.EngineerAspiration,
	}

	result, err := u.pr.CreateOrUpdateProfile(c, profile)
	if err != nil {
		return nil, err
	}

	// 更新されたフィールドのログを記録
	u.logFieldUpdates(userID, req)

	return result, nil
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
