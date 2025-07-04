# Base Gin DDD 项目 Makefile

.PHONY: help build run test clean wire fmt vet deps tools check-updates update-deps build-compress build-release

# 默认目标
help:
	@echo "可用的命令:"
	@echo "  build        		- 构建应用"
	@echo "  build-release		- 构建发布版本（优化体积）"
	@echo "  build-compress 	- 构建并使用 UPX 压缩应用（可能在 macOS 上因为签名问题无法运行）"
	@echo "  run          		- 运行应用"
	@echo "  test         		- 运行测试"
	@echo "  wire         		- 生成 Wire 依赖注入代码"
	@echo "  fmt          		- 格式化代码"
	@echo "  vet          		- 检查代码"
	@echo "  deps         		- 安装依赖"
	@echo "  tools        		- 安装开发工具"
	@echo "  check-updates		- 检查依赖更新"
	@echo "  update-deps  		- 更新所有依赖"
	@echo "  clean        		- 清理构建文件"

# 安装依赖
deps:
	@echo "安装项目依赖..."
	go mod download
	go mod tidy

# 安装工具
tools:
	@echo "安装开发工具..."
	go install github.com/google/wire/cmd/wire@latest
	go install golang.org/x/tools/gopls@latest
	go install github.com/go-delve/delve/cmd/dlv@latest

# 检查依赖更新
check-updates:
	@echo "检查依赖更新..."
	@echo "当前 Go 版本:" && go version
	@echo "项目要求版本:" && grep 'go ' go.mod
	@echo "可更新的依赖:"
	@go list -m -u all | grep "\[" || echo "所有依赖都是最新的"

# 更新依赖
update-deps:
	@echo "更新依赖..."
	go get -u github.com/gin-gonic/gin@latest
	go get -u github.com/google/wire@latest
	go get -u ./...
	go mod tidy
	@echo "依赖更新完成"

# 生成 Wire 代码
wire:
	@echo "生成 Wire 依赖注入代码..."
	cd wire && wire

# 格式化代码
fmt:
	@echo "格式化代码..."
	go fmt ./...

# 检查代码
vet:
	@echo "检查代码..."
	go vet ./...

# 构建应用
build: wire fmt vet
	@echo "构建应用..."
	go build -o bin/app cmd/main.go

# 构建发布版本（优化体积）
build-release: wire fmt vet
	@echo "构建发布版本（优化体积）..."
	go build -ldflags="-s -w" -o bin/app cmd/main.go
	@echo "发布版本构建完成"

# 运行应用
run: wire
	@echo "运行应用..."
	go run cmd/main.go

# 运行测试
test:
	@echo "运行测试..."
	go test ./...

# 清理构建文件
clean:
	@echo "清理构建文件..."
	rm -f bin/app

# 构建并使用 UPX 压缩应用
build-compress: wire fmt vet
	@echo "构建并压缩应用..."
	go build -o bin/app cmd/main.go
	upx --best --lzma --force-macos bin/app
	@echo "应用已压缩"
