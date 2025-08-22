package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type CareerSelectUsecase interface {
	GetCareerSelectByUserID(c *gin.Context, userID uuid.UUID) (*entity.CareerSelect, error)
	CreateOrUpdateCareerSelect(c *gin.Context, userID uuid.UUID, req entity.CareerSelectData) (*entity.CareerSelect, error)
}

type careerSelectUsecase struct {
	csr repository.CareerSelectRepository
	lu  LogUsecase
}

func NewCareerSelectUsecase(r repository.CareerSelectRepository, l LogUsecase) CareerSelectUsecase {
	return &careerSelectUsecase{csr: r, lu: l}
}

func (u *careerSelectUsecase) GetCareerSelectByUserID(c *gin.Context, userID uuid.UUID) (*entity.CareerSelect, error) {
	careerSelect, err := u.csr.GetCareerSelectByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return careerSelect, nil
}

func (u *careerSelectUsecase) CreateOrUpdateCareerSelect(c *gin.Context, userID uuid.UUID, req entity.CareerSelectData) (*entity.CareerSelect, error) {
	careerSelect := &entity.CareerSelect{
		ID:                                   userID,
		Skills:                               pq.StringArray(req.Skills),
		SkillDescriptions:                    pq.StringArray(req.SkillDescriptions),
		CompanySelectionCriteria:             pq.StringArray(req.CompanySelectionCriteria),
		CompanySelectionCriteriaDescriptions: pq.StringArray(req.CompanySelectionCriteriaDescriptions),
		CareerVision:                         req.CareerVision,
		SelfPromotion:                        req.SelfPromotion,
		Research:                             req.Research,
		Products:                             pq.StringArray(req.Products),
		ProductDescriptions:                  pq.StringArray(req.ProductDescriptions),
		Experiences:                          pq.StringArray(req.Experiences),
		ExperienceDescriptions:               pq.StringArray(req.ExperienceDescriptions),
		InternExperiences:                    pq.StringArray(req.InternExperiences),
		InternExperienceDescriptions:         pq.StringArray(req.InternExperienceDescriptions),
		Certifications:                       pq.StringArray(req.Certifications),
		CertificationDescriptions:            pq.StringArray(req.CertificationDescriptions),
	}

	result, err := u.csr.CreateOrUpdateCareerSelect(c, careerSelect)
	if err != nil {
		return nil, err
	}

	// 更新されたフィールドのログを記録
	u.logFieldUpdates(c, userID, req)

	return result, nil
}

// フィールド更新のログを記録
func (u *careerSelectUsecase) logFieldUpdates(c *gin.Context, userID uuid.UUID, req entity.CareerSelectData) {
	targetTable := "career_select"

	// 各フィールドが空でなければログを記録
	if len(req.Skills) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "skills")
	}
	if len(req.SkillDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "skill_descriptions")
	}
	if len(req.CompanySelectionCriteria) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "company_selection_criteria")
	}
	if len(req.CompanySelectionCriteriaDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "company_selection_criteria_descriptions")
	}
	if req.CareerVision != "" {
		u.lu.UpsertLog(c, userID, targetTable, "career_vision")
	}
	if req.SelfPromotion != "" {
		u.lu.UpsertLog(c, userID, targetTable, "self_promotion")
	}
	if req.Research != "" {
		u.lu.UpsertLog(c, userID, targetTable, "research")
	}
	if len(req.Products) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "products")
	}
	if len(req.ProductDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "product_descriptions")
	}
	if len(req.Experiences) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "experiences")
	}
	if len(req.ExperienceDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "experience_descriptions")
	}
	if len(req.InternExperiences) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "intern_experiences")
	}
	if len(req.InternExperienceDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "intern_experience_descriptions")
	}
	if len(req.Certifications) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "certifications")
	}
	if len(req.CertificationDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "certification_descriptions")
	}
}
