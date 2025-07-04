# 故障排除指南

## 环境问题

### Go 环境问题

#### 问题：VS Code 无法找到 Go 二进制文件

**错误信息**:

```txt
Failed to find the "go" binary in either GOROOT() or PATH
```

**解决方案**:

1. **检查 Go 安装**:

   ```bash
   which go
   go version
   ```

2. **检查环境变量**:

   ```bash
   echo $GOROOT
   echo $GOPATH
   echo $PATH
   ```

3. **配置 VS Code 设置**:

   创建或更新 `.vscode/settings.json`:

   ```json
   {
       "go.goroot": "/path/to/your/go",
       "go.gopath": "/path/to/your/gopath",
       "go.toolsGopath": "/path/to/your/gopath"
   }
   ```

4. **使用 goenv 的用户**:

   在 `~/.zshrc` 或 `~/.bashrc` 中添加：

   ```bash
   export GOENV_ROOT="$HOME/.goenv"
   export PATH="$GOENV_ROOT/bin:$PATH"
   eval "$(goenv init -)"
   export GOPATH="$HOME/go/$(goenv version-name)"
   export GOROOT="$GOENV_ROOT/versions/$(goenv version-name)"
   export PATH="$GOROOT/bin:$GOPATH/bin:$PATH"
   ```

5. **重启 VS Code**:

   ```bash
   # 重新加载配置
   source ~/.zshrc  # 或 ~/.bashrc
   
   # 重启 VS Code
   code --reload-window
   ```

#### 问题：Go 工具安装失败

**解决方案**:

1. **确保网络正常**:

   ```bash
   go env GOPROXY
   ```

2. **设置代理（中国用户）**:

   ```bash
   go env -w GOPROXY=https://goproxy.cn,direct
   ```

3. **手动安装工具**:

   ```bash
   go install golang.org/x/tools/gopls@latest
   go install github.com/go-delve/delve/cmd/dlv@latest
   go install github.com/google/wire/cmd/wire@latest
   ```

### Wire 问题

#### 问题：Wire 构建标签警告

**错误信息**:

```txt
No packages found for open file wire.go.
This file may be excluded due to its build tags
```

**解决方案**:

1. **更新 VS Code 设置**:

   在 `.vscode/settings.json` 中添加：

   ```json
   {
       "go.buildFlags": ["-tags=wireinject"],
       "go.testFlags": ["-tags=wireinject"]
   }
   ```

2. **或者在工作区设置中配置**:

   ```json
   {
       "gopls": {
           "build.buildFlags": ["-tags=wireinject"]
       }
   }
   ```

#### 问题：Wire 生成失败

**常见错误**:

```txt
wire: generate failed: could not load packages
```

**解决方案**:

1. **检查导入路径**:
   确保所有导入路径正确，特别是相对于 `go.mod` 的路径

2. **检查依赖循环**:

   ```bash
   # 在项目根目录运行
   go list -deps ./wire
   ```

3. **清理并重新生成**:

   ```bash
   cd wire
   rm -f wire_gen.go
   wire
   ```

4. **检查 provider 定义**:
   确保所有 provider 都正确定义：

   ```go
   var ServiceSet = wire.NewSet(
       service.NewUserService,
       // 确保这里没有缺失的依赖
   )
   ```

## 依赖问题

### 模块下载问题

#### 问题：依赖下载慢或失败

**解决方案**:

1. **设置代理**:

   ```bash
   go env -w GOPROXY=https://goproxy.cn,direct
   go env -w GOSUMDB=sum.golang.google.cn
   ```

2. **清理模块缓存**:

   ```bash
   go clean -modcache
   go mod download
   ```

3. **检查网络和防火墙**:

   ```bash
   curl -I https://proxy.golang.org
   ```

#### 问题：依赖版本冲突

**解决方案**:

1. **查看依赖图**:

   ```bash
   go mod graph
   ```

2. **更新依赖**:

   ```bash
   go get -u ./...
   go mod tidy
   ```

3. **指定特定版本**:

   ```bash
   go get github.com/gin-gonic/gin@v1.9.1
   ```

## 编译问题

### 构建失败

#### 问题：找不到包

**错误示例**:

```txt
cannot find package "base-gin/internal/domain/user/entity"
```

**解决方案**:

1. **检查模块名称**:
   确保 `go.mod` 中的模块名称与导入路径一致

2. **检查文件路径**:
   确保所有文件都在正确的目录中

3. **重新初始化模块**:

   ```bash
   go mod init base-gin
   go mod tidy
   ```

#### 问题：接口实现错误

**错误示例**:

```txt
cannot use userRepo (type *user_impl.UserRepo) as type repository.UserRepository
```

**解决方案**:

1. **检查接口实现**:
   确保实现了接口的所有方法

2. **检查方法签名**:
   确保方法签名完全匹配

3. **使用接口检查**:

   ```go
   var _ repository.UserRepository = (*user_impl.UserRepo)(nil)
   ```

## 运行时问题

### 应用启动问题

#### 问题：端口被占用

**错误信息**:

```txt
bind: address already in use
```

**解决方案**:

1. **查找占用端口的进程**:

   ```bash
   lsof -i :8080
   ```

2. **杀死进程**:

   ```bash
   kill -9 <PID>
   ```

3. **使用不同端口**:

   ```bash
   export SERVER_PORT=8081
   go run cmd/main.go
   ```

#### 问题：依赖注入失败

**错误信息**:

```txt
panic: wire: no provider found for type *service.UserService
```

**解决方案**:

1. **检查 provider 配置**:
   确保所有服务都在 `wire/provider.go` 中定义

2. **重新生成 Wire 代码**:

   ```bash
   cd wire
   wire
   ```

3. **检查依赖链**:
   确保所有依赖都有对应的 provider

### API 问题

#### 问题：404 Not Found

**解决方案**:

1. **检查路由注册**:
   确保路由在 `router.go` 中正确注册

2. **检查 URL 路径**:
   确保请求的 URL 与注册的路由匹配

3. **启用详细日志**:

   ```go
   gin.SetMode(gin.DebugMode)
   ```

#### 问题：500 Internal Server Error

**解决方案**:

1. **查看日志**:
   检查应用日志获取详细错误信息

2. **添加错误处理**:

   ```go
   if err != nil {
       log.Printf("Error: %v", err)
       c.JSON(500, gin.H{"error": err.Error()})
       return
   }
   ```

3. **使用调试器**:
   在 VS Code 中设置断点进行调试

## 测试问题

### 测试运行失败

#### 问题：测试找不到包

**解决方案**:

1. **确保测试文件在正确位置**:

   ```txt
   test/
   ├── unit/
   └── integration/
   ```

2. **检查测试导入**:
   确保导入路径正确

3. **运行特定测试**:

   ```bash
   go test ./test/unit/user_test.go
   ```

#### 问题：Mock 对象问题

**解决方案**:

1. **生成 Mock**:

   ```bash
   go install github.com/golang/mock/mockgen@latest
   mockgen -source=internal/domain/user/repository/user_repository.go -destination=test/mocks/user_repository_mock.go
   ```

2. **检查 Mock 接口**:
   确保 Mock 实现了正确的接口

## 性能问题

### 响应慢

**诊断步骤**:

1. **启用性能分析**:

   ```go
   import _ "net/http/pprof"
   
   go func() {
       log.Println(http.ListenAndServe("localhost:6060", nil))
   }()
   ```

2. **查看性能报告**:

   ```bash
   go tool pprof http://localhost:6060/debug/pprof/profile
   ```

3. **检查数据库查询**:
   添加查询日志

4. **添加缓存**:
   对频繁查询的数据添加缓存

### 内存泄漏

**解决方案**:

1. **检查 Goroutine 泄漏**:

   ```bash
   go tool pprof http://localhost:6060/debug/pprof/goroutine
   ```

2. **检查内存使用**:

   ```bash
   go tool pprof http://localhost:6060/debug/pprof/heap
   ```

3. **确保资源正确释放**:

   ```go
   defer db.Close()
   defer cancel() // context
   ```

## 调试技巧

### 使用日志调试

```go
// 添加详细日志
logger.Debug("函数开始", "function", "CreateUser", "params", fmt.Sprintf("%+v", req))

// 记录中间状态
logger.Debug("验证通过", "email", req.Email)

// 记录错误详情
logger.Error("创建失败", "error", err.Error(), "stack", fmt.Sprintf("%+v", err))
```

### 使用断点调试

1. **在 VS Code 中设置断点**
2. **按 F5 启动调试**
3. **逐步执行代码**

### 使用 Delve 命令行调试

```bash
# 启动调试
dlv debug cmd/main.go

# 设置断点
(dlv) break main.main
(dlv) break internal/app/user/service.(*UserService).CreateUser

# 运行
(dlv) continue

# 查看变量
(dlv) print user
(dlv) locals
```

## 常用命令快速参考

### 项目管理

```bash
# 初始化项目
go mod init base-gin

# 整理依赖
go mod tidy

# 下载依赖
go mod download

# 查看依赖
go list -m all
```

### 构建和运行

```bash
# 格式化代码
go fmt ./...

# 检查代码
go vet ./...

# 运行测试
go test ./...

# 构建项目
go build ./...

# 运行应用
go run cmd/main.go
```

### Wire 相关

```bash
# 生成依赖注入代码
cd wire && wire && cd ..

# 清理生成的代码
rm wire/wire_gen.go
```

### 调试相关

```bash
# 使用 Delve 调试
dlv debug cmd/main.go

# 查看性能
go tool pprof http://localhost:6060/debug/pprof/profile
```

如果遇到其他问题，请检查：

1. **Go 版本兼容性**
2. **依赖版本是否最新**
3. **环境变量配置**
4. **文件权限**
5. **网络连接**

如果问题仍然存在，可以在项目仓库中提交 Issue 或查看现有的 Issue。
