package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type OneCareerUsecase interface {
	GetOneCareerByUserID(c *gin.Context, userID uuid.UUID) (*entity.OneCareer, error)
	CreateOrUpdateOneCareer(c *gin.Context, userID uuid.UUID, req entity.OneCareerData) (*entity.OneCareer, error)
}

type oneCareerUsecase struct {
	ocr repository.OneCareerRepository
}

func NewOneCareerUsecase(r repository.OneCareerRepository) OneCareerUsecase {
	return &oneCareerUsecase{ocr: r}
}

func (u *oneCareerUsecase) GetOneCareerByUserID(c *gin.Context, userID uuid.UUID) (*entity.OneCareer, error) {
	oneCareer, err := u.ocr.GetOneCareerByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return oneCareer, nil
}

func (u *oneCareerUsecase) CreateOrUpdateOneCareer(c *gin.Context, userID uuid.UUID, req entity.OneCareerData) (*entity.OneCareer, error) {
	// リクエストから entity に変換
	oneCareer := &entity.OneCareer{
		ID:                           userID, // IDにuser_idを設定
		Skills:                       pq.StringArray(req.Skills),
		SkillDescriptions:            pq.StringArray(req.SkillDescriptions),
		Researches:                   pq.StringArray(req.Researches),
		ResearchDescriptions:         pq.StringArray(req.ResearchDescriptions),
		InternExperiences:            pq.StringArray(req.InternExperiences),
		InternExperienceDescriptions: pq.StringArray(req.InternExperienceDescriptions),
		Products:                     pq.StringArray(req.Products),
		ProductDescriptions:          pq.StringArray(req.ProductDescriptions),
		EngineerAspiration:           req.EngineerAspiration,
	}

	return u.ocr.CreateOrUpdateOneCareer(c, oneCareer)
}
