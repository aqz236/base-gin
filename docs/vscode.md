# VS Code 开发环境配置

本文档详细说明了如何配置 VS Code 以获得最佳的 Go 开发体验。

## 目录结构

```text
.vscode/
├── settings.json      # 工作区设置
├── launch.json        # 调试配置
└── extensions.json    # 推荐扩展列表
```

## 推荐扩展

### 必装扩展

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Go | `golang.go` | Go 官方扩展，提供语法高亮、智能感知、调试、测试等 |
| GitLens | `eamodio.gitlens` | Git 增强工具，代码历史、blame、分支可视化 |

### API 开发

| 扩展名 | ID | 说明 |
|--------|----|----- |
| REST Client | `humao.rest-client` | 在 VS Code 中测试 HTTP API |
| Thunder Client | `rangav.vscode-thunder-client` | 类似 Postman 的 API 测试工具 |

### 数据库工具

| 扩展名 | ID | 说明 |
|--------|----|----- |
| SQLTools | `mtxr.sqltools` | 通用数据库连接工具 |
| SQLTools MySQL/MariaDB | `mtxr.sqltools-driver-mysql` | MySQL 数据库驱动 |
| SQLTools PostgreSQL | `mtxr.sqltools-driver-pg` | PostgreSQL 数据库驱动 |

### 容器化开发

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Docker | `ms-azuretools.vscode-docker` | Docker 文件支持和容器管理 |
| Remote - Containers | `ms-vscode-remote.remote-containers` | 在容器中开发 |

### 文档编写

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Markdown All in One | `yzhang.markdown-all-in-one` | Markdown 编写增强 |
| Markdown Preview Enhanced | `shd101wyy.markdown-preview-enhanced` | 增强预览功能 |

### 代码质量

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Better Comments | `aaron-bond.better-comments` | 更好的注释高亮 |
| Document This | `mintlify.document` | 自动生成文档注释 |

### 项目管理

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Project Manager | `alefragnani.project-manager` | 快速切换项目 |
| Makefile Tools | `ms-vscode.makefile-tools` | Makefile 支持 |

### 主题美化

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Dracula Official | `dracula-theme.theme-dracula` | 流行的暗色主题 |
| Material Icon Theme | `pkief.material-icon-theme` | 文件图标主题 |

### 协作工具

| 扩展名 | ID | 说明 |
|--------|----|----- |
| Live Share | `ms-vsliveshare.vsliveshare` | 实时协作编程 |

## 安装扩展

### 自动安装（推荐）

1. 用 VS Code 打开项目：

   ```bash
   code .
   ```

2. VS Code 会自动检测 `.vscode/extensions.json` 并提示安装推荐扩展

3. 点击右下角的"安装"按钮安装所有推荐扩展

### 手动安装

1. 打开命令面板：`Ctrl+Shift+P` (Windows/Linux) 或 `Cmd+Shift+P` (macOS)

2. 输入：`Extensions: Show Recommended Extensions`

3. 选择要安装的扩展

### 命令行安装

```bash
# 安装 Go 扩展
code --install-extension golang.go

# 安装 GitLens
code --install-extension eamodio.gitlens

# 安装 REST Client
code --install-extension humao.rest-client

# 安装 Docker 扩展
code --install-extension ms-azuretools.vscode-docker
```

## 配置说明

### settings.json

主要配置项：

```json
{
  "go.gopath": "",
  "go.goroot": "",
  "go.toolsManagement.autoUpdate": true,
  "go.useLanguageServer": true,
  "go.formatTool": "goimports",
  "go.lintTool": "golangci-lint",
  "go.testTimeout": "30s"
}
```

### launch.json

调试配置支持：

- 启动和调试 main.go
- 附加到正在运行的进程
- 运行当前文件
- 运行和调试测试

### 快捷键

| 功能 | 快捷键 (Windows/Linux) | 快捷键 (macOS) |
|------|----------------------|---------------|
| Go to Definition | `F12` | `F12` |
| Go to References | `Shift+F12` | `Shift+F12` |
| Rename Symbol | `F2` | `F2` |
| Format Document | `Shift+Alt+F` | `Shift+Option+F` |
| Run Tests | `Ctrl+F5` | `Cmd+F5` |
| Debug | `F5` | `F5` |

## 常用功能

### 1. 代码导航

- **跳转到定义**：`F12` 或 `Ctrl+Click`
- **查看引用**：`Shift+F12`
- **搜索符号**：`Ctrl+Shift+O`
- **搜索工作区符号**：`Ctrl+T`

### 2. 代码编辑

- **自动完成**：`Ctrl+Space`
- **快速修复**：`Ctrl+.`
- **重命名符号**：`F2`
- **格式化文档**：`Shift+Alt+F`

### 3. 测试和调试

- **运行测试**：点击测试函数上方的 "run test" 链接
- **调试测试**：点击 "debug test" 链接
- **运行所有测试**：`Ctrl+Shift+P` -> `Go: Test All Packages In Workspace`

### 4. Git 集成

- **查看文件历史**：右键文件 -> `GitLens: Open File History`
- **比较版本**：右键文件 -> `GitLens: Compare with Previous`
- **查看 blame**：右键代码行 -> `GitLens: Toggle Line Blame`

### 5. API 测试

创建 `.http` 文件进行 API 测试：

```http
### 健康检查
GET http://localhost:8080/health

### 获取用户列表
GET http://localhost:8080/api/v1/users

### 创建用户
POST http://localhost:8080/api/v1/users
Content-Type: application/json

{
  "name": "张三",
  "email": "zhangsan@example.com"
}
```

## 故障排除

### Go 扩展无法工作

1. 检查 Go 安装：

   ```bash
   go version
   which go
   ```

2. 重新安装 Go 工具：

   ```text
   Ctrl+Shift+P -> Go: Install/Update Tools
   ```

3. 检查 GOPATH 和 GOROOT 设置

### 智能感知不工作

1. 重启 Go 语言服务器：

   ```text
   Ctrl+Shift+P -> Go: Restart Language Server
   ```

2. 检查 `go.mod` 文件是否正确

3. 运行 `go mod tidy`

### 调试无法启动

1. 检查 `launch.json` 配置
2. 确保 `dlv` 调试器已安装：

   ```bash
   go install github.com/go-delve/delve/cmd/dlv@latest
   ```

## 最佳实践

1. **定期更新扩展**：保持扩展版本最新
2. **使用工作区设置**：团队统一配置
3. **善用代码片段**：提高编码效率
4. **配置自动保存**：避免忘记保存文件
5. **使用 Git 集成**：充分利用版本控制功能

## 相关文档

- [开发指南](development.md)
- [快速开始](quick-start.md)
- [故障排除](troubleshooting.md)
