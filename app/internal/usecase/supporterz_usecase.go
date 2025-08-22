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
}

func NewSupporterzUsecase(r repository.SupporterzRepository) SupporterzUsecase {
	return &supporterzUsecase{sr: r}
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

	return u.sr.CreateOrUpdateSupporterz(c, supporterz)
}
