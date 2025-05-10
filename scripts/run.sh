#!/bin/bash
# Linux/Mac shell脚本，提供常见操作的快捷方式

# 确保目录存在
mkdir -p bin logs reports exports

# 显示帮助
function show_help {
    echo "用法: ./run.sh [命令]"
    echo "可用的命令:"
    echo "  run        - 运行应用"
    echo "  build      - 构建应用"
    echo "  clean      - 清理构建产物"
    echo "  test       - 运行测试"
    echo "  db-init    - 初始化数据库"
    echo "  db-seed    - 添加测试数据"
    echo "  db-verify  - 验证数据库连接"
    echo "  migrate    - 迁移文件到MinIO"
    echo "  help       - 显示帮助信息"
}

# 无参数时显示帮助
if [ -z "$1" ]; then
    show_help
    exit 0
fi

# 处理命令
case "$1" in
    run)
        echo "正在运行应用..."
        go run cmd/main.go
        ;;
    build)
        echo "正在构建应用..."
        go build -o bin/shop-go cmd/main.go
        go build -o bin/migrate cmd/migrate/main.go
        echo "构建完成! 可执行文件在bin目录中"
        ;;
    clean)
        echo "正在清理..."
        rm -rf bin/*
        rm -rf logs/*.log
        rm -rf reports/*
        rm -rf exports/*
        echo "清理完成!"
        ;;
    test)
        echo "正在运行测试..."
        go test -v ./...
        ;;
    db-init)
        echo "正在初始化数据库..."
        echo "请确保MySQL已启动，并且已配置正确的数据库用户名和密码"
        read -p "按Enter键继续..."
        mysql -u root -p < database/schema.sql
        ;;
    db-seed)
        echo "正在添加测试数据..."
        mysql -u root -p < database/init.sql
        ;;
    db-verify)
        echo "正在验证数据库连接..."
        cd database && go run verify_connection.go
        cd ..
        ;;
    migrate)
        echo "正在迁移文件到MinIO..."
        go run cmd/migrate/main.go
        ;;
    help)
        show_help
        ;;
    *)
        echo "未知命令: $1"
        show_help
        exit 1
        ;;
esac

exit 0 