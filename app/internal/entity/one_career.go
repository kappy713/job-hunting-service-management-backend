package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// OneCareer ワンキャリア用のプロフィール情報
type OneCareer struct {
	ID                           uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`                    // ユーザーID（主キー）
	Skills                       pq.StringArray `gorm:"type:text[]" json:"skills"`                         // 保有スキル一覧
	SkillDescriptions            pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`             // 各スキルの詳細説明
	Researches                   pq.StringArray `gorm:"type:text[]" json:"researches"`                     // 研究一覧
	ResearchDescriptions         pq.StringArray `gorm:"type:text[]" json:"research_descriptions"`          // 各研究の詳細説明
	InternExperiences            pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`             // インターン経験一覧
	InternExperienceDescriptions pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"` // 各インターン経験の詳細説明
	Products                     pq.StringArray `gorm:"type:text[]" json:"products"`                       // 制作物・プロダクト一覧
	ProductDescriptions          pq.StringArray `gorm:"type:text[]" json:"product_descriptions"`           // 各制作物の詳細説明
	EngineerAspiration           string         `gorm:"size:1000" json:"engineer_aspiration"`              // エンジニアとしての志望動機
}

// OneCareerData ワンキャリア用のリクエストデータ構造体
type OneCareerData struct {
	Skills                       []string `json:"skills"`                         // 保有スキル一覧
	SkillDescriptions            []string `json:"skill_descriptions"`             // 各スキルの詳細説明
	Researches                   []string `json:"researches"`                     // 研究一覧
	ResearchDescriptions         []string `json:"research_descriptions"`          // 各研究の詳細説明
	InternExperiences            []string `json:"intern_experiences"`             // インターン経験一覧
	InternExperienceDescriptions []string `json:"intern_experience_descriptions"` // 各インターン経験の詳細説明
	Products                     []string `json:"products"`                       // 制作物・プロダクト一覧
	ProductDescriptions          []string `json:"product_descriptions"`           // 各制作物の詳細説明
	EngineerAspiration           string   `json:"engineer_aspiration"`            // エンジニアとしての志望動機
}

// CreateOneCareerRequest ワンキャリア作成リクエスト
type CreateOneCareerRequest struct {
	UserID string        `json:"user_id" binding:"required"` // ユーザーID
	Data   OneCareerData `json:"data" binding:"required"`    // ワンキャリアデータ
}

func (OneCareer) TableName() string {
	return "one_career"
}
