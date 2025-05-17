package constant

// 订单状态
const (
	// 待支付
	OrderStatusPending = iota + 1
	// 已支付
	OrderStatusPaid
	// 已发货
	OrderStatusShipped
	// 已完成
	OrderStatusCompleted
	// 已取消
	OrderStatusCancelled
	// 已退款
	OrderStatusRefunded
)

// 订单状态描述
var OrderStatusDesc = map[int]string{
	OrderStatusPending:   "待支付",
	OrderStatusPaid:      "已支付",
	OrderStatusShipped:   "已发货",
	OrderStatusCompleted: "已完成",
	OrderStatusCancelled: "已取消",
	OrderStatusRefunded:  "已退款",
}

// 支付方式
const (
	// 微信支付
	PaymentMethodWechat = iota + 1
	// 支付宝
	PaymentMethodAlipay
)

// 支付方式描述
var PaymentMethodDesc = map[int]string{
	PaymentMethodWechat: "微信支付",
	PaymentMethodAlipay: "支付宝",
}
