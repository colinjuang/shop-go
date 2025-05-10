@echo off
REM Windows批处理文件，提供常见操作的快捷方式
SETLOCAL EnableDelayedExpansion

REM 确保bin目录存在
if not exist "bin" mkdir bin

if "%1"=="" goto :help

if "%1"=="run" (
    echo 正在运行应用...
    go run cmd/main.go
    goto :EOF
)

if "%1"=="build" (
    echo 正在构建应用...
    go build -o bin/shop-go.exe cmd/main.go
    go build -o bin/migrate.exe cmd/migrate/main.go
    echo 构建完成! 可执行文件在bin目录中
    goto :EOF
)

if "%1"=="clean" (
    echo 正在清理...
    if exist "bin\*" del /Q bin\*
    if exist "logs\*.log" del /Q logs\*.log
    if exist "reports\*" del /Q reports\*
    if exist "exports\*" del /Q exports\*
    echo 清理完成!
    goto :EOF
)

if "%1"=="test" (
    echo 正在运行测试...
    go test -v ./...
    goto :EOF
)

if "%1"=="db-init" (
    echo 正在初始化数据库...
    echo 请确保MySQL已启动，并且已配置正确的数据库用户名和密码
    pause
    mysql -u root -p < database/schema.sql
    goto :EOF
)

if "%1"=="db-seed" (
    echo 正在添加测试数据...
    mysql -u root -p < database/init.sql
    goto :EOF
)

if "%1"=="db-verify" (
    echo 正在验证数据库连接...
    cd database && go run verify_connection.go
    cd ..
    goto :EOF
)

if "%1"=="migrate" (
    echo 正在迁移文件到MinIO...
    go run cmd/migrate/main.go
    goto :EOF
)

:help
echo 用法: run.bat [命令]
echo 可用的命令:
echo   run        - 运行应用
echo   build      - 构建应用
echo   clean      - 清理构建产物
echo   test       - 运行测试
echo   db-init    - 初始化数据库
echo   db-seed    - 添加测试数据
echo   db-verify  - 验证数据库连接
echo   migrate    - 迁移文件到MinIO
echo   help       - 显示帮助信息

:EOF
ENDLOCAL 