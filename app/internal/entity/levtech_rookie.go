package entity

import (
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// LevtechRookie レバテックルーキー用のプロフィール情報
type LevtechRookie struct {
	ID                              uuid.UUID      `gorm:"type:uuid;primarykey" json:"id"`                       // ユーザーID（主キー）
	DesiredJobType                  pq.StringArray `gorm:"type:text[]" json:"desired_job_type"`                  // 希望職種一覧
	CareerAspiration                pq.StringArray `gorm:"type:text[]" json:"career_aspiration"`                 // キャリア志向一覧
	InterestedTasks                 pq.StringArray `gorm:"type:text[]" json:"interested_tasks"`                  // 興味のある業務一覧
	JobRequirements                 pq.StringArray `gorm:"type:text[]" json:"job_requirements"`                  // 求人への要望一覧
	InterestedIndustries            pq.StringArray `gorm:"type:text[]" json:"interested_industries"`             // 興味のある業界一覧
	PreferredCompanySize            pq.StringArray `gorm:"type:text[]" json:"preferred_company_size"`            // 希望会社規模一覧
	InterestedBusinessTypes         pq.StringArray `gorm:"type:text[]" json:"interested_business_types"`         // 興味のある事業形態一覧
	PreferredWorkLocation           pq.StringArray `gorm:"type:text[]" json:"preferred_work_location"`           // 希望勤務地一覧
	Skills                          pq.StringArray `gorm:"type:text[]" json:"skills"`                            // 保有スキル一覧
	SkillDescriptions               pq.StringArray `gorm:"type:text[]" json:"skill_descriptions"`                // 各スキルの詳細説明
	Portfolio                       string         `gorm:"size:200" json:"portfolio"`                            // ポートフォリオURL
	PortfolioDescription            string         `gorm:"size:2000" json:"portfolio_description"`               // ポートフォリオの詳細説明
	InternExperiences               pq.StringArray `gorm:"type:text[]" json:"intern_experiences"`                // インターン経験一覧
	InternExperienceDescriptions    pq.StringArray `gorm:"type:text[]" json:"intern_experience_descriptions"`    // 各インターン経験の詳細説明
	HackathonExperiences            pq.StringArray `gorm:"type:text[]" json:"hackathon_experiences"`             // ハッカソン経験一覧
	HackathonExperienceDescriptions pq.StringArray `gorm:"type:text[]" json:"hackathon_experience_descriptions"` // 各ハッカソン経験の詳細説明
	Research                        string         `gorm:"size:2000" json:"research"`                            // 研究内容
	Organization                    string         `gorm:"size:2000" json:"organization"`                        // 所属組織・団体
	Other                           string         `gorm:"size:2000" json:"other"`                               // その他の活動・経験
	Certifications                  pq.StringArray `gorm:"type:text[]" json:"certifications"`                    // 取得資格一覧
	Languages                       pq.StringArray `gorm:"type:text[]" json:"languages"`                         // 使用可能言語一覧
	LanguageLevels                  pq.StringArray `gorm:"type:text[]" json:"language_levels"`                   // 各言語のレベル一覧
}

// LevtechRookieData レバテックルーキー用のリクエストデータ構造体
type LevtechRookieData struct {
	DesiredJobType                  []string `json:"desired_job_type"`                  // 希望職種一覧
	CareerAspiration                []string `json:"career_aspiration"`                 // キャリア志向一覧
	InterestedTasks                 []string `json:"interested_tasks"`                  // 興味のある業務一覧
	JobRequirements                 []string `json:"job_requirements"`                  // 求人への要望一覧
	InterestedIndustries            []string `json:"interested_industries"`             // 興味のある業界一覧
	PreferredCompanySize            []string `json:"preferred_company_size"`            // 希望会社規模一覧
	InterestedBusinessTypes         []string `json:"interested_business_types"`         // 興味のある事業形態一覧
	PreferredWorkLocation           []string `json:"preferred_work_location"`           // 希望勤務地一覧
	Skills                          []string `json:"skills"`                            // 保有スキル一覧
	SkillDescriptions               []string `json:"skill_descriptions"`                // 各スキルの詳細説明
	Portfolio                       string   `json:"portfolio"`                         // ポートフォリオURL
	PortfolioDescription            string   `json:"portfolio_description"`             // ポートフォリオの詳細説明
	InternExperiences               []string `json:"intern_experiences"`                // インターン経験一覧
	InternExperienceDescriptions    []string `json:"intern_experience_descriptions"`    // 各インターン経験の詳細説明
	HackathonExperiences            []string `json:"hackathon_experiences"`             // ハッカソン経験一覧
	HackathonExperienceDescriptions []string `json:"hackathon_experience_descriptions"` // 各ハッカソン経験の詳細説明
	Research                        string   `json:"research"`                          // 研究内容
	Organization                    string   `json:"organization"`                      // 所属組織・団体
	Other                           string   `json:"other"`                             // その他の活動・経験
	Certifications                  []string `json:"certifications"`                    // 取得資格一覧
	Languages                       []string `json:"languages"`                         // 使用可能言語一覧
	LanguageLevels                  []string `json:"language_levels"`                   // 各言語のレベル一覧
}

// CreateLevtechRookieRequest レバテックルーキー作成リクエスト
type CreateLevtechRookieRequest struct {
	UserID string            `json:"user_id" binding:"required"` // ユーザーID
	Data   LevtechRookieData `json:"data" binding:"required"`    // レバテックルーキーデータ
}

func (LevtechRookie) TableName() string {
	return "levtech_rookie"
}
