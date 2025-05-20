package order

import (
	"fmt"
	"sync/atomic"
	"time"
)

const (
	// 订单号前缀
	OrderPrefix = "ORD"
	// 订单号时间格式
	TimeFormat = "20060102150405"
	// 随机数长度
	RandomLength = 4
)

var (
	// 用于生成唯一序号的计数器
	counter uint64
)

// GenerateOrderNo 生成唯一的订单号
// 格式: ORD + 时间戳(14位) + 用户ID(4位) + 序号(4位)
// 例如: ORD2024031512345612345678
func GenerateOrderNo(userID uint64) string {
	// 获取当前时间
	now := time.Now()

	// 获取并递增计数器
	seq := atomic.AddUint64(&counter, 1) % 10000

	// 格式化用户ID为4位数字
	userIDStr := fmt.Sprintf("%04d", userID%10000)

	// 格式化序号为4位数字
	seqStr := fmt.Sprintf("%04d", seq)

	// 组合订单号
	return fmt.Sprintf("%s%s%s%s",
		OrderPrefix,
		now.Format(TimeFormat),
		userIDStr,
		seqStr,
	)
}
