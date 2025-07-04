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
	configs.LoadConfig,
	database.NewDB,
	cache.NewRedisClient,
	logging.NewLogger,
)

// 仓储层依赖
var RepositorySet = wire.NewSet(
	user_impl.NewUserRepository,
	wire.Bind(new(repository.UserRepository), new(*user_impl.UserRepo)),
)

// 领域服务依赖
var DomainServiceSet = wire.NewSet(
	domainService.NewUserDomainService,
)

// 应用服务依赖
var ServiceSet = wire.NewSet(
	service.NewUserService,
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
