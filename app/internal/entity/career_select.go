package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type CareerSelect struct {
	ID                                   uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`
	Skills                               pq.StringArray `gorm:"type:text[]" json:"skills"`
	SkillDescriptions                    pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`
	CompanySelectionCriteria             pq.StringArray `gorm:"type:text[]" json:"company_selection_criteria"`
	CompanySelectionCriteriaDescriptions pq.StringArray `gorm:"type:text[]" json:"company_selection_criteria_descriptions"`
	CareerVision                         string         `gorm:"size:2000" json:"career_vision"`
	SelfPromotion                        string         `gorm:"size:5000" json:"self_promotion"`
	Research                             string         `gorm:"size:500" json:"research"`
	Products                             pq.StringArray `gorm:"type:text[]" json:"products"`
	ProductDescriptions                  pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`
	Experiences                          pq.StringArray `gorm:"type:text[]" json:"experiences"`
	ExperienceDescriptions               pq.StringArray `gorm:"type:text[]" json:"experience_descriptions"`
	InternExperiences                    pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`
	InternExperienceDescriptions         pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"`
	Certifications                       pq.StringArray `gorm:"type:text[]" json:"certifications"`
	CertificationDescriptions            pq.StringArray `gorm:"type:text[]" json:"certification_descriptions"`
}

// リクエスト用の構造体
type CareerSelectData struct {
	Skills                               []string `json:"skills"`
	SkillDescriptions                    []string `json:"skill_descriptions"`
	CompanySelectionCriteria             []string `json:"company_selection_criteria"`
	CompanySelectionCriteriaDescriptions []string `json:"company_selection_criteria_descriptions"`
	CareerVision                         string   `json:"career_vision"`
	SelfPromotion                        string   `json:"self_promotion"`
	Research                             string   `json:"research"`
	Products                             []string `json:"products"`
	ProductDescriptions                  []string `json:"product_descriptions"`
	Experiences                          []string `json:"experiences"`
	ExperienceDescriptions               []string `json:"experience_descriptions"`
	InternExperiences                    []string `json:"intern_experiences"`
	InternExperienceDescriptions         []string `json:"intern_experience_descriptions"`
	Certifications                       []string `json:"certifications"`
	CertificationDescriptions            []string `json:"certification_descriptions"`
}

type CreateCareerSelectRequest struct {
	UserID string           `json:"user_id" binding:"required"`
	Data   CareerSelectData `json:"data" binding:"required"`
}

func (CareerSelect) TableName() string {
	return "career_select"
}
