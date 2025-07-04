# 架构设计文档

## 概述

Base Gin 采用领域驱动设计（DDD）的分层架构，结合依赖注入模式，实现了高内聚、低耦合的代码结构。

## 整体架构

```txt
┌─────────────────────────────────────────────────────────────┐
│                    接口层 (Interfaces)                      │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌──────────┐ │
│  │   Handler   │ │   Router    │ │ Middleware  │ │Validation│ │
│  │   (控制器)   │ │   (路由)    │ │  (中间件)   │ │  (验证)   │ │
│  └─────────────┘ └─────────────┘ └─────────────┘ └──────────┘ │
└─────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────┐
│                    应用层 (Application)                     │
│  ┌─────────────────────────────────────────────────────────┐ │
│  │              Application Service                        │ │
│  │              (应用服务)                                  │ │
│  └─────────────────────────────────────────────────────────┘ │
└─────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────┐
│                     领域层 (Domain)                         │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌──────────┐ │
│  │   Entity    │ │ Repository  │ │   Service   │ │    VO    │ │
│  │   (实体)    │ │ (仓储接口)  │ │ (领域服务)  │ │ (值对象) │ │
│  └─────────────┘ └─────────────┘ └─────────────┘ └──────────┘ │
└─────────────────────────────────────────────────────────────┘
                               │
                               ▼
┌─────────────────────────────────────────────────────────────┐
│                  基础设施层 (Infrastructure)                │
│  ┌─────────────┐ ┌─────────────┐ ┌─────────────┐ ┌──────────┐ │
│  │  Database   │ │    Cache    │ │   Logging   │ │Repository│ │
│  │  (数据库)   │ │   (缓存)    │ │   (日志)    │ │ (仓储实现)│ │
│  └─────────────┘ └─────────────┘ └─────────────┘ └──────────┘ │
└─────────────────────────────────────────────────────────────┘
```

## 分层详解

### 1. 接口层 (Interfaces)

**职责**: 处理外部请求，提供 HTTP API 接口

**组件**:

- **Handler**: HTTP 控制器，处理 HTTP 请求和响应
- **Router**: 路由配置，定义 API 端点
- **Middleware**: 中间件，处理横切关注点（日志、认证、CORS 等）
- **Validation**: 请求验证，验证输入参数

**关键文件**:

```txt
internal/interfaces/
├── handler/user/user_handler.go    # 用户控制器
├── router/router.go                # 路由配置
├── middleware/middleware.go        # 中间件
└── validation/validator.go         # 验证器
```

### 2. 应用层 (Application)

**职责**: 协调领域对象完成业务用例，处理事务边界

**特点**:

- 薄薄的一层，不包含业务逻辑
- 协调多个领域对象完成复杂用例
- 管理事务边界
- 转换数据格式

**关键文件**:

```txt
internal/app/user/service/user_service.go    # 用户应用服务
```

### 3. 领域层 (Domain)

**职责**: 核心业务逻辑，系统的心脏

**组件**:

- **Entity**: 领域实体，包含标识和业务逻辑
- **Repository**: 仓储接口，定义数据访问抽象
- **Service**: 领域服务，处理跨实体的业务逻辑
- **VO**: 值对象，不可变的业务概念

**关键文件**:

```txt
internal/domain/user/
├── entity/user.go                   # 用户实体
├── repository/user_repository.go    # 用户仓储接口
├── service/user_domain_service.go   # 用户领域服务
└── vo/user_vo.go                   # 用户值对象
```

### 4. 基础设施层 (Infrastructure)

**职责**: 提供技术实现，支撑上层业务逻辑

**组件**:

- **Database**: 数据库连接和配置
- **Cache**: 缓存实现
- **Logging**: 日志实现
- **Repository**: 仓储接口的具体实现

**关键文件**:

```txt
internal/infrastructure/
├── database/database.go                        # 数据库
├── cache/redis.go                             # 缓存
├── logging/logger.go                          # 日志
└── repository/user_impl/user_repository.go    # 用户仓储实现
```

## 依赖关系

### 依赖方向

```txt
接口层 ──┐
        ├──► 应用层 ──► 领域层 ◄── 基础设施层
        └──────────────► 领域层
```

**核心原则**:

1. **高层模块不依赖低层模块**，都依赖于抽象
2. **抽象不依赖细节**，细节依赖抽象
3. **内层不知道外层的存在**

### 依赖注入

使用 Google Wire 进行编译时依赖注入：

```txt
wire/
├── provider.go    # 依赖提供者定义
├── wire.go        # Wire 配置文件
└── wire_gen.go    # 自动生成的依赖注入代码
```

**依赖注入流程**:

1. 在 `provider.go` 中定义各层的依赖提供者
2. 在 `wire.go` 中配置依赖关系
3. 运行 `wire` 命令生成注入代码
4. 在 `main.go` 中调用生成的初始化函数

## 模块结构

### 用户模块示例

```txt
User Module
├── Domain Layer
│   ├── Entity: User (用户实体)
│   ├── Repository: UserRepository (仓储接口)
│   ├── Service: UserDomainService (领域服务)
│   └── VO: UserCreateRequest, UserResponse (值对象)
├── Application Layer
│   └── Service: UserService (应用服务)
├── Infrastructure Layer
│   └── Repository: UserRepo (仓储实现)
└── Interface Layer
    └── Handler: UserHandler (HTTP 控制器)
```

## 数据流

### 请求处理流程

```txt
HTTP Request
    │
    ▼
[Router] ──► [Middleware] ──► [Handler]
                                │
                                ▼
                         [Application Service]
                                │
                                ▼
                         [Domain Service] ◄──► [Domain Entity]
                                │
                                ▼
                         [Repository Interface]
                                │
                                ▼
                         [Repository Implementation]
                                │
                                ▼
                            [Database]
```

### 响应处理流程

```txt
[Database]
    │
    ▼
[Repository Implementation]
    │
    ▼
[Repository Interface]
    │
    ▼
[Domain Entity] ──► [Domain Service]
    │
    ▼
[Application Service]
    │
    ▼
[Handler] ──► [HTTP Response]
```

## 设计优势

### 1. 可测试性

- **单元测试**: 每层都可以独立测试
- **Mock 测试**: 通过接口可以轻松创建 Mock 对象
- **集成测试**: 可以测试多层协作

### 2. 可维护性

- **关注点分离**: 每层职责明确
- **低耦合**: 层间通过接口通信
- **高内聚**: 相关逻辑聚合在一起

### 3. 可扩展性

- **新增功能**: 只需在相应层添加代码
- **替换实现**: 通过依赖注入轻松替换组件
- **技术栈迁移**: 核心业务逻辑不受技术变更影响

### 4. 代码质量

- **业务逻辑集中**: 所有业务规则在领域层
- **框架无关**: 领域层不依赖任何框架
- **SOLID 原则**: 遵循面向对象设计原则

## 扩展指南

### 添加新模块

1. **创建领域模型**

   ```txt
   internal/domain/[module]/
   ├── entity/
   ├── repository/
   ├── service/
   └── vo/
   ```

2. **实现基础设施**

   ```txt
   internal/infrastructure/repository/[module]_impl/
   ```

3. **创建应用服务**

   ```txt
   internal/app/[module]/service/
   ```

4. **添加接口层**

   ```txt
   internal/interfaces/handler/[module]/
   ```

5. **配置依赖注入**
   - 在 `wire/provider.go` 中添加相关提供者
   - 重新生成 Wire 代码

6. **注册路由**
   - 在 `internal/interfaces/router/router.go` 中添加路由
