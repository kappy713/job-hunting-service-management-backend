package entity

import "github.com/google/uuid"

type Profile struct {
	ID                       uuid.UUID `gorm:"type:uuid;primarykey" json:"id"`
	CareerVision             string    `gorm:"size:2000" json:"career_vision"`              // キャリアビジョン
	SelfPromotion            string    `gorm:"size:5000" json:"self_promotion"`             // 自己PR
	StudentExperience        string    `gorm:"size:5000" json:"student_experience"`         // ガクチカ
	Research                 string    `gorm:"size:2000" json:"research"`                   // 研究内容
	Product                  string    `gorm:"size:5000" json:"product"`                    // 製作物・開発経験
	Skill                    string    `gorm:"size:5000" json:"skill"`                      // スキル(プログラミング言語・フレームワーク等)
	Intern                   string    `gorm:"size:5000" json:"intern"`                     // インターン・アルバイト経験
	Organization             string    `gorm:"size:2000" json:"organization"`               // 部活・サークル・団体活動経験
	Certification            string    `gorm:"size:2000" json:"certification"`              // 資格
	DesiredJobType           string    `gorm:"size:2000" json:"desired_job_type"`           // 希望職種
	CompanySelectionCriteria string    `gorm:"size:2000" json:"company_selection_criteria"` // 企業選びの軸
	EngineerAspiration       string    `gorm:"size:2000" json:"engineer_aspiration"`        // 理想のエンジニア像
}

func (Profile) TableName() string {
	return "profiles"
}
