package response

// PaymentResponse represents the payment response
type PaymentResponse struct {
	PaymentID string `json:"paymentID"`
	AppID     string `json:"appID"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
