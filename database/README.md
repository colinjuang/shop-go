# 花店系统数据库初始化

这个目录包含了初始化花店系统数据库的SQL脚本。

## 文件说明

- `schema.sql`: 包含创建数据库和所有表的SQL语句，以及基础的示例数据
- `init.sql`: 包含额外的测试数据和示例记录
- `verify_connection.go`: Go程序，用于验证数据库连接和表结构
- `verify.bat`: Windows批处理文件，用于在Windows系统上运行验证脚本
- `verify.sh`: Shell脚本，用于在Linux/Mac系统上运行验证脚本

## 如何使用

### 1. 初始化数据库

#### 方法一：使用命令行

1. 确保已安装MySQL并且服务已启动
2. 打开命令行终端，登录到MySQL
   ```bash
   mysql -u root -p
   ```
3. 执行schema.sql创建数据库和表
   ```bash
   source /path/to/schema.sql
   ```
4. 执行init.sql添加额外的测试数据（可选）
   ```bash
   source /path/to/init.sql
   ```

#### 方法二：使用MySQL客户端工具

1. 使用MySQL Workbench、Navicat、HeidiSQL等客户端工具连接到MySQL服务器
2. 创建一个新的查询窗口
3. 打开并执行`schema.sql`脚本
4. 打开并执行`init.sql`脚本（可选）

### 2. 验证数据库连接

#### Windows系统

1. 如有需要，编辑`verify.bat`文件中的数据库连接参数
2. 双击运行`verify.bat`或在命令行中执行：
   ```
   verify.bat
   ```

#### Linux/Mac系统

1. 如有需要，编辑`verify.sh`文件中的数据库连接参数
2. 赋予脚本执行权限：
   ```bash
   chmod +x verify.sh
   ```
3. 运行脚本：
   ```bash
   ./verify.sh
   ```

#### 直接使用Go

1. 设置环境变量：
   ```bash
   # Windows
   set DB_USER=root
   set DB_PASS=密码
   
   # Linux/Mac
   export DB_USER=root
   export DB_PASS=密码
   ```
2. 运行Go程序：
   ```bash
   go run verify_connection.go
   ```

## 数据库结构

系统包含以下表：

- `users`: 用户表
- `addresses`: 收货地址表
- `categories`: 商品分类表
- `products`: 商品表
- `banners`: 首页轮播图表
- `promotions`: 促销活动表
- `cart_items`: 购物车表
- `orders`: 订单表
- `order_items`: 订单商品表

## 测试数据

初始化脚本会创建以下测试数据：

- 花店的基本商品分类（鲜花、绿植、花篮、礼盒等）
- 一些示例商品（玫瑰花束、百合花束等）
- 首页轮播图和促销活动
- 一个测试用户和示例订单

## 注意事项

- 执行`schema.sql`会创建一个名为`flower_shop`的新数据库。如果已存在同名数据库，将会继续使用该数据库而不删除现有数据
- 所有SQL语句使用了`IF NOT EXISTS`子句，避免在表已存在的情况下报错
- 默认的数据库字符集为`utf8mb4`，支持完整的Unicode字符集（包括Emoji表情符号）
- 可以根据需要修改`configs/config.yaml`中的数据库连接配置 