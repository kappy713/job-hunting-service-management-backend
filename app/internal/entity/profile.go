package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Profile struct {
	ID                        uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`
	CareerVision              string         `gorm:"size:2000" json:"career_vision"`                // キャリアビジョン
	SelfPromotion             string         `gorm:"size:5000" json:"self_promotion"`               // 自己PR
	StudentExperience         string         `gorm:"size:5000" json:"student_experience"`           // ガクチカ
	Research                  string         `gorm:"size:2000" json:"research"`                     // 研究内容
	Products                  pq.StringArray `gorm:"type:text[]" json:"products"`                   // 製作物・開発経験（配列）
	ProductDescriptions       pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`       // 製作物説明（配列）
	Skills                    pq.StringArray `gorm:"type:text[]" json:"skills"`                     // スキル（配列）
	SkillDescriptions         pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`         // スキル説明（配列）
	Interns                   pq.StringArray `gorm:"type:text[]" json:"interns"`                    // インターン・アルバイト経験（配列）
	InternDescriptions        pq.StringArray `gorm:"type:text[]" json:"intern_descriptions"`        // インターン説明（配列）
	Organization              string         `gorm:"size:2000" json:"organization"`                 // 部活・サークル・団体活動経験
	Certifications            pq.StringArray `gorm:"type:text[]" json:"certifications"`             // 資格（配列）
	CertificationDescriptions pq.StringArray `gorm:"type:text[]" json:"certification_descriptions"` // 資格説明（配列）
	DesiredJobType            string         `gorm:"size:2000" json:"desired_job_type"`             // 希望職種
	CompanySelectionCriteria  string         `gorm:"size:2000" json:"company_selection_criteria"`   // 企業選びの軸
	EngineerAspiration        string         `gorm:"size:2000" json:"engineer_aspiration"`          // 理想のエンジニア像
}

func (Profile) TableName() string {
	return "profiles"
}
