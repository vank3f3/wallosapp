package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/vank3f3/wallosapp/internal/config"
	"github.com/vank3f3/wallosapp/internal/database"
	"github.com/vank3f3/wallosapp/internal/handlers"
	"github.com/vank3f3/wallosapp/internal/middleware"
)

func main() {
	// 加载配置
	cfg := config.LoadConfig()

	// 初始化数据库
	database.InitDB(cfg)

	// 初始化Gin引擎
	r := gin.Default()

	// 将配置存储到上下文中
	r.Use(func(c *gin.Context) {
		c.Set("config", cfg)
		c.Next()
	})

	// 设置路由
	setupRoutes(r, cfg)

	// 启动服务器
	if err := r.Run(":" + cfg.ServerPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}

func setupRoutes(r *gin.Engine, cfg *config.Config) {
	// 公共路由
	public := r.Group("/api")
	{
		public.POST("/register", handlers.Register)
		public.POST("/login", handlers.Login)
	}

	// 需要认证的路由
	auth := r.Group("/api")
	auth.Use(middleware.AuthMiddleware(cfg))
	{
		auth.GET("/profile", handlers.GetUserProfile)
	}

	// 健康检查路由
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})
} 