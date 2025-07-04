# Base Gin - DDD 架构脚手架

这是一个基于 Gin 框架的 DDD（领域驱动设计）架构脚手架，使用 Wire 进行依赖注入。

## 技术特性

- 🏗️ **DDD 分层架构**：清晰的领域驱动设计分层
- 💉 **依赖注入**：使用 Google Wire 自动管理依赖
- 🔄 **依赖倒置**：高层模块不依赖低层模块
- 🧪 **高可测试性**：各层解耦，便于单元测试
- 📦 **模块化设计**：独立的领域模块
- 🛡️ **框架独立**：核心业务逻辑不依赖框架
- 🚀 **Mock 数据**：开箱即用的示例数据

## 架构概览

项目采用严格的 DDD 四层架构：

```txt
接口层 (HTTP API) → 应用层 (用例协调) → 领域层 (业务逻辑) → 基础设施层 (技术实现)
```

详细的架构说明请参考 [架构设计文档](./docs/architecture.md)。

## 项目结构

```txt
.
├── cmd/                       # 应用入口
├── configs/                   # 配置文件
├── internal/                  # 内部包（不对外暴露）
│   ├── app/                   # 应用层
│   │   └── user/service/      # 用户应用服务
│   ├── domain/                # 领域层
│   │   └── user/              # 用户领域
│   │       ├── entity/        # 实体
│   │       ├── repository/    # 仓储接口
│   │       ├── service/       # 领域服务
│   │       └── vo/            # 值对象
│   ├── infrastructure/        # 基础设施层
│   │   ├── database/          # 数据库
│   │   ├── cache/             # 缓存
│   │   ├── logging/           # 日志
│   │   └── repository/        # 仓储实现
│   ├── interfaces/            # 接口层
│   │   ├── handler/           # HTTP 控制器
│   │   ├── router/            # 路由
│   │   ├── middleware/        # 中间件
│   │   └── validation/        # 验证
│   └── pkg/                   # 内部共享包
├── pkg/                       # 公共包（可对外暴露）
├── test/                      # 测试
├── wire/                      # Wire 依赖注入配置
└── go.mod
```

## 快速开始

```bash
# 1. 安装依赖
go mod download

# 2. 生成 Wire 依赖注入代码
cd wire && wire && cd ..

# 3. 运行应用
go run cmd/main.go
```

服务器将在 `:8080` 端口启动。

### 验证安装

```bash
# 健康检查
curl http://localhost:8080/health

# 获取用户列表
curl http://localhost:8080/api/v1/users
```

## 文档

- 📚 [快速开始指南](./docs/quick-start.md) - 详细的安装和配置步骤
- 🔌 [API 接口文档](./docs/api.md) - 完整的 API 使用说明
- 🏗️ [架构设计文档](./docs/architecture.md) - DDD 分层架构详解
- 👨‍💻 [开发指南](./docs/development.md) - 开发规范和最佳实践
- 🔧 [故障排除指南](./docs/troubleshooting.md) - 常见问题解决方案

## 运行测试

```bash
# 运行所有测试
go test ./...

# 查看覆盖率
go test -cover ./...
```

详细的测试指南请参考 [开发指南](./docs/development.md#测试策略)。

## 许可证

MIT License
