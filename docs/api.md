# API 接口文档

## 基础信息

- **Base URL**: `http://localhost:8080`
- **API 版本**: `v1`
- **API 前缀**: `/api/v1`

## 健康检查

### GET /health

检查服务健康状态。

**请求示例：**

```bash
curl http://localhost:8080/health
```

**响应示例：**

```json
{
  "status": "ok"
}
```

## 用户管理

### GET /api/v1/users

获取所有用户列表。

**请求示例：**

```bash
curl http://localhost:8080/api/v1/users
```

**响应示例：**

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
    }
  ]
}
```

### GET /api/v1/users/{id}

根据 ID 获取单个用户。

**路径参数：**

- `id` (integer): 用户 ID

**请求示例：**

```bash
curl http://localhost:8080/api/v1/users/1
```

**成功响应 (200)：**

```json
{
  "data": {
    "id": 1,
    "name": "张三",
    "email": "zhangsan@example.com"
  }
}
```

**用户不存在 (404)：**

```json
{
  "error": "用户不存在"
}
```

### POST /api/v1/users

创建新用户。

**请求体：**

```json
{
  "name": "string",     // 用户名，2-50 个字符
  "email": "string",    // 邮箱地址，必须是有效格式
  "password": "string"  // 密码，至少 6 个字符
}
```

**请求示例：**

```bash
curl -X POST http://localhost:8080/api/v1/users \
  -H "Content-Type: application/json" \
  -d '{
    "name": "新用户",
    "email": "newuser@example.com",
    "password": "password123"
  }'
```

**成功响应 (201)：**

```json
{
  "data": {
    "id": 4,
    "name": "新用户",
    "email": "newuser@example.com"
  },
  "message": "用户创建成功"
}
```

**验证失败 (400)：**

```json
{
  "error": "邮箱格式不正确"
}
```

### PUT /api/v1/users/{id}

更新用户信息。

**路径参数：**

- `id` (integer): 用户 ID

**请求体：**

```json
{
  "name": "string",   // 用户名，2-50 个字符
  "email": "string"   // 邮箱地址，必须是有效格式
}
```

**请求示例：**

```bash
curl -X PUT http://localhost:8080/api/v1/users/1 \
  -H "Content-Type: application/json" \
  -d '{
    "name": "更新后的用户名",
    "email": "updated@example.com"
  }'
```

**成功响应 (200)：**

```json
{
  "data": {
    "id": 1,
    "name": "更新后的用户名",
    "email": "updated@example.com"
  },
  "message": "用户更新成功"
}
```

### DELETE /api/v1/users/{id}

删除用户。

**路径参数：**

- `id` (integer): 用户 ID

**请求示例：**

```bash
curl -X DELETE http://localhost:8080/api/v1/users/1
```

**成功响应 (200)：**

```json
{
  "message": "用户删除成功"
}
```

**用户不存在 (404)：**

```json
{
  "error": "用户不存在"
}
```

## 错误处理

### 常见错误码

- `400 Bad Request`: 请求参数错误或验证失败
- `404 Not Found`: 资源不存在
- `500 Internal Server Error`: 服务器内部错误

### 错误响应格式

```json
{
  "error": "错误描述信息"
}
```

## 验证规则

### 用户字段验证

- **name**:
  - 必填
  - 长度：2-50 个字符
  - 不能为空或只包含空格

- **email**:
  - 必填
  - 必须是有效的邮箱格式
  - 在系统中必须唯一

- **password** (仅创建时需要):
  - 必填
  - 最少 6 个字符
