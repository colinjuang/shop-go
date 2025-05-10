package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 从命令行参数获取数据库配置，如果没有则使用默认值
	username := getEnvOrDefault("DB_USER", "root")
	password := getEnvOrDefault("DB_PASS", "123456")
	host := getEnvOrDefault("DB_HOST", "127.0.0.1")
	port := getEnvOrDefault("DB_PORT", "3306")
	dbname := getEnvOrDefault("DB_NAME", "flower_shop")

	// 构建DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		username, password, host, port, dbname)

	// 尝试连接数据库
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("无法创建数据库连接：%v", err)
	}
	defer db.Close()

	// 测试连接
	err = db.Ping()
	if err != nil {
		log.Fatalf("无法连接到数据库：%v", err)
	}

	fmt.Println("数据库连接成功!")

	// 检查数据库表是否存在
	var tableCount int
	tables := []string{"users", "addresses", "categories", "products", "banners", "promotions", "cart_items", "orders", "order_items"}

	err = db.QueryRow("SELECT COUNT(*) FROM information_schema.tables WHERE table_schema = ? AND table_name IN (?, ?, ?, ?, ?, ?, ?, ?, ?)",
		dbname, tables[0], tables[1], tables[2], tables[3], tables[4], tables[5], tables[6], tables[7], tables[8]).Scan(&tableCount)

	if err != nil {
		log.Fatalf("查询表信息时出错：%v", err)
	}

	fmt.Printf("数据库中找到 %d/%d 张表\n", tableCount, len(tables))

	// 检查每个表的记录数
	for _, table := range tables {
		var count int
		err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM %s", table)).Scan(&count)
		if err != nil {
			fmt.Printf("表 %s 查询失败：%v\n", table, err)
		} else {
			fmt.Printf("表 %s 中包含 %d 条记录\n", table, count)
		}
	}
}

// 从环境变量获取值，如果不存在则返回默认值
func getEnvOrDefault(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
