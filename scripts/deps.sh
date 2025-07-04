#!/bin/bash

# Base Gin 依赖管理脚本

set -e

# 颜色定义
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

echo -e "${BLUE}=== Base Gin 依赖管理工具 ===${NC}"

# 检查 Go 版本
check_go_version() {
    echo -e "\n${YELLOW}📋 检查 Go 版本...${NC}"
    local go_version=$(go version | awk '{print $3}' | sed 's/go//')
    local mod_version=$(grep "go " go.mod | awk '{print $2}')
    
    echo "系统 Go 版本: $go_version"
    echo "项目要求版本: $mod_version"
    
    if [[ "$go_version" != "$mod_version"* ]]; then
        echo -e "${YELLOW}⚠️  Go 版本不匹配，建议更新 go.mod 中的版本${NC}"
    else
        echo -e "${GREEN}✅ Go 版本匹配${NC}"
    fi
}

# 检查依赖更新
check_updates() {
    echo -e "\n${YELLOW}🔍 检查依赖更新...${NC}"
    
    # 检查直接依赖
    echo -e "\n${BLUE}直接依赖:${NC}"
    go list -m -u $(go list -m -f '{{if not .Indirect}}{{.Path}}{{end}}' all | grep -v "^$")
    
    # 检查所有可更新的依赖
    echo -e "\n${BLUE}所有可更新的依赖:${NC}"
    outdated=$(go list -m -u all | grep "\[" | wc -l)
    if [ "$outdated" -gt 0 ]; then
        echo -e "${YELLOW}发现 $outdated 个可更新的依赖:${NC}"
        go list -m -u all | grep "\["
    else
        echo -e "${GREEN}✅ 所有依赖都是最新的${NC}"
    fi
}

# 更新依赖
update_deps() {
    echo -e "\n${YELLOW}🔄 更新依赖...${NC}"
    
    # 更新主要依赖
    echo "更新 Gin 框架..."
    go get -u github.com/gin-gonic/gin@latest
    
    echo "更新 Wire..."
    go get -u github.com/google/wire@latest
    
    # 更新所有依赖
    echo "更新所有依赖..."
    go get -u ./...
    
    # 整理依赖
    echo "整理依赖..."
    go mod tidy
    
    echo -e "${GREEN}✅ 依赖更新完成${NC}"
}

# 重新生成 Wire 代码
regenerate_wire() {
    echo -e "\n${YELLOW}🔧 重新生成 Wire 代码...${NC}"
    
    if [ -d "wire" ]; then
        cd wire
        if command -v wire >/dev/null 2>&1; then
            wire
            echo -e "${GREEN}✅ Wire 代码生成完成${NC}"
        else
            echo -e "${RED}❌ Wire 工具未安装，请运行: go install github.com/google/wire/cmd/wire@latest${NC}"
        fi
        cd ..
    else
        echo -e "${RED}❌ 未找到 wire 目录${NC}"
    fi
}

# 验证构建
verify_build() {
    echo -e "\n${YELLOW}🏗️  验证构建...${NC}"
    
    if go build -v ./...; then
        echo -e "${GREEN}✅ 构建成功${NC}"
    else
        echo -e "${RED}❌ 构建失败${NC}"
        exit 1
    fi
}

# 运行测试
run_tests() {
    echo -e "\n${YELLOW}🧪 运行测试...${NC}"
    
    if go test ./...; then
        echo -e "${GREEN}✅ 测试通过${NC}"
    else
        echo -e "${RED}❌ 测试失败${NC}"
        exit 1
    fi
}

# 显示依赖漏洞
check_vulnerabilities() {
    echo -e "\n${YELLOW}🔒 检查依赖漏洞...${NC}"
    
    if command -v govulncheck >/dev/null 2>&1; then
        govulncheck ./...
    else
        echo -e "${YELLOW}⚠️  govulncheck 未安装，可运行以下命令安装:${NC}"
        echo "go install golang.org/x/vuln/cmd/govulncheck@latest"
    fi
}

# 显示依赖树
show_dependency_tree() {
    echo -e "\n${YELLOW}🌳 依赖树:${NC}"
    go mod graph | head -20
    echo "..."
    echo "完整依赖树请运行: go mod graph"
}

# 清理依赖缓存
clean_cache() {
    echo -e "\n${YELLOW}🧹 清理依赖缓存...${NC}"
    go clean -modcache
    echo -e "${GREEN}✅ 缓存清理完成${NC}"
}

# 主菜单
show_menu() {
    echo -e "\n${BLUE}请选择操作:${NC}"
    echo "1) 检查 Go 版本"
    echo "2) 检查依赖更新"
    echo "3) 更新所有依赖"
    echo "4) 重新生成 Wire 代码"
    echo "5) 验证构建"
    echo "6) 运行测试"
    echo "7) 检查安全漏洞"
    echo "8) 显示依赖树"
    echo "9) 清理依赖缓存"
    echo "a) 执行完整检查更新流程"
    echo "q) 退出"
    echo -n -e "${YELLOW}请输入选择: ${NC}"
}

# 完整流程
full_update() {
    echo -e "\n${BLUE}=== 执行完整检查更新流程 ===${NC}"
    check_go_version
    check_updates
    
    echo -n -e "\n${YELLOW}是否继续更新依赖? [y/N]: ${NC}"
    read -r response
    if [[ "$response" =~ ^[Yy]$ ]]; then
        update_deps
        regenerate_wire
        verify_build
        run_tests
        echo -e "\n${GREEN}🎉 完整更新流程完成！${NC}"
    else
        echo -e "${YELLOW}取消更新${NC}"
    fi
}

# 检查是否在项目根目录
if [ ! -f "go.mod" ]; then
    echo -e "${RED}❌ 请在项目根目录运行此脚本${NC}"
    exit 1
fi

# 主循环
while true; do
    show_menu
    read -r choice
    
    case $choice in
        1) check_go_version ;;
        2) check_updates ;;
        3) update_deps ;;
        4) regenerate_wire ;;
        5) verify_build ;;
        6) run_tests ;;
        7) check_vulnerabilities ;;
        8) show_dependency_tree ;;
        9) clean_cache ;;
        a) full_update ;;
        q) echo -e "${GREEN}再见！${NC}"; exit 0 ;;
        *) echo -e "${RED}无效选择，请重试${NC}" ;;
    esac
done
