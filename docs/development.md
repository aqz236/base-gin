# 开发指南

## 开发环境设置

### 必需工具

- Go 1.21+
- Git
- VS Code (推荐) 或其他 Go IDE

### VS Code 扩展推荐

- Go (官方扩展)
- REST Client (测试 API)
- GitLens (Git 增强)

### Wire 工具安装

```bash
go install github.com/google/wire/cmd/wire@latest
```

## VS Code 配置

项目已配置好 VS Code 开发环境，包含以下配置文件：

### 配置文件说明

- `.vscode/settings.json` - VS Code 工作区设置
- `.vscode/launch.json` - 调试配置
- `.vscode/extensions.json` - 推荐扩展列表

### 推荐扩展

项目在 `.vscode/extensions.json` 中配置了推荐扩展，主要包括：

#### 核心开发扩展

- **golang.go** - Go 官方扩展，提供语法高亮、智能感知、调试等
- **eamodio.gitlens** - Git 增强工具
- **humao.rest-client** - API 测试工具
- **ms-azuretools.vscode-docker** - Docker 支持
- **yzhang.markdown-all-in-one** - Markdown 文档编写

#### 安装推荐扩展

```bash
# VS Code 会自动提示安装推荐扩展
# 也可以通过命令面板手动安装：
# Ctrl+Shift+P -> Extensions: Show Recommended Extensions
```

#### 扩展用途说明

- **API 开发**: REST Client、Thunder Client 用于接口测试
- **数据库**: SQL Tools 系列用于数据库连接和查询
- **容器化**: Docker 扩展支持容器开发
- **文档**: Markdown 扩展用于编写项目文档
- **代码质量**: GitLens、Better Comments 提升代码可读性

### 调试配置

## 开发工作流

### 1. 代码生成

每次修改 Wire 配置后，需要重新生成依赖注入代码：

```bash
cd wire
wire
cd ..
```

### 2. 代码检查

```bash
# 格式化代码
go fmt ./...

# 检查代码
go vet ./...

# 运行测试
go test ./...

# 构建项目
go build ./...
```

### 3. 依赖管理

```bash
# 添加依赖
go get github.com/example/package

# 更新依赖
go get -u ./...

# 清理无用依赖
go mod tidy
```

## 编码规范

### 目录结构规范

```txt
internal/domain/[module]/
├── entity/          # 实体：业务对象
├── repository/      # 仓储接口：数据访问抽象
├── service/         # 领域服务：跨实体业务逻辑
└── vo/             # 值对象：数据传输对象
```

### 命名规范

#### 文件命名

- 使用小写字母和下划线：`user_service.go`
- 接口文件以 `_interface` 结尾：`user_repository.go`
- 实现文件以 `_impl` 结尾：`user_repository.go`

#### 变量命名

- 使用驼峰命名：`userService`
- 常量使用大写：`const DefaultPageSize = 10`
- 私有变量小写开头：`userRepo`
- 公有变量大写开头：`UserService`

#### 函数命名

- 构造函数以 `New` 开头：`NewUserService()`
- 获取器以 `Get` 开头：`GetUser()`
- 设置器以 `Set` 开头：`SetEmail()`
- 布尔函数以 `Is` 或 `Has` 开头：`IsValid()`, `HasPermission()`

### 错误处理

#### 错误定义

在 `internal/pkg/errors/errors.go` 中定义业务错误：

```go
var (
    ErrUserNotFound = NewAppError("USER_NOT_FOUND", "用户不存在", "")
    ErrEmailExists  = NewAppError("EMAIL_EXISTS", "邮箱已存在", "")
)
```

#### 错误返回

```go
// 好的做法
func (s *UserService) GetUser(id int) (*vo.UserResponse, error) {
    user, err := s.userRepo.FindByID(id)
    if err != nil {
        return nil, errors.ErrUserNotFound
    }
    // ...
}

// 避免的做法
func (s *UserService) GetUser(id int) (*vo.UserResponse, error) {
    user, err := s.userRepo.FindByID(id)
    if err != nil {
        return nil, fmt.Errorf("用户不存在") // 避免硬编码错误信息
    }
    // ...
}
```

### 日志规范

```go
// 使用结构化日志
logger.Info("用户创建成功", 
    "user_id", user.ID,
    "email", user.Email,
)

// 错误日志包含上下文
logger.Error("用户创建失败",
    "error", err.Error(),
    "email", req.Email,
)
```

## 添加新功能

### 1. 创建新模块

以添加 `product` 模块为例：

#### Step 1: 创建领域模型

```bash
mkdir -p internal/domain/product/{entity,repository,service,vo}
```

创建产品实体：

```go
// internal/domain/product/entity/product.go
package entity

type Product struct {
    ID          int
    Name        string
    Description string
    Price       float64
    CreatedAt   time.Time
    UpdatedAt   time.Time
}

func (p *Product) Validate() error {
    // 验证逻辑
}
```

创建仓储接口：

```go
// internal/domain/product/repository/product_repository.go
package repository

import "base-gin/internal/domain/product/entity"

type ProductRepository interface {
    FindByID(id int) (*entity.Product, error)
    Save(product *entity.Product) error
    // 其他方法
}
```

#### Step 2: 实现基础设施

```go
// internal/infrastructure/repository/product_impl/product_repository.go
package product_impl

type ProductRepo struct {
    // 实现细节
}

func NewProductRepository() *ProductRepo {
    return &ProductRepo{}
}

func (r *ProductRepo) FindByID(id int) (*entity.Product, error) {
    // 实现逻辑
}
```

#### Step 3: 创建应用服务

```go
// internal/app/product/service/product_service.go
package service

type ProductService struct {
    productRepo repository.ProductRepository
}

func NewProductService(repo repository.ProductRepository) *ProductService {
    return &ProductService{productRepo: repo}
}
```

#### Step 4: 添加控制器

```go
// internal/interfaces/handler/product/product_handler.go
package product

type ProductHandler struct {
    productService *service.ProductService
}

func NewProductHandler(service *service.ProductService) *ProductHandler {
    return &ProductHandler{productService: service}
}
```

#### Step 5: 配置依赖注入

在 `wire/provider.go` 中添加：

```go
// 产品模块依赖
var ProductSet = wire.NewSet(
    product_impl.NewProductRepository,
    wire.Bind(new(product_repository.ProductRepository), new(*product_impl.ProductRepo)),
    product_service.NewProductService,
    product_handler.NewProductHandler,
)
```

在 `wire/wire.go` 中添加到构建集合：

```go
func InitializeApp() (*App, func(), error) {
    panic(wire.Build(
        InfraSet,
        RepositorySet,
        ProductSet,        // 新增
        // ...
    ))
}
```

#### Step 6: 注册路由

在 `internal/interfaces/router/router.go` 中添加：

```go
func NewRouter(
    userHandler *user.UserHandler,
    productHandler *product.ProductHandler, // 新增
) *gin.Engine {
    // ...
    
    api := r.Group("/api/v1")
    {
        // 产品路由
        productGroup := api.Group("/products")
        {
            productGroup.GET("", productHandler.GetAllProducts)
            productGroup.GET("/:id", productHandler.GetProduct)
            productGroup.POST("", productHandler.CreateProduct)
            // ...
        }
    }
}
```

#### Step 7: 重新生成 Wire 代码

```bash
cd wire
wire
cd ..
```

### 2. 修改现有功能

#### 修改实体

1. 在 `internal/domain/[module]/entity/` 中修改实体
2. 确保通过验证方法保持业务规则
3. 运行测试确保修改不破坏现有功能

#### 修改仓储

1. 在仓储接口中添加新方法
2. 在仓储实现中实现新方法
3. 更新相关测试

#### 修改 API

1. 在控制器中添加新的处理方法
2. 在路由中注册新的端点
3. 更新 API 文档

## 测试策略

### 单元测试

每个模块都应该有对应的单元测试：

```txt
test/unit/
├── user_test.go         # 用户模块测试
├── product_test.go      # 产品模块测试
└── ...
```

#### 测试领域实体

```go
func TestUser_Validate(t *testing.T) {
    tests := []struct {
        name    string
        user    *entity.User
        wantErr bool
    }{
        {
            name: "valid user",
            user: &entity.User{
                Name:     "张三",
                Email:    "zhangsan@example.com",
                Password: "password123",
            },
            wantErr: false,
        },
        // 更多测试案例
    }
    
    for _, tt := range tests {
        t.Run(tt.name, func(t *testing.T) {
            err := tt.user.Validate()
            if (err != nil) != tt.wantErr {
                t.Errorf("Validate() error = %v, wantErr %v", err, tt.wantErr)
            }
        })
    }
}
```

#### 测试应用服务

```go
func TestUserService_CreateUser(t *testing.T) {
    // 创建 mock 仓储
    mockRepo := &MockUserRepository{}
    mockDomainService := &MockUserDomainService{}
    
    service := NewUserService(mockRepo, mockDomainService)
    
    // 测试逻辑
}
```

### 集成测试

测试多层协作：

```go
func TestAPI_CreateUser(t *testing.T) {
    // 初始化完整应用
    app, cleanup, err := wire.InitializeApp()
    if err != nil {
        t.Fatal(err)
    }
    defer cleanup()
    
    // 创建测试服务器
    ts := httptest.NewServer(app.Router)
    defer ts.Close()
    
    // 发送请求并验证响应
}
```

### 运行测试

```bash
# 运行所有测试
go test ./...

# 运行特定包的测试
go test ./test/unit/

# 运行测试并显示覆盖率
go test -cover ./...

# 生成覆盖率报告
go test -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## 调试技巧

### 使用 Delve 调试器

```bash
# 安装 Delve
go install github.com/go-delve/delve/cmd/dlv@latest

# 启动调试
dlv debug cmd/main.go

# 在 VS Code 中调试
# 使用 F5 或点击调试按钮，使用预配置的 launch.json
```

### 日志调试

在代码中添加调试日志：

```go
logger.Debug("处理用户创建请求", 
    "request", fmt.Sprintf("%+v", req))

logger.Debug("调用仓储保存用户", 
    "user_id", user.ID)
```

### 性能分析

```bash
# 启用性能分析
go run -pprof=cpu cmd/main.go

# 查看性能报告
go tool pprof http://localhost:8080/debug/pprof/profile
```

## 常见问题

### Wire 相关问题

**问题**: Wire 生成失败

**解决方案**:

1. 检查 provider 定义是否正确
2. 确保所有依赖都有对应的 provider
3. 检查循环依赖问题

**问题**: 找不到 Wire 命令

**解决方案**:

```bash
# 确保 GOPATH/bin 在 PATH 中
export PATH=$PATH:$(go env GOPATH)/bin

# 重新安装 Wire
go install github.com/google/wire/cmd/wire@latest
```

### 依赖注入问题

**问题**: 接口绑定失败

**解决方案**:
确保在 provider 中正确绑定接口：

```go
wire.Bind(new(repository.UserRepository), new(*user_impl.UserRepo))
```

### 性能问题

**问题**: 响应慢

**解决方案**:

1. 检查数据库查询效率
2. 添加必要的缓存
3. 使用性能分析工具定位瓶颈

## 部署指南

### 构建

```bash
# 构建应用
go build -o app cmd/main.go

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o app-linux cmd/main.go
```

### Docker 部署

创建 Dockerfile：

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o main cmd/main.go

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/main .

CMD ["./main"]
```

构建和运行：

```bash
docker build -t base-gin .
docker run -p 8080:8080 base-gin
```
