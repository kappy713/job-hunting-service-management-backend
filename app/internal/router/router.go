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
	uh handler.UserHandler,
	sh handler.SupporterzHandler,
	csh handler.CareerSelectHandler,
	lrh handler.LevtechRookieHandler,
	mh handler.MynaviHandler,
	och handler.OneCareerHandler,
	lh handler.LogHandler,
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
	r.POST("/api/user", uh.UpdateUser)
	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("/services", uh.UpdateUserServices) // 新しいエンドポイント
		userRoutes.POST("", uh.CreateUser)
	}

	// サポーターズ
	r.GET("/api/supporterz/:id", sh.GetSupporterzByID)
	r.POST("/api/supporterz", sh.CreateOrUpdateSupporterz)

	// キャリアセレクト
	r.GET("/api/career-select/:id", csh.GetCareerSelectByID)
	r.POST("/api/career-select", csh.CreateOrUpdateCareerSelect)

	// レバテックルーキー
	r.GET("/api/levtech-rookie/:id", lrh.GetLevtechRookieByID)
	r.POST("/api/levtech-rookie", lrh.CreateOrUpdateLevtechRookie)

	// マイナビ
	r.GET("/api/mynavi/:id", mh.GetMynaviByID)
	r.POST("/api/mynavi", mh.CreateOrUpdateMynavi)

	// ワンキャリア
	r.GET("/api/one-career/:id", och.GetOneCareerByID)
	r.POST("/api/one-career", och.CreateOrUpdateOneCareer)

	// ログ
	r.GET("/api/log/:id", lh.GetLogsByUserID)

	return r
}
