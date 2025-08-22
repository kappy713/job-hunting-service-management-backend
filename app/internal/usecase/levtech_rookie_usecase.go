package usecase

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/lib/pq"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/repository"
)

type LevtechRookieUsecase interface {
	GetLevtechRookieByUserID(c *gin.Context, userID uuid.UUID) (*entity.LevtechRookie, error)
	CreateOrUpdateLevtechRookie(c *gin.Context, userID uuid.UUID, req entity.LevtechRookieData) (*entity.LevtechRookie, error)
}

type levtechRookieUsecase struct {
	lrr repository.LevtechRookieRepository
}

func NewLevtechRookieUsecase(r repository.LevtechRookieRepository) LevtechRookieUsecase {
	return &levtechRookieUsecase{lrr: r}
}

func (u *levtechRookieUsecase) GetLevtechRookieByUserID(c *gin.Context, userID uuid.UUID) (*entity.LevtechRookie, error) {
	levtechRookie, err := u.lrr.GetLevtechRookieByUserID(c, userID)
	if err != nil {
		return nil, err
	}

	return levtechRookie, nil
}

func (u *levtechRookieUsecase) CreateOrUpdateLevtechRookie(c *gin.Context, userID uuid.UUID, req entity.LevtechRookieData) (*entity.LevtechRookie, error) {
	levtechRookie := &entity.LevtechRookie{
		ID:                              userID,
		DesiredJobType:                  pq.StringArray(req.DesiredJobType),
		CareerAspiration:                pq.StringArray(req.CareerAspiration),
		InterestedTasks:                 pq.StringArray(req.InterestedTasks),
		JobRequirements:                 pq.StringArray(req.JobRequirements),
		InterestedIndustries:            pq.StringArray(req.InterestedIndustries),
		PreferredCompanySize:            pq.StringArray(req.PreferredCompanySize),
		InterestedBusinessTypes:         pq.StringArray(req.InterestedBusinessTypes),
		PreferredWorkLocation:           pq.StringArray(req.PreferredWorkLocation),
		Skills:                          pq.StringArray(req.Skills),
		SkillDescriptions:               pq.StringArray(req.SkillDescriptions),
		Portfolio:                       req.Portfolio,
		PortfolioDescription:            req.PortfolioDescription,
		InternExperiences:               pq.StringArray(req.InternExperiences),
		InternExperienceDescriptions:    pq.StringArray(req.InternExperienceDescriptions),
		HackathonExperiences:            pq.StringArray(req.HackathonExperiences),
		HackathonExperienceDescriptions: pq.StringArray(req.HackathonExperienceDescriptions),
		Research:                        req.Research,
		Organization:                    req.Organization,
		Other:                           req.Other,
		Certifications:                  pq.StringArray(req.Certifications),
		Languages:                       pq.StringArray(req.Languages),
		LanguageLevels:                  pq.StringArray(req.LanguageLevels),
	}

	return u.lrr.CreateOrUpdateLevtechRookie(c, levtechRookie)
}
