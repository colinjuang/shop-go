.PHONY: run build clean test db migrate

# 默认目标：运行应用
run:
	go run cmd/main.go

# 构建应用
build:
	go build -o bin/shop-go cmd/main.go
	go build -o bin/migrate cmd/migrate/main.go

# 清理构建产物
clean:
	rm -rf bin/*
	rm -rf logs/*.log
	rm -rf reports/*
	rm -rf exports/*

# 运行测试
test:
	go test -v ./...

# 初始化数据库
db-init:
	@echo "正在初始化数据库..."
	@echo "请确保MySQL已启动，并且已配置正确的数据库用户名和密码"
	@read -p "按任意键继续..." KEY
	mysql -u root -p < database/schema.sql

# 添加测试数据
db-seed:
	@echo "正在添加测试数据..."
	mysql -u root -p < database/init.sql

# 验证数据库连接
db-verify:
	cd database && go run verify_connection.go

# 迁移文件到MinIO
migrate:
	go run cmd/migrate/main.go

# 帮助信息
help:
	@echo "可用的命令:"
	@echo "  make run        - 运行应用"
	@echo "  make build      - 构建应用"
	@echo "  make clean      - 清理构建产物"
	@echo "  make test       - 运行测试"
	@echo "  make db-init    - 初始化数据库"
	@echo "  make db-seed    - 添加测试数据"
	@echo "  make db-verify  - 验证数据库连接"
	@echo "  make migrate    - 迁移文件到MinIO"
	@echo "  make help       - 显示帮助信息" 