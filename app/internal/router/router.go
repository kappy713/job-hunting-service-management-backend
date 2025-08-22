package router

import (
	"github.com/gin-gonic/gin"

	"job-hunting-service-management-backend/app/internal/handler"
)

func NewRouter(
	suh handler.SampleUserHandler,
	uh handler.UserHandler,
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

	// ユーザー
	userRoutes := r.Group("/api/users")
	{
		userRoutes.POST("/services", uh.UpdateUserServices)
		userRoutes.POST("", uh.CreateUser)
	}

	return r
}
