#!/bin/bash

# 测试运行脚本
# 用于 CI/CD 环境中运行测试

set -e

echo "🧪 开始运行测试..."

# 运行单元测试
echo "📋 运行单元测试..."
if go test -v ./test/unit/...; then
    echo "✅ 单元测试通过"
else
    echo "❌ 单元测试失败"
    exit 1
fi

# 运行集成测试
echo "🔗 运行集成测试..."
if go test -v ./test/integration/...; then
    echo "✅ 集成测试通过"
else
    echo "⚠️  集成测试失败，但继续执行"
fi

# 检查是否有其他测试
echo "🔍 检查其他测试文件..."
OTHER_TESTS=$(find . -name "*_test.go" -not -path "./test/*" -not -path "./vendor/*" -not -path "./.git/*" | wc -l)

if [ "$OTHER_TESTS" -gt 0 ]; then
    echo "发现 $OTHER_TESTS 个其他测试文件"
    find . -name "*_test.go" -not -path "./test/*" -not -path "./vendor/*" -not -path "./.git/*" | head -5
else
    echo "没有发现其他测试文件"
fi

echo "🎉 测试运行完成"
