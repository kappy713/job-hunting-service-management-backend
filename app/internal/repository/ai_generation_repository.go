package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

// ヘルパー関数: string値の取得とマッピング
func mapStringField(data map[string]interface{}, key string, target *string) {
	if val, ok := data[key].(string); ok {
		*target = val
	}
}

// ヘルパー関数: pq.StringArray型の取得とマッピング
func mapPQStringArrayField(data map[string]interface{}, key string, target *pq.StringArray) {
	if val, ok := data[key].([]interface{}); ok {
		result := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				result = append(result, str)
			}
		}
		*target = result
	}
}

type AIGenerationRepository interface {
	GetUserByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
	SaveSupporterzData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error
	SaveCareerSelectData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error
	SaveOneCareerData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error
	SaveMynaviData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error
	SaveLevtechRookieData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error
}

type aiGenerationRepository struct {
	db *gorm.DB
}

func NewAIGenerationRepository(db *gorm.DB) AIGenerationRepository {
	return &aiGenerationRepository{db: db}
}

func (r *aiGenerationRepository) GetUserByID(ctx context.Context, userID uuid.UUID) (*entity.User, error) {
	var user entity.User
	if err := r.db.WithContext(ctx).Where("user_id = ?", userID).First(&user).Error; err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}
	return &user, nil
}

// Supporterz用のマッピングヘルパー関数
func (r *aiGenerationRepository) mapSupporterzFields(data map[string]interface{}, supporterzData *entity.Supporterz) {
	mapStringField(data, "career_vision", &supporterzData.CareerVision)
	mapStringField(data, "self_promotion", &supporterzData.SelfPromotion)

	mapPQStringArrayField(data, "skills", &supporterzData.Skills)
	mapPQStringArrayField(data, "skill_descriptions", &supporterzData.SkillDescriptions)
	mapPQStringArrayField(data, "intern_experiences", &supporterzData.InternExperiences)
	mapPQStringArrayField(data, "intern_experience_descriptions", &supporterzData.InternExperienceDescriptions)
	mapPQStringArrayField(data, "products", &supporterzData.Products)
	mapPQStringArrayField(data, "product_tech_stacks", &supporterzData.ProductTechStacks)
	mapPQStringArrayField(data, "product_descriptions", &supporterzData.ProductDescriptions)
	mapPQStringArrayField(data, "researches", &supporterzData.Researches)
	mapPQStringArrayField(data, "research_descriptions", &supporterzData.ResearchDescriptions)
}

func (r *aiGenerationRepository) SaveSupporterzData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	supporterzData := &entity.Supporterz{
		ID: userID, // user_idを直接IDとして使用
	}

	r.mapSupporterzFields(data, supporterzData)

	if err := r.db.WithContext(ctx).Save(supporterzData).Error; err != nil {
		return fmt.Errorf("failed to save supporterz data: %w", err)
	}
	return nil
}

// CareerSelect用のマッピングヘルパー関数
func (r *aiGenerationRepository) mapCareerSelectFields(data map[string]interface{}, careerSelectData *entity.CareerSelect) {
	mapStringField(data, "career_vision", &careerSelectData.CareerVision)
	mapStringField(data, "self_promotion", &careerSelectData.SelfPromotion)
	mapStringField(data, "research", &careerSelectData.Research)

	mapPQStringArrayField(data, "skills", &careerSelectData.Skills)
	mapPQStringArrayField(data, "skill_descriptions", &careerSelectData.SkillDescriptions)
	mapPQStringArrayField(data, "company_selection_criteria", &careerSelectData.CompanySelectionCriteria)
	mapPQStringArrayField(data, "company_selection_criteria_descriptions", &careerSelectData.CompanySelectionCriteriaDescriptions)
	mapPQStringArrayField(data, "products", &careerSelectData.Products)
	mapPQStringArrayField(data, "product_descriptions", &careerSelectData.ProductDescriptions)
	mapPQStringArrayField(data, "experiences", &careerSelectData.Experiences)
	mapPQStringArrayField(data, "experience_descriptions", &careerSelectData.ExperienceDescriptions)
	mapPQStringArrayField(data, "intern_experiences", &careerSelectData.InternExperiences)
	mapPQStringArrayField(data, "intern_experience_descriptions", &careerSelectData.InternExperienceDescriptions)
	mapPQStringArrayField(data, "certifications", &careerSelectData.Certifications)
	mapPQStringArrayField(data, "certification_descriptions", &careerSelectData.CertificationDescriptions)
}

func (r *aiGenerationRepository) SaveCareerSelectData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	careerSelectData := &entity.CareerSelect{
		ID: userID, // user_idを直接IDとして使用
	}

	r.mapCareerSelectFields(data, careerSelectData)

	if err := r.db.WithContext(ctx).Save(careerSelectData).Error; err != nil {
		return fmt.Errorf("failed to save career select data: %w", err)
	}
	return nil
}

// OneCareer用のマッピングヘルパー関数
func (r *aiGenerationRepository) mapOneCareerFields(data map[string]interface{}, oneCareerData *entity.OneCareer) {
	mapStringField(data, "engineer_aspiration", &oneCareerData.EngineerAspiration)

	mapPQStringArrayField(data, "skills", &oneCareerData.Skills)
	mapPQStringArrayField(data, "skill_descriptions", &oneCareerData.SkillDescriptions)
	mapPQStringArrayField(data, "researches", &oneCareerData.Researches)
	mapPQStringArrayField(data, "research_descriptions", &oneCareerData.ResearchDescriptions)
	mapPQStringArrayField(data, "intern_experiences", &oneCareerData.InternExperiences)
	mapPQStringArrayField(data, "intern_experience_descriptions", &oneCareerData.InternExperienceDescriptions)
	mapPQStringArrayField(data, "products", &oneCareerData.Products)
	mapPQStringArrayField(data, "product_descriptions", &oneCareerData.ProductDescriptions)
}

func (r *aiGenerationRepository) SaveOneCareerData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	oneCareerData := &entity.OneCareer{
		ID: userID, // user_idを直接IDとして使用
	}

	r.mapOneCareerFields(data, oneCareerData)

	if err := r.db.WithContext(ctx).Save(oneCareerData).Error; err != nil {
		return fmt.Errorf("failed to save one career data: %w", err)
	}
	return nil
}

func (r *aiGenerationRepository) SaveMynaviData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	mynaviData := &entity.Mynavi{
		ID: userID, // user_idを直接IDとして使用
	}

	// データマッピング
	if val, ok := data["self_promotion"].(string); ok {
		mynaviData.SelfPromotion = val
	}
	if val, ok := data["future_plan"].(string); ok {
		mynaviData.FuturePlan = val
	}

	if err := r.db.WithContext(ctx).Save(mynaviData).Error; err != nil {
		return fmt.Errorf("failed to save mynavi data: %w", err)
	}
	return nil
}

// LevtechRookie用のマッピングヘルパー関数
func (r *aiGenerationRepository) mapLevtechRookieStringFields(data map[string]interface{}, levtechData *entity.LevtechRookie) {
	mapStringField(data, "portfolio", &levtechData.Portfolio)
	mapStringField(data, "portfolio_description", &levtechData.PortfolioDescription)
	mapStringField(data, "research", &levtechData.Research)
	mapStringField(data, "organization", &levtechData.Organization)
	mapStringField(data, "other", &levtechData.Other)
}

func (r *aiGenerationRepository) mapLevtechRookieArrayFields(data map[string]interface{}, levtechData *entity.LevtechRookie) {
	mapPQStringArrayField(data, "desired_job_type", &levtechData.DesiredJobType)
	mapPQStringArrayField(data, "career_aspiration", &levtechData.CareerAspiration)
	mapPQStringArrayField(data, "interested_tasks", &levtechData.InterestedTasks)
	mapPQStringArrayField(data, "job_requirements", &levtechData.JobRequirements)
	mapPQStringArrayField(data, "interested_industries", &levtechData.InterestedIndustries)
	mapPQStringArrayField(data, "preferred_company_size", &levtechData.PreferredCompanySize)
	mapPQStringArrayField(data, "interested_business_types", &levtechData.InterestedBusinessTypes)
	mapPQStringArrayField(data, "preferred_work_location", &levtechData.PreferredWorkLocation)
	mapPQStringArrayField(data, "skills", &levtechData.Skills)
	mapPQStringArrayField(data, "skill_descriptions", &levtechData.SkillDescriptions)
	mapPQStringArrayField(data, "intern_experiences", &levtechData.InternExperiences)
	mapPQStringArrayField(data, "intern_experience_descriptions", &levtechData.InternExperienceDescriptions)
	mapPQStringArrayField(data, "hackathon_experiences", &levtechData.HackathonExperiences)
	mapPQStringArrayField(data, "hackathon_experience_descriptions", &levtechData.HackathonExperienceDescriptions)
	mapPQStringArrayField(data, "certifications", &levtechData.Certifications)
	mapPQStringArrayField(data, "languages", &levtechData.Languages)
	mapPQStringArrayField(data, "language_levels", &levtechData.LanguageLevels)
}

func (r *aiGenerationRepository) SaveLevtechRookieData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	levtechData := &entity.LevtechRookie{
		ID: userID, // user_idを直接IDとして使用
	}

	r.mapLevtechRookieStringFields(data, levtechData)
	r.mapLevtechRookieArrayFields(data, levtechData)

	if err := r.db.WithContext(ctx).Save(levtechData).Error; err != nil {
		return fmt.Errorf("failed to save levtech rookie data: %w", err)
	}
	return nil
}
