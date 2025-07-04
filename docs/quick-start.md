# 快速开始指南

## 环境要求

- Go 1.21+
- Git

## 0. VS Code 环境配置

项目已配置 VS Code 开发环境，建议安装推荐扩展：

```bash
# 1. 用 VS Code 打开项目
code .

# 2. VS Code 会自动提示安装推荐扩展
# 点击右下角的"安装"按钮安装所有推荐扩展

# 3. 或者手动安装：
# Ctrl+Shift+P -> Extensions: Show Recommended Extensions
```

**核心扩展说明：**

- `golang.go` - Go 官方扩展，必装
- `eamodio.gitlens` - Git 增强工具
- `humao.rest-client` - API 测试工具
- `ms-azuretools.vscode-docker` - Docker 支持

## 1. 安装依赖

```bash
# 下载项目依赖
go mod download

# 安装 Wire 工具（用于依赖注入代码生成）
go install github.com/google/wire/cmd/wire@latest
```

## 2. 生成 Wire 代码

```bash
# 进入 wire 目录
cd wire

# 生成依赖注入代码
wire

# 返回项目根目录
cd ..
```

## 3. 运行应用

```bash
# 直接运行
go run cmd/main.go

# 或者先编译再运行
go build -o app cmd/main.go
./app
```

服务器将在 `:8080` 端口启动。

## 4. 验证服务

### 健康检查

```bash
curl http://localhost:8080/health
```

预期响应：

```json
{
  "status": "ok"
}
```

### 获取用户列表

```bash
curl http://localhost:8080/api/v1/users
```

预期响应：

```json
{
  "data": [
    {
      "id": 1,
      "name": "张三",
      "email": "zhangsan@example.com"
    },
    {
      "id": 2,
      "name": "李四",
      "email": "lisi@example.com"
    },
    {
      "id": 3,
      "name": "王五",
      "email": "wangwu@example.com"
    }
  ]
}
```

## 5. 开发环境配置

### 环境变量

复制环境变量示例文件：

```bash
cp .env.example .env
```

根据需要修改 `.env` 文件中的配置。

### VS Code 配置

项目已包含 VS Code 配置文件：

- `.vscode/settings.json` - Go 语言服务配置
- `.vscode/launch.json` - 调试配置

如果遇到 Go 扩展问题，请参考 [故障排除文档](./troubleshooting.md)。

## 下一步

- 查看 [API 文档](./api.md) 了解所有可用的 API 端点
- 阅读 [架构文档](./architecture.md) 了解项目结构
- 参考 [开发指南](./development.md) 开始开发新功能
