package entity

import (
	"github.com/google/uuid"
)

// AI生成リクエストの構造体
type AIGenerationRequest struct {
	UserID   uuid.UUID `json:"user_id" binding:"required"`
	Services []string  `json:"services" binding:"required"`
}

// AI生成レスポンスの構造体
type AIGenerationResponse struct {
	UserID  uuid.UUID              `json:"user_id"`
	Results map[string]interface{} `json:"results"`
	Status  string                 `json:"status"`
	Message string                 `json:"message,omitempty"`
}

// プロンプトテンプレートの構造体
type PromptTemplate struct {
	ServiceName string `json:"service_name"`
	Template    string `json:"template"`
}

// 日本語サービス名からアルファベットへの変換マップ
var ServiceNameMap = map[string]string{
	"サポーターズ":    "supporterz",
	"キャリアセレクト":  "career_select",
	"ワンキャリア":    "one_career",
	"レバテックルーキー": "levtech_rookie",
	"マイナビ":      "mynavi",
}

// アルファベットから日本語サービス名への逆変換マップ
var ReverseServiceNameMap = map[string]string{
	"supporterz":     "サポーターズ",
	"career_select":  "キャリアセレクト",
	"one_career":     "ワンキャリア",
	"levtech_rookie": "レバテックルーキー",
	"mynavi":         "マイナビ",
}

// サービス名変換関数
func ConvertServiceName(japaneseName string) (string, bool) {
	englishName, exists := ServiceNameMap[japaneseName]
	return englishName, exists
}

// 逆変換関数
func ConvertServiceNameToJapanese(englishName string) (string, bool) {
	japaneseName, exists := ReverseServiceNameMap[englishName]
	return japaneseName, exists
}

// 各サービス用のプロンプトテンプレート
var ServicePrompts = map[string]string{
	"supporterz": `
あなたはサポーターズの就活支援AIです。以下のユーザー情報に基づいて、サポーターズのプロフィール項目を日本語で生成してください。

ユーザー情報:
- 氏名: {{.LastName}} {{.FirstName}}
- 年齢: {{.Age}}歳
- 大学: {{.University}}
- 学部: {{.Faculty}}
- 学年: {{.Grade}}年
- 志望職種: {{.TargetJobType}}

技術スキルについては、以下の観点から幅広く生成してください:
- プログラミング言語（Java, Python, JavaScript, Go, C++, Swift, Kotlin等）
- フレームワーク・ライブラリ（Spring Boot, React, Vue.js, Django, Express.js等）
- データベース（MySQL, PostgreSQL, MongoDB, Redis等）
- クラウド・インフラ（AWS, GCP, Azure, Docker, Kubernetes等）
- 開発ツール（Git, GitHub, Jenkins, Jira等）
- その他の技術（API設計, テスト自動化, CI/CD等）

JSONオブジェクトのみを返してください。マークダウンのコードブロックは使用せず、純粋なJSONで回答してください（日本語で記述）:
{
  "career_vision": "キャリアビジョンを3-5行で記述",
  "self_promotion": "自己PRを5-7行で記述",
  "skills": ["プログラミング言語1", "プログラミング言語2", "フレームワーク1", "フレームワーク2", "データベース1", "クラウド技術1", "開発ツール1", "その他技術1"],
  "skill_descriptions": ["スキル1の具体的な経験と習熟度", "スキル2の具体的な経験と習熟度", "スキル3の具体的な経験と習熟度", "スキル4の具体的な経験と習熟度", "スキル5の具体的な経験と習熟度", "スキル6の具体的な経験と習熟度", "スキル7の具体的な経験と習熟度", "スキル8の具体的な経験と習熟度"],
  "intern_experiences": ["インターン経験1", "インターン経験2"],
  "intern_experience_descriptions": ["インターン経験1の詳細", "インターン経験2の詳細"],
  "products": ["制作物1", "制作物2"],
  "product_tech_stacks": ["技術スタック1", "技術スタック2"],
  "product_descriptions": ["制作物1の説明", "制作物2の説明"],
  "researches": ["研究テーマ1"],
  "research_descriptions": ["研究テーマ1の詳細説明"]
}`,

	"career_select": `
あなたはキャリアセレクトの就活支援AIです。以下のユーザー情報に基づいて、キャリアセレクトのプロフィール項目を日本語で生成してください。

ユーザー情報:
- 氏名: {{.LastName}} {{.FirstName}}
- 年齢: {{.Age}}歳
- 大学: {{.University}}
- 学部: {{.Faculty}}
- 学年: {{.Grade}}年
- 志望職種: {{.TargetJobType}}

技術スキルについては、プログラミング言語、フレームワーク、データベース、クラウド技術、開発ツールなど幅広く生成してください。

JSONオブジェクトのみを返してください。マークダウンのコードブロックは使用せず、純粋なJSONで回答してください（日本語で記述）:
{
  "skills": ["プログラミング言語1", "プログラミング言語2", "フレームワーク1", "フレームワーク2", "データベース1", "クラウド技術1", "開発ツール1"],
  "skill_descriptions": ["スキル1の具体的な経験と習熟度", "スキル2の具体的な経験と習熟度", "スキル3の具体的な経験と習熟度", "スキル4の具体的な経験と習熟度", "スキル5の具体的な経験と習熟度", "スキル6の具体的な経験と習熟度", "スキル7の具体的な経験と習熟度"],
  "company_selection_criteria": ["企業選択基準1", "企業選択基準2"],
  "company_selection_criteria_descriptions": ["基準1の詳細", "基準2の詳細"],
  "career_vision": "キャリアビジョンを3-5行で記述",
  "self_promotion": "自己PRを5-7行で記述",
  "research": "研究内容の詳細説明",
  "products": ["制作物1", "制作物2"],
  "product_descriptions": ["制作物1の説明", "制作物2の説明"],
  "experiences": ["経験1", "経験2"],
  "experience_descriptions": ["経験1の詳細", "経験2の詳細"],
  "intern_experiences": ["インターン経験1"],
  "intern_experience_descriptions": ["インターン経験1の詳細"],
  "certifications": ["資格1", "資格2"],
  "certification_descriptions": ["資格1の詳細", "資格2の詳細"]
}`,

	"one_career": `
あなたはワンキャリアの就活支援AIです。以下のユーザー情報に基づいて、ワンキャリアのプロフィール項目を日本語で生成してください。

ユーザー情報:
- 氏名: {{.LastName}} {{.FirstName}}
- 年齢: {{.Age}}歳
- 大学: {{.University}}
- 学部: {{.Faculty}}
- 学年: {{.Grade}}年
- 志望職種: {{.TargetJobType}}

技術スキルについては、プログラミング言語、フレームワーク、データベース、クラウド技術、開発ツールなど幅広く生成してください。

JSONオブジェクトのみを返してください。マークダウンのコードブロックは使用せず、純粋なJSONで回答してください（日本語で記述）:
{
  "skills": ["プログラミング言語1", "プログラミング言語2", "フレームワーク1", "フレームワーク2", "データベース1", "開発ツール1"],
  "skill_descriptions": ["スキル1の具体的な経験と習熟度", "スキル2の具体的な経験と習熟度", "スキル3の具体的な経験と習熟度", "スキル4の具体的な経験と習熟度", "スキル5の具体的な経験と習熟度", "スキル6の具体的な経験と習熟度"],
  "researches": ["研究テーマ1"],
  "research_descriptions": ["研究テーマ1の詳細説明"],
  "intern_experiences": ["インターン経験1"],
  "intern_experience_descriptions": ["インターン経験1の詳細"],
  "products": ["制作物1", "制作物2"],
  "product_descriptions": ["制作物1の説明", "制作物2の説明"],
  "engineer_aspiration": "エンジニア志望動機を5-7行で記述"
}`,

	"mynavi": `
あなたはマイナビの就活支援AIです。以下のユーザー情報に基づいて、マイナビのプロフィール項目を日本語で生成してください。

ユーザー情報:
- 氏名: {{.LastName}} {{.FirstName}}
- 年齢: {{.Age}}歳
- 大学: {{.University}}
- 学部: {{.Faculty}}
- 学年: {{.Grade}}年
- 志望職種: {{.TargetJobType}}

JSONオブジェクトのみを返してください。マークダウンのコードブロックは使用せず、純粋なJSONで回答してください（日本語で記述）:
{
  "self_promotion": "自己PRを5-7行で記述",
  "future_plan": "将来のキャリアプランを3-5行で記述"
}`,

	"levtech_rookie": `
あなたはレバテックルーキーの就活支援AIです。以下のユーザー情報に基づいて、レバテックルーキーのプロフィール項目を日本語で生成してください。

ユーザー情報:
- 氏名: {{.LastName}} {{.FirstName}}
- 年齢: {{.Age}}歳
- 大学: {{.University}}
- 学部: {{.Faculty}}
- 学年: {{.Grade}}年
- 志望職種: {{.TargetJobType}}

技術スキルについては、プログラミング言語、フレームワーク、データベース、クラウド技術、開発ツールなど幅広く生成してください。エンジニア向けのサービスなので、技術的なスキルを特に充実させてください。

JSONオブジェクトのみを返してください。マークダウンのコードブロックは使用せず、純粋なJSONで回答してください（日本語で記述）:
{
  "desired_job_type": ["希望職種1", "希望職種2"],
  "career_aspiration": ["キャリア志向1", "キャリア志向2"],
  "interested_tasks": ["興味のある業務1", "興味のある業務2"],
  "job_requirements": ["希望条件1", "希望条件2"],
  "interested_industries": ["興味のある業界1", "興味のある業界2"],
  "preferred_company_size": ["希望企業規模1", "希望企業規模2"],
  "interested_business_types": ["興味のある事業形態1"],
  "preferred_work_location": ["希望勤務地1", "希望勤務地2"],
  "skills": ["プログラミング言語1", "プログラミング言語2", "フレームワーク1", "フレームワーク2", "データベース1", "クラウド技術1", "開発ツール1", "その他技術1"],
  "skill_descriptions": ["スキル1の具体的な経験と習熟度", "スキル2の具体的な経験と習熟度", "スキル3の具体的な経験と習熟度", "スキル4の具体的な経験と習熟度", "スキル5の具体的な経験と習熟度", "スキル6の具体的な経験と習熟度", "スキル7の具体的な経験と習熟度", "スキル8の具体的な経験と習熟度"],
  "portfolio": "ポートフォリオURL（例: https://github.com/username）",
  "portfolio_description": "ポートフォリオの詳細説明",
  "intern_experiences": ["インターン経験1"],
  "intern_experience_descriptions": ["インターン経験1の詳細"],
  "hackathon_experiences": ["ハッカソン経験1"],
  "hackathon_experience_descriptions": ["ハッカソン経験1の詳細"],
  "research": "研究内容の詳細説明",
  "organization": "所属組織・団体活動",
  "other": "その他のアピールポイント",
  "certifications": ["資格1", "資格2"],
  "languages": ["使用言語1", "使用言語2"],
  "language_levels": ["言語1のレベル", "言語2のレベル"]
}`,
}
