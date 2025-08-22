package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type SupporterzUsecase interface {
	GetSupporterzByUserID(c *gin.Context, userID uuid.UUID) (*entity.Supporterz, error)
	CreateOrUpdateSupporterz(c *gin.Context, userID uuid.UUID, req entity.SupporterzData) (*entity.Supporterz, error)
}

type supporterzUsecase struct {
	sr repository.SupporterzRepository
	lu LogUsecase
}

func NewSupporterzUsecase(r repository.SupporterzRepository, l LogUsecase) SupporterzUsecase {
	return &supporterzUsecase{sr: r, lu: l}
}

func (u *supporterzUsecase) GetSupporterzByUserID(c *gin.Context, userID uuid.UUID) (*entity.Supporterz, error) {
	supporterz, err := u.sr.GetSupporterzByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return supporterz, nil
}

func (u *supporterzUsecase) CreateOrUpdateSupporterz(c *gin.Context, userID uuid.UUID, req entity.SupporterzData) (*entity.Supporterz, error) {
	supporterz := &entity.Supporterz{
		ID:                           userID,
		CareerVision:                 req.CareerVision,
		SelfPromotion:                req.SelfPromotion,
		Skills:                       pq.StringArray(req.Skills),
		SkillDescriptions:            pq.StringArray(req.SkillDescriptions),
		InternExperiences:            pq.StringArray(req.InternExperiences),
		InternExperienceDescriptions: pq.StringArray(req.InternExperienceDescriptions),
		Products:                     pq.StringArray(req.Products),
		ProductTechStacks:            pq.StringArray(req.ProductTechStacks),
		ProductDescriptions:          pq.StringArray(req.ProductDescriptions),
		Researches:                   pq.StringArray(req.Researches),
		ResearchDescriptions:         pq.StringArray(req.ResearchDescriptions),
	}

	result, err := u.sr.CreateOrUpdateSupporterz(c, supporterz)
	if err != nil {
		return nil, err
	}

	// 更新されたフィールドのログを記録
	u.logFieldUpdates(c, userID, req)

	return result, nil
}

// フィールド更新のログを記録
func (u *supporterzUsecase) logFieldUpdates(c *gin.Context, userID uuid.UUID, req entity.SupporterzData) {
	targetTable := "supporterz"

	// 各フィールドが空でなければログを記録
	if req.CareerVision != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "career_vision")
	}
	if req.SelfPromotion != "" {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "self_promotion")
	}
	if len(req.Skills) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "skills")
	}
	if len(req.SkillDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "skill_descriptions")
	}
	if len(req.InternExperiences) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "intern_experiences")
	}
	if len(req.InternExperienceDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "intern_experience_descriptions")
	}
	if len(req.Products) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "products")
	}
	if len(req.ProductTechStacks) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "product_tech_stacks")
	}
	if len(req.ProductDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "product_descriptions")
	}
	if len(req.Researches) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "researches")
	}
	if len(req.ResearchDescriptions) > 0 {
		u.lu.LogFieldUpdateWithErrorHandling(userID, targetTable, "research_descriptions")
	}
}
