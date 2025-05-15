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
│   ├── handler/       # HTTP handlers
│   ├── middleware/    # Middleware functions
│   ├── model/         # Domain models
│   ├── pkg/           # Internal shared packages
│   │   ├── redis/     # Redis client and tools
│   │   ├── minio/     # MinIO client and tools
│   │   └── logger/    # Zap logger configuration
│   ├── repository/    # Data access layer
│   ├── service/       # Business logic
│   └── server/        # Server setup
├── pkg/               # Shareable packages 
├── uploads/           # File uploads (created at runtime)
├── logs/              # Application logs
├── reports/           # Generated reports
└── exports/           # Exported data
```

## Requirements

- Go 1.22 or higher
- MySQL
- Redis (for caching and rate limiting)
- MinIO (for object storage)

## Configuration

Edit the `configs/config.yaml` file to customize:

- Server settings
- Database connection
- Redis configuration
- MinIO configuration
- JWT keys
- WeChat credentials
- Upload settings
- Logger configuration

## Quick Start

1. Clone the repository
2. Set up the database using the scripts in `database/`
3. Install and configure a Redis server
4. Install and configure a MinIO server
5. Configure the application (configs/config.yaml)
6. Run the application:

```bash
go run cmd/main.go
```

## Migrating Files to MinIO

To migrate existing files from local storage to MinIO:

```bash
go run cmd/migrate/main.go
```

Options:
- `--config` - Path to config file (default: configs/config.yaml)
- `--file` - Migrate a single file (optional)

## API Endpoints

### Home
- `GET /mall-api/home/banners` - Get banners for homepage
- `GET /mall-api/home/categories` - Get top categories
- `GET /mall-api/home/promotions` - Get promotions
- `GET /mall-api/home/recommend` - Get recommended products
- `GET /mall-api/home/hot` - Get hot products

### Categories
- `GET /mall-api/categories` - Get all categories
- `GET /mall-api/categories/:id/subs` - Get subcategories

### Products
- `GET /mall-api/products` - Get paginated products
- `GET /mall-api/products/:id` - Get product details

### Reports and Exports
- `GET /mall-api/reports/catalog` - Generate PDF product catalog
- `GET /mall-api/reports/export` - Export products to CSV
- `GET /mall-api/order/:id/invoice` - Generate order invoice (requires auth)

### User
- `GET /mall-api/users/login` - WeChat login
- `GET /mall-api/users/info` - Get user info (requires auth)
- `POST /mall-api/users/update` - Update user info (requires auth)

### Address
- `POST /mall-api/address/add` - Add new address (requires auth)
- `GET /mall-api/address/list` - Get address list (requires auth)
- `GET /mall-api/address/:id` - Get address details (requires auth)
- `POST /mall-api/address/:id/update` - Update address (requires auth)
- `GET /mall-api/address/:id/delete` - Delete address (requires auth)

### Cart
- `GET /mall-api/cart/add` - Add to cart (requires auth)
- `GET /mall-api/cart/list` - Get cart items (requires auth)
- `GET /mall-api/cart/update` - Update item status (requires auth)
- `GET /mall-api/cart/update-all` - Update all items status (requires auth)
- `GET /mall-api/cart/delete` - Delete cart item (requires auth)

### Order
- `GET /mall-api/order/detail` - Get order details (requires auth)
- `GET /mall-api/order/address` - Get order address (requires auth)
- `POST /mall-api/order/submit` - Submit order (requires auth)
- `GET /mall-api/order/pay` - Get WeChat pay info (requires auth)
- `GET /mall-api/order/pay/status` - Check payment status (requires auth)
- `GET /mall-api/order/list` - Get order list (requires auth)

### Upload
- `POST /mall-api/upload` - Upload file (requires auth)
- `POST /mall-api/upload/batch` - Upload multiple files (requires auth)
- `POST /mall-api/upload/delete` - Delete uploaded file (requires auth)

## Key Features

### Redis Caching
- Product and category caching
- Rate limiting for API endpoints
- Distributed locking for order processing

### MinIO Object Storage
- Efficient file storage for uploaded images
- Auto-migration from local storage
- PDF report and CSV export generation
- File caching for generated reports

### Zap Logging
- Structured logging with different log levels
- Console and JSON encoding support
- Rotating log files support 