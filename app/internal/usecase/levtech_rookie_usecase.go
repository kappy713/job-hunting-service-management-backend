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
	lu  LogUsecase
}

func NewLevtechRookieUsecase(r repository.LevtechRookieRepository, l LogUsecase) LevtechRookieUsecase {
	return &levtechRookieUsecase{lrr: r, lu: l}
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

	result, err := u.lrr.CreateOrUpdateLevtechRookie(c, levtechRookie)
	if err != nil {
		return nil, err
	}

	// 更新されたフィールドのログを記録
	u.logFieldUpdates(c, userID, req)

	return result, nil
}

// フィールド更新のログを記録
func (u *levtechRookieUsecase) logFieldUpdates(c *gin.Context, userID uuid.UUID, req entity.LevtechRookieData) {
	targetTable := "levtech_rookie"

	// 各フィールドが空でなければログを記録
	if len(req.DesiredJobType) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "desired_job_type")
	}
	if len(req.CareerAspiration) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "career_aspiration")
	}
	if len(req.InterestedTasks) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "interested_tasks")
	}
	if len(req.JobRequirements) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "job_requirements")
	}
	if len(req.InterestedIndustries) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "interested_industries")
	}
	if len(req.PreferredCompanySize) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "preferred_company_size")
	}
	if len(req.InterestedBusinessTypes) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "interested_business_types")
	}
	if len(req.PreferredWorkLocation) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "preferred_work_location")
	}
	if len(req.Skills) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "skills")
	}
	if len(req.SkillDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "skill_descriptions")
	}
	if req.Portfolio != "" {
		u.lu.UpsertLog(c, userID, targetTable, "portfolio")
	}
	if req.PortfolioDescription != "" {
		u.lu.UpsertLog(c, userID, targetTable, "portfolio_description")
	}
	if len(req.InternExperiences) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "intern_experiences")
	}
	if len(req.InternExperienceDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "intern_experience_descriptions")
	}
	if len(req.HackathonExperiences) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "hackathon_experiences")
	}
	if len(req.HackathonExperienceDescriptions) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "hackathon_experience_descriptions")
	}
	if req.Research != "" {
		u.lu.UpsertLog(c, userID, targetTable, "research")
	}
	if req.Organization != "" {
		u.lu.UpsertLog(c, userID, targetTable, "organization")
	}
	if req.Other != "" {
		u.lu.UpsertLog(c, userID, targetTable, "other")
	}
	if len(req.Certifications) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "certifications")
	}
	if len(req.Languages) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "languages")
	}
	if len(req.LanguageLevels) > 0 {
		u.lu.UpsertLog(c, userID, targetTable, "language_levels")
	}
}
