# 文档索引

欢迎来到 Base Gin DDD 脚手架项目文档！这里包含了项目的完整指南和参考资料。

## 📖 主要文档

### 🚀 快速开始

- **[快速开始指南](./quick-start.md)** - 5分钟快速上手项目

### 🔌 API 参考

- **[API 接口文档](./api.md)** - 完整的 REST API 文档

### 🏗️ 架构设计

- **[架构设计文档](./architecture.md)** - DDD 分层架构详解

### 🛠️ 开发环境

- **[VS Code 配置指南](./vscode.md)** - 开发环境配置和扩展推荐

### 👨‍💻 开发指南

- **[开发指南](./development.md)** - 开发规范、测试策略和部署指南

### 🔧 问题排查

- **[故障排除指南](./troubleshooting.md)** - 常见问题和解决方案

## 🎯 快速导航

### 新手入门

1. [环境要求](./quick-start.md#环境要求)
2. [安装依赖](./quick-start.md#1-安装依赖)
3. [运行应用](./quick-start.md#3-运行应用)
4. [验证服务](./quick-start.md#4-验证服务)

### API 使用

- [用户管理 API](./api.md#用户管理)
- [错误处理](./api.md#错误处理)
- [验证规则](./api.md#验证规则)

### 架构理解

- [分层架构](./architecture.md#分层详解)
- [依赖关系](./architecture.md#依赖关系)
- [数据流](./architecture.md#数据流)

### 开发实践

- [编码规范](./development.md#编码规范)
- [添加新功能](./development.md#添加新功能)
- [测试策略](./development.md#测试策略)

### 问题解决

- [环境问题](./troubleshooting.md#环境问题)
- [Wire 问题](./troubleshooting.md#wire-问题)
- [运行时问题](./troubleshooting.md#运行时问题)

## 🎨 项目特色

### DDD 架构优势

- **清晰的职责分离**：每层都有明确的职责边界
- **业务逻辑集中**：核心业务规则位于领域层
- **技术无关性**：业务逻辑不依赖技术框架
- **高度可测试**：每层都可以独立测试

### Wire 依赖注入

- **编译时生成**：零运行时开销
- **类型安全**：编译期检查依赖关系
- **易于维护**：自动管理复杂的依赖图

### Mock 数据系统

- **开箱即用**：无需配置数据库即可运行
- **真实模拟**：模拟真实的数据访问模式
- **易于替换**：一键切换到真实数据库

## 📋 常用命令

### 开发命令

```bash
# 生成 Wire 代码
cd wire && wire && cd ..

# 运行应用
go run cmd/main.go

# 运行测试
go test ./...

# 格式化代码
go fmt ./...
```

### 构建部署

```bash
# 构建应用
go build -o app cmd/main.go

# 交叉编译
GOOS=linux GOARCH=amd64 go build -o app-linux cmd/main.go
```

## 🤝 贡献指南

### 开发流程

1. Fork 项目
2. 创建功能分支
3. 提交变更
4. 发起 Pull Request

### 代码规范

- 遵循 Go 官方编码规范
- 添加适当的注释和文档
- 编写单元测试
- 确保所有测试通过

## 📚 学习资源

### DDD 相关

- [领域驱动设计](https://book.douban.com/subject/26819666/)
- [实现领域驱动设计](https://book.douban.com/subject/25844633/)

### Go 相关

- [Go 官方文档](https://golang.org/doc/)
- [Effective Go](https://golang.org/doc/effective_go)

### 架构相关

- [Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)
- [Hexagonal Architecture](https://alistair.cockburn.us/hexagonal-architecture/)

## 📞 获取帮助

- **文档问题**：查看 [故障排除指南](./troubleshooting.md)
- **功能请求**：在 GitHub 上提交 Issue
- **Bug 报告**：在 GitHub 上提交 Issue 并提供详细信息

---

💡 **提示**: 建议按照文档顺序阅读，先从快速开始指南开始，然后深入了解架构和开发指南。
