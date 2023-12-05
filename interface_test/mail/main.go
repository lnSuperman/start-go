package main

// 电商平台多种支付方式 微信 支付宝 银行卡
// 定义协议 提供 创建订单 支付 查询支付状态  退款

type AliPay struct {
}
type WeChat struct {
}
type Bank struct {
}

type BasePay interface {
	CreateCredit()
	CreditPay()
	GetCreditStatus()
	RePay()
}
