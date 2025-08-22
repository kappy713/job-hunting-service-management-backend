package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type LevtechRookie struct {
	ID                              uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`
	DesiredJobType                  pq.StringArray `gorm:"type:text[]" json:"desired_job_type"`
	CareerAspiration                pq.StringArray `gorm:"type:text[]" json:"career_aspiration"`
	InterestedTasks                 pq.StringArray `gorm:"type:text[]" json:"interested_tasks"`
	JobRequirements                 pq.StringArray `gorm:"type:text[]" json:"job_requirements"`
	InterestedIndustries            pq.StringArray `gorm:"type:text[]" json:"interested_industries"`
	PreferredCompanySize            pq.StringArray `gorm:"type:text[]" json:"preferred_company_size"`
	InterestedBusinessTypes         pq.StringArray `gorm:"type:text[]" json:"interested_business_types"`
	PreferredWorkLocation           pq.StringArray `gorm:"type:text[]" json:"preferred_work_location"`
	Skills                          pq.StringArray `gorm:"type:text[]" json:"skills"`
	SkillDescriptions               pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`
	Portfolio                       string         `gorm:"size:200" json:"portfolio"`
	PortfolioDescription            string         `gorm:"size:2000" json:"portfolio_description"`
	InternExperiences               pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`
	InternExperienceDescriptions    pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"`
	HackathonExperiences            pq.StringArray `gorm:"type:text[]" json:"hackathon_experiences"`
	HackathonExperienceDescriptions pq.StringArray `gorm:"type:text[]" json:"hackathon_experience_descriptions"`
	Research                        string         `gorm:"size:2000" json:"research"`
	Organization                    string         `gorm:"size:2000" json:"organization"`
	Other                           string         `gorm:"size:2000" json:"other"`
	Certifications                  pq.StringArray `gorm:"type:text[]" json:"certifications"`
	Languages                       pq.StringArray `gorm:"type:text[]" json:"languages"`
	LanguageLevels                  pq.StringArray `gorm:"type:text[]" json:"language_levels"`
}

// リクエスト用の構造体
type LevtechRookieData struct {
	DesiredJobType                  []string `json:"desired_job_type"`
	CareerAspiration                []string `json:"career_aspiration"`
	InterestedTasks                 []string `json:"interested_tasks"`
	JobRequirements                 []string `json:"job_requirements"`
	InterestedIndustries            []string `json:"interested_industries"`
	PreferredCompanySize            []string `json:"preferred_company_size"`
	InterestedBusinessTypes         []string `json:"interested_business_types"`
	PreferredWorkLocation           []string `json:"preferred_work_location"`
	Skills                          []string `json:"skills"`
	SkillDescriptions               []string `json:"skill_descriptions"`
	Portfolio                       string   `json:"portfolio"`
	PortfolioDescription            string   `json:"portfolio_description"`
	InternExperiences               []string `json:"intern_experiences"`
	InternExperienceDescriptions    []string `json:"intern_experience_descriptions"`
	HackathonExperiences            []string `json:"hackathon_experiences"`
	HackathonExperienceDescriptions []string `json:"hackathon_experience_descriptions"`
	Research                        string   `json:"research"`
	Organization                    string   `json:"organization"`
	Other                           string   `json:"other"`
	Certifications                  []string `json:"certifications"`
	Languages                       []string `json:"languages"`
	LanguageLevels                  []string `json:"language_levels"`
}

type CreateLevtechRookieRequest struct {
	UserID string            `json:"user_id" binding:"required"`
	Data   LevtechRookieData `json:"data" binding:"required"`
}

func (LevtechRookie) TableName() string {
	return "levtech_rookie"
}
