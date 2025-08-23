package main

import (
	"log"
	"time"

	"job-hunting-service-management-backend/app/infrastructure/client"
	"job-hunting-service-management-backend/app/infrastructure/db"
	"job-hunting-service-management-backend/app/internal/handler"
	"job-hunting-service-management-backend/app/internal/repository"
	"job-hunting-service-management-backend/app/internal/router"
	"job-hunting-service-management-backend/app/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {
	// 日本時間をデフォルトに設定
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)
	time.Local = jst

	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load .env file:", err)
	}

	// DB接続
	database, err := db.NewDB()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	defer func() {
		if err := db.Close(database); err != nil {
			log.Printf("Failed to close database connection: %v", err)
		}
	}()

	log.Println("Database connection successful!")

	sampleUserRepository := repository.NewSampleUserRepository(database)
	sampleUserUsecase := usecase.NewSampleUserUsecase(sampleUserRepository)
	sampleUserHandler := handler.NewSampleUserHandler(sampleUserUsecase)

	userRepository := repository.NewUserRepository(database)

	logRepository := repository.NewLogRepository(database)
	logUsecase := usecase.NewLogUsecase(logRepository)
	logHandler := handler.NewLogHandler(logUsecase)

	// AI生成機能（userUsecaseより先に初期化）
	geminiClient := client.NewGeminiClient()
	aiGenerationRepository := repository.NewAIGenerationRepository(database)
	aiGenerationUsecase := usecase.NewAIGenerationUsecase(aiGenerationRepository, geminiClient)
	aiGenerationHandler := handler.NewAIGenerationHandler(aiGenerationUsecase)

	// UserUsecaseにAI生成機能を依存として渡す
	userUsecase := usecase.NewUserUsecase(userRepository, aiGenerationUsecase)
	userHandler := handler.NewUserHandler(userUsecase)

	supporterzRepository := repository.NewSupporterzRepository(database)
	supporterzUsecase := usecase.NewSupporterzUsecase(supporterzRepository, logUsecase)
	supporterzHandler := handler.NewSupporterzHandler(supporterzUsecase)

	careerSelectRepository := repository.NewCareerSelectRepository(database)
	careerSelectUsecase := usecase.NewCareerSelectUsecase(careerSelectRepository, logUsecase)
	careerSelectHandler := handler.NewCareerSelectHandler(careerSelectUsecase)

	levtechRookieRepository := repository.NewLevtechRookieRepository(database)
	levtechRookieUsecase := usecase.NewLevtechRookieUsecase(levtechRookieRepository, logUsecase)
	levtechRookieHandler := handler.NewLevtechRookieHandler(levtechRookieUsecase)

	mynaviRepository := repository.NewMynaviRepository(database)
	mynaviUsecase := usecase.NewMynaviUsecase(mynaviRepository, logUsecase)
	mynaviHandler := handler.NewMynaviHandler(mynaviUsecase)

	oneCareerRepository := repository.NewOneCareerRepository(database)
	oneCareerUsecase := usecase.NewOneCareerUsecase(oneCareerRepository, logUsecase)
	oneCareerHandler := handler.NewOneCareerHandler(oneCareerUsecase)

	// AI生成機能
	geminiClient := client.NewGeminiClient()
	aiGenerationRepository := repository.NewAIGenerationRepository(database)
	aiGenerationUsecase := usecase.NewAIGenerationUsecase(aiGenerationRepository, geminiClient)
	aiGenerationHandler := handler.NewAIGenerationHandler(aiGenerationUsecase)

	// ES API関連のDI ---
	profileRepository := repository.NewProfileRepository(database)
	profileUsecase := usecase.NewProfileUsecase(profileRepository, logUsecase)
	profileHandler := handler.NewProfileHandler(profileUsecase)

	// ルーター設定
	r := router.NewRouter(
		sampleUserHandler,
		userHandler,
		supporterzHandler,
		careerSelectHandler,
		levtechRookieHandler,
		mynaviHandler,
		oneCareerHandler,
		logHandler,
		aiGenerationHandler,
		profileHandler,
	)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}
