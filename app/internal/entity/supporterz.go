package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Supporterz サポーターズ用のプロフィール情報
type Supporterz struct {
	ID                           uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`                    // ユーザーID（主キー）
	CareerVision                 string         `gorm:"size:200" json:"career_vision"`                     // 将来のキャリアビジョン
	SelfPromotion                string         `gorm:"size:5000" json:"self_promotion"`                   // 自己PR文
	Skills                       pq.StringArray `gorm:"type:text[]" json:"skills"`                         // 保有スキル一覧
	SkillDescriptions            pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`             // 各スキルの詳細説明
	InternExperiences            pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`             // インターン経験一覧
	InternExperienceDescriptions pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"` // 各インターン経験の詳細説明
	Products                     pq.StringArray `gorm:"type:text[]" json:"products"`                       // 制作物・プロダクト一覧
	ProductTechStacks            pq.StringArray `gorm:"type:text[]" json:"product_tech_stacks"`            // 各制作物の技術スタック
	ProductDescriptions          pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`           // 各制作物の詳細説明
	Researches                   pq.StringArray `gorm:"type:text[]" json:"researches"`                     // 研究一覧
	ResearchDescriptions         pq.StringArray `gorm:"type:text[]" json:"research_descriptions"`          // 各研究の詳細説明
}

// SupporterzData サポーターズ用のリクエストデータ構造体
type SupporterzData struct {
	CareerVision                 string   `json:"career_vision"`                  // 将来のキャリアビジョン
	SelfPromotion                string   `json:"self_promotion"`                 // 自己PR文
	Skills                       []string `json:"skills"`                         // 保有スキル一覧
	SkillDescriptions            []string `json:"skill_descriptions"`             // 各スキルの詳細説明
	InternExperiences            []string `json:"intern_experiences"`             // インターン経験一覧
	InternExperienceDescriptions []string `json:"intern_experience_descriptions"` // 各インターン経験の詳細説明
	Products                     []string `json:"products"`                       // 制作物・プロダクト一覧
	ProductTechStacks            []string `json:"product_tech_stacks"`            // 各制作物の技術スタック
	ProductDescriptions          []string `json:"product_descriptions"`           // 各制作物の詳細説明
	Researches                   []string `json:"researches"`                     // 研究一覧
	ResearchDescriptions         []string `json:"research_descriptions"`          // 各研究の詳細説明
}

// CreateSupporterzRequest サポーターズ作成リクエスト
type CreateSupporterzRequest struct {
	UserID string         `json:"user_id" binding:"required"` // ユーザーID
	Data   SupporterzData `json:"data" binding:"required"`    // サポーターズデータ
}

func (Supporterz) TableName() string {
	return "supporterz"
}
