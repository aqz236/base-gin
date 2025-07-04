package wire

import (
	"base-gin/configs"
	"base-gin/internal/app/user/service"
	"base-gin/internal/domain/user/repository"
	domainService "base-gin/internal/domain/user/service"
	"base-gin/internal/infrastructure/cache"
	"base-gin/internal/infrastructure/database"
	"base-gin/internal/infrastructure/logging"
	"base-gin/internal/infrastructure/repository/user_impl"
	"base-gin/internal/interfaces/handler/user"
	"base-gin/internal/interfaces/router"
	"base-gin/internal/interfaces/validation"

	"github.com/google/wire"
)

// 基础设施层依赖
var InfraSet = wire.NewSet(
	configs.LoadConfig,   // 提供 *configs.Config
	database.NewDB,       // 需要 *configs.Config，提供 *database.DB
	cache.NewRedisClient, // 需要 *configs.Config，提供 *cache.RedisClient
	logging.NewLogger,    // 需要 *configs.Config，提供 *logging.Logger
)

// 仓储层依赖
var RepositorySet = wire.NewSet(
	user_impl.NewGormUserRepository, // 需要 *database.DB，提供 *user_impl.GormUserRepository
	wire.Bind(
		new(repository.UserRepository),
		new(*user_impl.GormUserRepository)),
) // 接口绑定

// 领域服务依赖
var DomainServiceSet = wire.NewSet(
	domainService.NewUserDomainService, // 需要 repository.UserRepository，提供 *UserDomainService
)

// 应用服务依赖
var ServiceSet = wire.NewSet(
	service.NewUserService, // 需要 repository.UserRepository 和 *UserDomainService
)

// 验证器依赖
var ValidationSet = wire.NewSet(
	validation.NewValidator,
)

// 控制器依赖
var HandlerSet = wire.NewSet(
	user.NewUserHandler,
)

// 路由依赖
var RouterSet = wire.NewSet(
	router.NewRouter,
)
