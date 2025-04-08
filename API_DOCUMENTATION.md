# WallosApp API 文档

## 基础信息
- 基础URL: `http://localhost:8080`
- 所有API响应格式均为JSON
- 需要认证的接口需要在请求头中添加 `Authorization: Bearer <token>`

## 健康检查

### 检查服务状态
- **描述**: 检查服务是否正常运行
- **路由**: `GET /health`
- **参数**: 无
- **响应示例**:
```json
{
    "status": "ok"
}
```

## 用户认证

### 用户注册
- **描述**: 创建新用户账号
- **路由**: `POST /api/register`
- **请求参数**:
```json
{
    "username": "string",  // 用户名，必填
    "email": "string",     // 邮箱，必填，需要符合邮箱格式
    "password": "string"   // 密码，必填，最少6位
}
```
- **响应示例**:
```json
{
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "is_active": true
}
```
- **错误响应**:
  - 400: 参数验证失败
  - 409: 用户名或邮箱已存在
  - 500: 服务器内部错误

### 用户登录
- **描述**: 用户登录获取访问令牌
- **路由**: `POST /api/login`
- **请求参数**:
```json
{
    "username": "string",  // 用户名，必填
    "password": "string"   // 密码，必填
}
```
- **响应示例**:
```json
{
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
        "id": 1,
        "username": "testuser",
        "email": "test@example.com",
        "created_at": "2024-01-01T00:00:00Z",
        "is_active": true
    }
}
```
- **错误响应**:
  - 400: 参数验证失败
  - 401: 用户名或密码错误
  - 500: 服务器内部错误

### 获取用户资料
- **描述**: 获取当前登录用户的详细信息
- **路由**: `GET /api/profile`
- **认证**: 需要Bearer Token
- **请求参数**: 无
- **响应示例**:
```json
{
    "id": 1,
    "username": "testuser",
    "email": "test@example.com",
    "created_at": "2024-01-01T00:00:00Z",
    "is_active": true
}
```
- **错误响应**:
  - 401: 未认证或token无效
  - 404: 用户不存在
  - 500: 服务器内部错误

## 错误响应格式
所有错误响应都遵循以下格式：
```json
{
    "error": "错误描述信息"
}
```

## 认证说明
1. 登录成功后，服务器会返回一个JWT token
2. 需要在后续请求的Header中添加：
   ```
   Authorization: Bearer <token>
   ```
3. Token有效期为24小时
4. 如果token过期或无效，需要重新登录获取新token 