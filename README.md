# 微信小程序商城后端 API (Shop-Go)

这是一个基于Gin框架实现的微信小程序商城后端API。

## 项目结构

项目采用清晰的架构方式组织：

```
├── cmd/              # 应用程序入口点
│   └── main.go       # 主应用程序入口
├── configs/          # 配置文件
├── internal/         # 内部包
│   ├── api/          # API层
│   │   └── v1/          # API V1路由定义
│   ├── handler/      # HTTP处理器
│   ├── middleware/   # HTTP中间件
│   ├── request/      # HTTP请求
│   ├── response/     # HTTP响应
│   ├── router/       # API路由定义
│   ├── config/       # 配置代码
│   ├── model/        # 领域模型
│   ├── pkg/          # 内部共享包
│   │   ├── database/ # 数据库客户端和工具
│   │   ├── redis/    # Redis客户端和工具
│   │   ├── minio/    # MinIO客户端和工具
│   │   └── logger/   # Zap日志配置
│   ├── repository/   # 数据访问层
│   ├── service/      # 业务逻辑
│   └── server/       # 服务器设置
├── scripts/          # 实用脚本
├── uploads/          # 文件上传（运行时创建）
├── logs/             # 应用程序日志（运行时创建）
├── reports/          # 生成的报告（运行时创建）
└── exports/          # 导出的数据（运行时创建）
```

## 系统要求

- Go 1.23或更高版本
- MySQL
- Redis（用于缓存和速率限制）
- MinIO（用于对象存储）

## 配置

编辑`configs/config.yaml`文件以自定义：

- 服务器设置
- 数据库连接
- Redis配置
- MinIO配置
- JWT密钥
- 微信凭证
- 上传设置

## 快速开始

1. 克隆仓库
2. 设置数据库
3. 安装并配置Redis服务器
4. 安装并配置MinIO服务器
5. 配置应用程序(configs/config.yaml)
6. 运行应用程序：

```bash
go run cmd/main.go
```

## 将文件迁移到MinIO

要将现有文件从本地存储迁移到MinIO：

```bash
go run cmd/migrate/main.go
```

选项：
- `--config` - 配置文件路径（默认：configs/config.yaml）
- `--file` - 迁移单个文件（可选）

## API端点

### 首页
- `GET /api/banner` - 获取首页轮播图
- `GET /api/category/level1` - 获取顶级分类
- `GET /api/promotion` - 获取促销信息
- `GET /api/product/recommend` - 获取推荐商品
- `GET /api/product/hot` - 获取热门商品

### 分类
- `GET /api/category` - 获取所有分类
- `GET /api/category/:id/subs` - 获取子分类

### 商品
- `GET /api/product` - 分页获取商品
- `GET /api/product/:id` - 获取商品详情

### 报表和导出
- `GET /api/report/catalog` - 生成PDF商品目录
- `GET /api/report/export` - 导出商品到CSV

### 用户
- `GET /api/user/login` - 微信登录
- `GET /api/user/info` - 获取用户信息（需要认证）
- `POST /api/user/update` - 更新用户信息（需要认证）

### 地址
- `POST /api/address/add` - 添加新地址（需要认证）
- `GET /api/address/list` - 获取地址列表（需要认证）
- `GET /api/address/:id` - 获取地址详情（需要认证）
- `POST /api/address/:id/update` - 更新地址（需要认证）
- `GET /api/address/:id/delete` - 删除地址（需要认证）

### 购物车
- `GET /api/cart/add` - 加入购物车（需要认证）
- `GET /api/cart/list` - 获取购物车项目（需要认证）
- `GET /api/cart/update` - 更新项目状态（需要认证）
- `GET /api/cart/update-all` - 更新所有项目状态（需要认证）
- `GET /api/cart/delete` - 删除购物车项目（需要认证）

### 订单
- `GET /api/order/:id/invoice` - 生成订单发票（需要认证）
- `GET /api/order/detail` - 获取订单详情（需要认证）
- `GET /api/order/address` - 获取订单地址（需要认证）
- `POST /api/order/submit` - 提交订单（需要认证）
- `GET /api/order/pay` - 获取支付信息（需要认证）
- `GET /api/order/pay/status` - 检查支付状态（需要认证）
- `GET /api/order/list` - 获取订单列表（需要认证）

### 上传
- `POST /api/upload` - 上传文件（需要认证）
- `POST /api/upload/batch` - 上传多个文件（需要认证）
- `POST /api/upload/delete` - 删除已上传文件（需要认证）

## 主要功能

### Redis缓存
- 商品和分类缓存
- API端点的速率限制
- 订单处理的分布式锁

### MinIO对象存储
- 上传图片的高效文件存储
- 自动从本地存储迁移
- PDF报表和CSV导出生成
- 生成报表的文件缓存 