package router

import (
	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/handler"
)

func NewRouter(
	suh handler.SampleUserHandler,
	uh handler.UserHandler, // 新しいハンドラーを追加
) *gin.Engine {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World",
			"status":  "Database connected successfully",
		})
	})

	// サンプルユーザー
	r.GET("/api/sample-users", suh.GetAllSampleUsers)

	// 新しいエンドポイント
	r.GET("/api/users/:userID", uh.GetUserByID)

	return r
}
