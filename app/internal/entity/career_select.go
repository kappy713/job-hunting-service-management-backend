package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// CareerSelect キャリアセレクト用のプロフィール情報
type CareerSelect struct {
	ID                                   uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`                             // ユーザーID（主キー）
	Skills                               pq.StringArray `gorm:"type:text[]" json:"skills"`                                  // 保有スキル一覧
	SkillDescriptions                    pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`                      // 各スキルの詳細説明
	CompanySelectionCriteria             pq.StringArray `gorm:"type:text[]" json:"company_selection_criteria"`              // 会社選びの軸一覧
	CompanySelectionCriteriaDescriptions pq.StringArray `gorm:"type:text[]" json:"company_selection_criteria_descriptions"` // 各会社選びの軸の詳細説明
	CareerVision                         string         `gorm:"size:2000" json:"career_vision"`                             // 将来のキャリアビジョン
	SelfPromotion                        string         `gorm:"size:5000" json:"self_promotion"`                            // 自己PR文
	Research                             string         `gorm:"size:500" json:"research"`                                   // 研究内容
	Products                             pq.StringArray `gorm:"type:text[]" json:"products"`                                // 制作物・プロダクト一覧
	ProductDescriptions                  pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`                    // 各制作物の詳細説明
	Experiences                          pq.StringArray `gorm:"type:text[]" json:"experiences"`                             // その他の経験一覧
	ExperienceDescriptions               pq.StringArray `gorm:"type:text[]" json:"experience_descriptions"`                 // 各経験の詳細説明
	InternExperiences                    pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`                      // インターン経験一覧
	InternExperienceDescriptions         pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"`          // 各インターン経験の詳細説明
	Certifications                       pq.StringArray `gorm:"type:text[]" json:"certifications"`                          // 取得資格一覧
	CertificationDescriptions            pq.StringArray `gorm:"type:text[]" json:"certification_descriptions"`              // 各資格の詳細説明
}

// CareerSelectData キャリアセレクト用のリクエストデータ構造体
type CareerSelectData struct {
	Skills                               []string `json:"skills"`                                  // 保有スキル一覧
	SkillDescriptions                    []string `json:"skill_descriptions"`                      // 各スキルの詳細説明
	CompanySelectionCriteria             []string `json:"company_selection_criteria"`              // 会社選びの軸一覧
	CompanySelectionCriteriaDescriptions []string `json:"company_selection_criteria_descriptions"` // 各会社選びの軸の詳細説明
	CareerVision                         string   `json:"career_vision"`                           // 将来のキャリアビジョン
	SelfPromotion                        string   `json:"self_promotion"`                          // 自己PR文
	Research                             string   `json:"research"`                                // 研究内容
	Products                             []string `json:"products"`                                // 制作物・プロダクト一覧
	ProductDescriptions                  []string `json:"product_descriptions"`                    // 各制作物の詳細説明
	Experiences                          []string `json:"experiences"`                             // その他の経験一覧
	ExperienceDescriptions               []string `json:"experience_descriptions"`                 // 各経験の詳細説明
	InternExperiences                    []string `json:"intern_experiences"`                      // インターン経験一覧
	InternExperienceDescriptions         []string `json:"intern_experience_descriptions"`          // 各インターン経験の詳細説明
	Certifications                       []string `json:"certifications"`                          // 取得資格一覧
	CertificationDescriptions            []string `json:"certification_descriptions"`              // 各資格の詳細説明
}

// CreateCareerSelectRequest キャリアセレクト作成リクエスト
type CreateCareerSelectRequest struct {
	UserID string           `json:"user_id" binding:"required"` // ユーザーID
	Data   CareerSelectData `json:"data" binding:"required"`    // キャリアセレクトデータ
}

func (CareerSelect) TableName() string {
	return "career_select"
}
