//go:build wireinject
// +build wireinject

package wire

import (
	"base-gin/internal/infrastructure/cache"
	"base-gin/internal/infrastructure/database"
	"base-gin/internal/infrastructure/logging"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

// App 应用结构
type App struct {
	Router *gin.Engine
	DB     *database.DB
	Cache  *cache.RedisClient
	Logger *logging.Logger
}

// NewApp 创建应用实例
func NewApp(
	router *gin.Engine,
	db *database.DB,
	cache *cache.RedisClient,
	logger *logging.Logger,
) *App {
	return &App{
		Router: router,
		DB:     db,
		Cache:  cache,
		Logger: logger,
	}
}

// InitializeApp 初始化应用
func InitializeApp() (*App, func(), error) {
	panic(wire.Build(
		InfraSet,         // 基础设施层
		RepositorySet,    // 仓储层
		DomainServiceSet, // 领域服务层
		ServiceSet,       // 应用服务层
		ValidationSet,    // 验证层
		HandlerSet,       // 控制器层
		RouterSet,        // 路由层
		NewApp,           // 应用构造函数
	))
}
