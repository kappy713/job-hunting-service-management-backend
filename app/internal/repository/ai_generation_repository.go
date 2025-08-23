package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"job-hunting-service-management-backend/app/internal/entity"
)

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

func (r *aiGenerationRepository) SaveSupporterzData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	supporterzData := &entity.Supporterz{
		ID: userID, // user_idを直接IDとして使用
	}

	// データマッピング
	if val, ok := data["career_vision"].(string); ok {
		supporterzData.CareerVision = val
	}
	if val, ok := data["self_promotion"].(string); ok {
		supporterzData.SelfPromotion = val
	}
	if val, ok := data["skills"].([]interface{}); ok {
		skills := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				skills = append(skills, str)
			}
		}
		supporterzData.Skills = skills
	}
	if val, ok := data["skill_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		supporterzData.SkillDescriptions = descriptions
	}
	if val, ok := data["intern_experiences"].([]interface{}); ok {
		experiences := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				experiences = append(experiences, str)
			}
		}
		supporterzData.InternExperiences = experiences
	}
	if val, ok := data["intern_experience_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		supporterzData.InternExperienceDescriptions = descriptions
	}
	if val, ok := data["products"].([]interface{}); ok {
		products := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				products = append(products, str)
			}
		}
		supporterzData.Products = products
	}
	if val, ok := data["product_tech_stacks"].([]interface{}); ok {
		stacks := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				stacks = append(stacks, str)
			}
		}
		supporterzData.ProductTechStacks = stacks
	}
	if val, ok := data["product_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		supporterzData.ProductDescriptions = descriptions
	}
	if val, ok := data["researches"].([]interface{}); ok {
		researches := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				researches = append(researches, str)
			}
		}
		supporterzData.Researches = researches
	}
	if val, ok := data["research_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		supporterzData.ResearchDescriptions = descriptions
	}

	if err := r.db.WithContext(ctx).Save(supporterzData).Error; err != nil {
		return fmt.Errorf("failed to save supporterz data: %w", err)
	}
	return nil
}

func (r *aiGenerationRepository) SaveCareerSelectData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	careerSelectData := &entity.CareerSelect{
		ID: userID, // user_idを直接IDとして使用
	}

	// 基本的なデータマッピング
	if val, ok := data["career_vision"].(string); ok {
		careerSelectData.CareerVision = val
	}
	if val, ok := data["self_promotion"].(string); ok {
		careerSelectData.SelfPromotion = val
	}
	if val, ok := data["research"].(string); ok {
		careerSelectData.Research = val
	}

	// スライス系フィールドのマッピング
	if val, ok := data["skills"].([]interface{}); ok {
		skills := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				skills = append(skills, str)
			}
		}
		careerSelectData.Skills = skills
	}
	if val, ok := data["skill_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		careerSelectData.SkillDescriptions = descriptions
	}
	if val, ok := data["company_selection_criteria"].([]interface{}); ok {
		criteria := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				criteria = append(criteria, str)
			}
		}
		careerSelectData.CompanySelectionCriteria = criteria
	}
	if val, ok := data["company_selection_criteria_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		careerSelectData.CompanySelectionCriteriaDescriptions = descriptions
	}
	if val, ok := data["products"].([]interface{}); ok {
		products := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				products = append(products, str)
			}
		}
		careerSelectData.Products = products
	}
	if val, ok := data["product_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		careerSelectData.ProductDescriptions = descriptions
	}
	if val, ok := data["experiences"].([]interface{}); ok {
		experiences := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				experiences = append(experiences, str)
			}
		}
		careerSelectData.Experiences = experiences
	}
	if val, ok := data["experience_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		careerSelectData.ExperienceDescriptions = descriptions
	}
	if val, ok := data["intern_experiences"].([]interface{}); ok {
		experiences := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				experiences = append(experiences, str)
			}
		}
		careerSelectData.InternExperiences = experiences
	}
	if val, ok := data["intern_experience_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		careerSelectData.InternExperienceDescriptions = descriptions
	}
	if val, ok := data["certifications"].([]interface{}); ok {
		certifications := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				certifications = append(certifications, str)
			}
		}
		careerSelectData.Certifications = certifications
	}
	if val, ok := data["certification_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		careerSelectData.CertificationDescriptions = descriptions
	}

	if err := r.db.WithContext(ctx).Save(careerSelectData).Error; err != nil {
		return fmt.Errorf("failed to save career select data: %w", err)
	}
	return nil
}

func (r *aiGenerationRepository) SaveOneCareerData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	oneCareerData := &entity.OneCareer{
		ID: userID, // user_idを直接IDとして使用
	}

	// データマッピング
	if val, ok := data["engineer_aspiration"].(string); ok {
		oneCareerData.EngineerAspiration = val
	}
	if val, ok := data["skills"].([]interface{}); ok {
		skills := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				skills = append(skills, str)
			}
		}
		oneCareerData.Skills = skills
	}
	if val, ok := data["skill_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		oneCareerData.SkillDescriptions = descriptions
	}
	if val, ok := data["researches"].([]interface{}); ok {
		researches := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				researches = append(researches, str)
			}
		}
		oneCareerData.Researches = researches
	}
	if val, ok := data["research_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		oneCareerData.ResearchDescriptions = descriptions
	}
	if val, ok := data["intern_experiences"].([]interface{}); ok {
		experiences := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				experiences = append(experiences, str)
			}
		}
		oneCareerData.InternExperiences = experiences
	}
	if val, ok := data["intern_experience_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		oneCareerData.InternExperienceDescriptions = descriptions
	}
	if val, ok := data["products"].([]interface{}); ok {
		products := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				products = append(products, str)
			}
		}
		oneCareerData.Products = products
	}
	if val, ok := data["product_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		oneCareerData.ProductDescriptions = descriptions
	}

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

func (r *aiGenerationRepository) SaveLevtechRookieData(ctx context.Context, userID uuid.UUID, data map[string]interface{}) error {
	levtechData := &entity.LevtechRookie{
		ID: userID, // user_idを直接IDとして使用
	}

	// 文字列フィールドのマッピング
	if val, ok := data["portfolio"].(string); ok {
		levtechData.Portfolio = val
	}
	if val, ok := data["portfolio_description"].(string); ok {
		levtechData.PortfolioDescription = val
	}
	if val, ok := data["research"].(string); ok {
		levtechData.Research = val
	}
	if val, ok := data["organization"].(string); ok {
		levtechData.Organization = val
	}
	if val, ok := data["other"].(string); ok {
		levtechData.Other = val
	}

	// 配列フィールドのマッピング
	if val, ok := data["desired_job_type"].([]interface{}); ok {
		jobTypes := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				jobTypes = append(jobTypes, str)
			}
		}
		levtechData.DesiredJobType = jobTypes
	}
	if val, ok := data["career_aspiration"].([]interface{}); ok {
		aspirations := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				aspirations = append(aspirations, str)
			}
		}
		levtechData.CareerAspiration = aspirations
	}
	if val, ok := data["interested_tasks"].([]interface{}); ok {
		tasks := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				tasks = append(tasks, str)
			}
		}
		levtechData.InterestedTasks = tasks
	}
	if val, ok := data["job_requirements"].([]interface{}); ok {
		requirements := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				requirements = append(requirements, str)
			}
		}
		levtechData.JobRequirements = requirements
	}
	if val, ok := data["interested_industries"].([]interface{}); ok {
		industries := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				industries = append(industries, str)
			}
		}
		levtechData.InterestedIndustries = industries
	}
	if val, ok := data["preferred_company_size"].([]interface{}); ok {
		sizes := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				sizes = append(sizes, str)
			}
		}
		levtechData.PreferredCompanySize = sizes
	}
	if val, ok := data["interested_business_types"].([]interface{}); ok {
		types := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				types = append(types, str)
			}
		}
		levtechData.InterestedBusinessTypes = types
	}
	if val, ok := data["preferred_work_location"].([]interface{}); ok {
		locations := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				locations = append(locations, str)
			}
		}
		levtechData.PreferredWorkLocation = locations
	}
	if val, ok := data["skills"].([]interface{}); ok {
		skills := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				skills = append(skills, str)
			}
		}
		levtechData.Skills = skills
	}
	if val, ok := data["skill_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		levtechData.SkillDescriptions = descriptions
	}
	if val, ok := data["intern_experiences"].([]interface{}); ok {
		experiences := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				experiences = append(experiences, str)
			}
		}
		levtechData.InternExperiences = experiences
	}
	if val, ok := data["intern_experience_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		levtechData.InternExperienceDescriptions = descriptions
	}
	if val, ok := data["hackathon_experiences"].([]interface{}); ok {
		experiences := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				experiences = append(experiences, str)
			}
		}
		levtechData.HackathonExperiences = experiences
	}
	if val, ok := data["hackathon_experience_descriptions"].([]interface{}); ok {
		descriptions := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				descriptions = append(descriptions, str)
			}
		}
		levtechData.HackathonExperienceDescriptions = descriptions
	}
	if val, ok := data["certifications"].([]interface{}); ok {
		certifications := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				certifications = append(certifications, str)
			}
		}
		levtechData.Certifications = certifications
	}
	if val, ok := data["languages"].([]interface{}); ok {
		languages := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				languages = append(languages, str)
			}
		}
		levtechData.Languages = languages
	}
	if val, ok := data["language_levels"].([]interface{}); ok {
		levels := make([]string, 0, len(val))
		for _, v := range val {
			if str, ok := v.(string); ok {
				levels = append(levels, str)
			}
		}
		levtechData.LanguageLevels = levels
	}

	if err := r.db.WithContext(ctx).Save(levtechData).Error; err != nil {
		return fmt.Errorf("failed to save levtech rookie data: %w", err)
	}
	return nil
}
