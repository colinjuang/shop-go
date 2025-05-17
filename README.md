# WeChat Mini Program Mall Backend API (Shop-Go)

A WeChat mini program mall backend API implemented with the Gin framework.

## Project Structure

The project is organized in a clean architecture:

```
├── cmd/               # Application entry points
│   └── migrate/       # MinIO file migration tool
├── configs/           # Configuration files
├── database/          # Database initialization and scripts
├── internal/          # Internal packages
│   ├── config/        # Configuration code
│   ├── dto/               # Data Transfer Objects
│   ├── handler/       # HTTP handlers
│   ├── middleware/    # Middleware functions
│   ├── model/         # Domain models
│   ├── pkg/           # Internal shared packages
│   │   ├── redis/     # Redis client and tools
│   │   ├── minio/     # MinIO client and tools
│   │   └── logger/    # Zap logger configuration
│   ├── repository/    # Data access layer
│   ├── router/        # API routes definition
│   ├── service/       # Business logic
│   └── server/        # Server setup
├── scripts/           # Utility scripts
├── uploads/           # File uploads (created at runtime)
├── logs/              # Application logs (created at runtime)
├── reports/           # Generated reports (created at runtime)
└── exports/           # Exported data (created at runtime)
```

## Requirements

- Go 1.23 or higher
- MySQL
- Redis
- MinIO

### Main Dependencies

- Gin Web Framework (v1.10.0) - HTTP web framework
- GORM (v1.26.1) - ORM library for database operations
- Zap (v1.27.0) - Structured logging
- Viper (v1.20.1) - Configuration management
- JWT (v4.5.2) - JSON Web Token authentication
- Redis (v9.8.0) - Redis client
- MinIO (v7.0.91) - Object storage
- Wechat SDK (v2.1.8) - WeChat integration

## Configuration

Edit the `configs/config.yaml` file to customize your application settings:

```yaml
server:
  port: ":8001"
  environment: "development"

database:
  host: "localhost"
  port: "3306"
  username: "root"
  password: "your-password"
  dbname: "flower_shop"

redis:
  host: "localhost"
  port: "6379"
  password: "your-redis-password"
  db: 0
  prefix: "flower_shop:"

minio:
  endpoint: "localhost:9000"
  access_key: "minioadmin"
  secret_key: "minioadmin"
  use_ssl: false
  bucket: "flower-shop"
  location: "us-east-1"

jwt:
  secret: "your-secret-key-here"
  expires_in: 72 # hours

wechat:
  app_id: "your-app-id-here"
  app_secret: "your-app-secret-here"

upload:
  save_path: "./uploads"
  max_size: 5242880 # 5MB

logger:
  level: "info"  # debug, info, warn, error, fatal
  encoding: "json"  # json or console
  output_path: "stdout"  # stdout or file path
```

## Quick Start

1. Clone the repository
2. Configure the application in `configs/config.yaml`
3. Set up a MySQL database
4. Install and configure Redis
5. Install and configure MinIO
6. Run the application:

```bash
go run cmd/main.go
```

For development with hot reload:
```bash
go install github.com/cosmtrek/air@latest
air
```

## Migrating Files to MinIO

To migrate existing files from local storage to MinIO:

```bash
go run cmd/migrate/main.go
```

Configuration options are read from the `configs/config.yaml` file.

## API Endpoints

All API endpoints are prefixed with `/api`.

### Banner
- `GET /api/banners` - Get banners for homepage

### Categories
- `GET /api/categories` - Get all categories
- `GET /api/categories/:id/subs` - Get subcategories

### Products
- `GET /api/products` - Get paginated products
- `GET /api/products/:id` - Get product details

### Promotions
- `GET /api/promotions` - Get promotions

### User Authentication
- `POST /api/login` - Login with username/password
- `GET /api/wechat/login` - WeChat login

### User Management
- `GET /api/users/info` - Get user info (requires auth)
- `POST /api/users/update` - Update user info (requires auth)

### Address
- `POST /api/address/add` - Add new address (requires auth)
- `GET /api/address/list` - Get address list (requires auth)
- `GET /api/address/:id` - Get address details (requires auth)
- `POST /api/address/:id/update` - Update address (requires auth)
- `GET /api/address/:id/delete` - Delete address (requires auth)

### Cart
- `GET /api/cart/add` - Add to cart (requires auth)
- `GET /api/cart/list` - Get cart items (requires auth)
- `GET /api/cart/update` - Update item status (requires auth)
- `GET /api/cart/update-all` - Update all items status (requires auth)
- `GET /api/cart/delete` - Delete cart item (requires auth)

### Order
- `GET /api/order/detail` - Get order details (requires auth)
- `GET /api/order/address` - Get order address (requires auth)
- `POST /api/order/submit` - Submit order (requires auth)
- `GET /api/order/pay` - Get WeChat pay info (requires auth)
- `GET /api/order/pay/status` - Check payment status (requires auth)
- `GET /api/order/list` - Get order list (requires auth)

### Upload
- `POST /api/upload` - Upload file (requires auth)
- `POST /api/upload/batch` - Upload multiple files (requires auth)
- `POST /api/upload/delete` - Delete uploaded file (requires auth)

## Key Features

### Redis Caching
- Product and category caching
- Rate limiting for API endpoints
- Distributed locking for order processing

### MinIO Object Storage
- Efficient file storage for uploaded images
- Auto-migration from local storage to MinIO

### Zap Logging
- Structured logging with different log levels
- Console and JSON encoding support
- Rotating log files support 