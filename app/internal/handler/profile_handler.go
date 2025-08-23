package handler

import (
	"fmt"
	"net/http"
	"strings"
	"unicode/utf8"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	"job-hunting-service-management-backend/app/internal/entity"
	"job-hunting-service-management-backend/app/internal/usecase"
)

// ProfileHandler プロフィール関連のHTTPリクエストを処理するインターフェース
type ProfileHandler interface {
	GetProfileByUserID(c *gin.Context)    // ユーザーIDでプロフィール情報を取得
	CreateOrUpdateProfile(c *gin.Context) // プロフィール情報を作成または更新
}

type profileHandler struct {
	pu usecase.ProfileUsecase
}

// NewProfileHandler 新しいProfileHandlerのインスタンスを作成します
func NewProfileHandler(u usecase.ProfileUsecase) ProfileHandler {
	return &profileHandler{pu: u}
}

// GetProfileByUserID はユーザーIDに基づいてプロフィール情報を取得します
func (h *profileHandler) GetProfileByUserID(c *gin.Context) {
	// URLパラメータからIDを取得
	idParam := c.Param("id")
	if idParam == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID parameter is required",
		})
		return
	}

	// UUIDのパース
	userID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"details": "User ID must be a valid UUID",
		})
		return
	}

	// 空のユーザーIDチェック
	if userID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID cannot be empty",
		})
		return
	}

	// プロフィール取得
	profile, err := h.pu.GetProfileByUserID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error":   "Failed to retrieve profile",
			"details": err.Error(),
		})
		return
	}

	// プロフィールが見つからない場合
	if profile == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "Profile not found for the specified user",
		})
		return
	}

	// 成功レスポンス（UTF-8で明示的に設定）
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile retrieved successfully",
		"profile": profile,
	})
}

// CreateOrUpdateProfile はプロフィール情報を作成または更新します
func (h *profileHandler) CreateOrUpdateProfile(c *gin.Context) {
	// Content-Typeの確認（文字化け対策）
	contentType := c.GetHeader("Content-Type")
	if !strings.Contains(contentType, "application/json") {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Content-Type must be application/json",
		})
		return
	}

	var req entity.CreateProfileRequest

	// JSONリクエストボディのバインド
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid request format",
			"details": err.Error(),
		})
		return
	}

	// UTF-8エンコーディングの検証（文字化け対策）
	if err := h.validateUTF8Encoding(req.Data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid character encoding",
			"details": err.Error(),
		})
		return
	}

	// ユーザーIDの検証
	userID, err := uuid.Parse(req.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error":   "Invalid user ID format",
			"details": "User ID must be a valid UUID",
		})
		return
	}

	// 空のユーザーIDチェック
	if userID == uuid.Nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "User ID cannot be empty",
		})
		return
	}

	// プロフィールの作成または更新
	profile, err := h.pu.CreateOrUpdateProfile(c, userID, req.Data)
	if err != nil {
		// エラーの種類に応じてステータスコードを決定
		statusCode := http.StatusInternalServerError
		if strings.Contains(err.Error(), "validation failed") {
			statusCode = http.StatusBadRequest
		}

		c.JSON(statusCode, gin.H{
			"error":   "Failed to create or update profile",
			"details": err.Error(),
		})
		return
	}

	// 成功レスポンス（UTF-8で明示的に設定）
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.JSON(http.StatusOK, gin.H{
		"message": "Profile created/updated successfully",
		"profile": profile,
	})
}

// validateUTF8Encoding はプロフィールデータのUTF-8エンコーディングを検証します（文字化け対策）
func (h *profileHandler) validateUTF8Encoding(data entity.ProfileData) error {
	// 文字列フィールドのUTF-8検証
	fields := []string{
		data.CareerVision,
		data.SelfPromotion,
		data.StudentExperience,
		data.Research,
		data.Organization,
		data.DesiredJobType,
		data.CompanySelectionCriteria,
		data.EngineerAspiration,
	}

	for _, field := range fields {
		if !utf8.ValidString(field) {
			return fmt.Errorf("invalid UTF-8 encoding detected in text field")
		}
	}

	// スライスフィールドのUTF-8検証
	sliceFields := [][]string{
		data.Products,
		data.ProductDescriptions,
		data.Skills,
		data.SkillDescriptions,
		data.Interns,
		data.InternDescriptions,
		data.Certifications,
		data.CertificationDescriptions,
	}

	for _, slice := range sliceFields {
		for _, item := range slice {
			if !utf8.ValidString(item) {
				return fmt.Errorf("invalid UTF-8 encoding detected in array field")
			}
		}
	}

	return nil
}
