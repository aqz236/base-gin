package router

import (
	"base-gin/internal/interfaces/handler/user"
	"base-gin/internal/interfaces/middleware"

	"github.com/gin-gonic/gin"
)

func NewRouter(userHandler *user.UserHandler) *gin.Engine {
	r := gin.New()

	// 注册中间件
	r.Use(middleware.Logger())
	r.Use(middleware.Recovery())
	r.Use(middleware.CORS())

	// 健康检查
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	// API 路由组
	api := r.Group("/api/v1")
	{
		// 用户路由
		userGroup := api.Group("/users")
		{
			userGroup.GET("", userHandler.GetAllUsers)
			userGroup.GET("/:id", userHandler.GetUser)
			userGroup.POST("", userHandler.CreateUser)
			userGroup.PUT("/:id", userHandler.UpdateUser)
			userGroup.DELETE("/:id", userHandler.DeleteUser)
		}
	}

	return r
}
