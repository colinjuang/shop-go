package utils

import (
	"regexp"
	"testing"
	"time"
)

func TestGenerateOrderNo(t *testing.T) {
	// 测试订单号格式
	t.Run("订单号格式测试", func(t *testing.T) {
		userID := uint64(1234)
		orderNo := GenerateOrderNo(userID)

		// 验证订单号长度
		if len(orderNo) != 25 {
			t.Errorf("订单号长度错误，期望25位，实际%d位", len(orderNo))
		}

		// 验证订单号格式
		pattern := `^ORD\d{14}\d{4}\d{4}$`
		matched, err := regexp.MatchString(pattern, orderNo)
		if err != nil {
			t.Fatalf("正则表达式匹配错误: %v", err)
		}
		if !matched {
			t.Errorf("订单号格式错误: %s", orderNo)
		}

		// 验证前缀
		if orderNo[:3] != OrderPrefix {
			t.Errorf("订单号前缀错误，期望%s，实际%s", OrderPrefix, orderNo[:3])
		}

		// 验证时间部分
		timeStr := orderNo[3:17]
		_, err = time.ParseInLocation(TimeFormat, timeStr, time.Local)
		if err != nil {
			t.Errorf("订单号时间部分格式错误: %s", timeStr)
		}

		// 验证用户ID部分
		userIDStr := orderNo[17:21]
		if userIDStr != "1234" {
			t.Errorf("用户ID部分错误，期望1234，实际%s", userIDStr)
		}
	})

	// 测试订单号唯一性
	t.Run("订单号唯一性测试", func(t *testing.T) {
		userID := uint64(1234)
		orderNos := make(map[string]bool)

		// 生成1000个订单号，验证唯一性
		for i := 0; i < 1000; i++ {
			orderNo := GenerateOrderNo(userID)
			if orderNos[orderNo] {
				t.Errorf("订单号重复: %s", orderNo)
			}
			orderNos[orderNo] = true
		}
	})

	// 测试不同用户ID生成的订单号
	t.Run("不同用户ID测试", func(t *testing.T) {
		userID1 := uint64(1234)
		userID2 := uint64(5678)

		orderNo1 := GenerateOrderNo(userID1)
		orderNo2 := GenerateOrderNo(userID2)

		// 验证用户ID部分不同
		if orderNo1[17:21] == orderNo2[17:21] {
			t.Errorf("不同用户ID生成的订单号用户ID部分相同: %s, %s", orderNo1, orderNo2)
		}
	})

	// 测试时间部分
	t.Run("时间部分测试", func(t *testing.T) {
		userID := uint64(1234)
		orderNo := GenerateOrderNo(userID)

		// 验证时间部分是否接近当前时间
		timeStr := orderNo[3:17]
		orderTime, err := time.ParseInLocation(TimeFormat, timeStr, time.Local)
		if err != nil {
			t.Fatalf("解析时间错误: %v", err)
		}

		// 验证时间差在1秒内
		timeDiff := time.Since(orderTime)
		if timeDiff < 0 || timeDiff > time.Second {
			t.Errorf("订单号时间部分与当前时间相差过大: %v", timeDiff)
		}
	})
}
