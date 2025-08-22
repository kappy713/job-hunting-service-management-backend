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
	lu  LogUsecase
}

func NewOneCareerUsecase(r repository.OneCareerRepository, l LogUsecase) OneCareerUsecase {
	return &oneCareerUsecase{ocr: r, lu: l}
}

func (u *oneCareerUsecase) GetOneCareerByUserID(c *gin.Context, userID uuid.UUID) (*entity.OneCareer, error) {
	oneCareer, err := u.ocr.GetOneCareerByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return oneCareer, nil
}

func (u *oneCareerUsecase) CreateOrUpdateOneCareer(c *gin.Context, userID uuid.UUID, req entity.OneCareerData) (*entity.OneCareer, error) {
	oneCareer := &entity.OneCareer{
		ID:                           userID,
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

	result, err := u.ocr.CreateOrUpdateOneCareer(c, oneCareer)
	if err != nil {
		return nil, err
	}

	// 更新されたフィールドのログを記録
	u.logFieldUpdates(c, userID, req)

	return result, nil
}

// フィールド更新のログを記録
func (u *oneCareerUsecase) logFieldUpdates(c *gin.Context, userID uuid.UUID, req entity.OneCareerData) {
	targetTable := "one_career"

	// 各フィールドが空でなければログを記録
	if len(req.Skills) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "skills")
	}
	if len(req.SkillDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "skill_descriptions")
	}
	if len(req.Researches) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "researches")
	}
	if len(req.ResearchDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "research_descriptions")
	}
	if len(req.InternExperiences) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "intern_experiences")
	}
	if len(req.InternExperienceDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "intern_experience_descriptions")
	}
	if len(req.Products) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "products")
	}
	if len(req.ProductDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "product_descriptions")
	}
	if req.EngineerAspiration != "" {
		u.lu.UpsertLog(c, userID, targetTable, "engineer_aspiration")
	}
}
