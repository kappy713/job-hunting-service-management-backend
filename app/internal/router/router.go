package router

import (
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/handler"
)

func NewRouter(
	suh handler.SampleUserHandler,
	uh handler.UserHandler, // 新しいハンドラーを追加
) *gin.Engine {
	r := gin.Default()

	frontendURL := os.Getenv("FRONTEND_URL")

	// CORS設定
	config := cors.Config{
		AllowOrigins:     []string{frontendURL},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(config))

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"status":  "Database connected successfully",
		})
	})

	// サンプルユーザー
	r.GET("/api/sample-users", suh.GetAllSampleUsers)

	// ユーザー
	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("/services", uh.UpdateUserServices) // 新しいエンドポイント
	}

	return r
}
