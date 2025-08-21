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

func (CareerSelect) TableName() string {
	return "career_select"
}
