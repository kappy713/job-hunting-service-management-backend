package main

import (
	"log"

	"job-hunting-service-management-backend/app/infrastructure/db"
	"job-hunting-service-management-backend/app/internal/handler"
	"job-hunting-service-management-backend/app/internal/repository"
	"job-hunting-service-management-backend/app/internal/router"
	"job-hunting-service-management-backend/app/internal/usecase"

	"github.com/joho/godotenv"
)

func main() {
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
	
	// 新しいエンティティのDI
	userRepository := repository.NewUserRepository(database)
	userUsecase := usecase.NewUserUsecase(userRepository)
	userHandler := handler.NewUserHandler(userUsecase)

	// ルーター設定
	r := router.NewRouter(sampleUserHandler, userHandler)

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}