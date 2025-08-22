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
}

func NewCareerSelectUsecase(r repository.CareerSelectRepository) CareerSelectUsecase {
	return &careerSelectUsecase{csr: r}
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

	return u.csr.CreateOrUpdateCareerSelect(c, careerSelect)
}
