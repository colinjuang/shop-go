@echo off
REM 设置数据库连接参数
set DB_USER=root
set DB_PASS=123456
set DB_HOST=127.0.0.1
set DB_PORT=3306
set DB_NAME=flower_shop

echo 正在验证数据库连接...
go run verify_connection.go

if %ERRORLEVEL% NEQ 0 (
    echo 连接失败! 请检查数据库配置和服务状态。
    echo 确保MySQL服务已启动，且配置信息正确。
) else (
    echo 验证完成!
)

pause 