package paysdk

const (
	TIME_FORMAT = "2006-01-02 15:04:05"

	ALI_PAY_TRADE_STATUS_WAIT_BUYER_PAY = "WAIT_BUYER_PAY" // 交易创建，等待买家付款
	ALI_PAY_TRADE_STATUS_TRADE_CLOSED   = "TRADE_CLOSED"   // 未付款交易超时关闭，或支付完成后全额退款
	ALI_PAY_TRADE_STATUS_TRADE_SUCCESS  = "TRADE_SUCCESS"  // 交易支付成功
	ALI_PAY_TRADE_STATUS_TRADE_FINISHED = "TRADE_FINISHED" // 交易结束，不可退款

	ALI_PAY_SANDBOX_API_URL     = "https://openapi.alipaydev.com/gateway.do"
	ALI_PAY_PRODUCTION_API_URL  = "https://openapi.alipay.com/gateway.do"
	ALI_PAY_PRODUCTION_MAPI_URL = "https://mapi.alipay.com/gateway.do"

	FORMAT  = "JSON"
	CHARSET = "utf-8"
	VERSION = "1.0"

	

	RESPONSE_SUFFIX = "_response"
	ERROR_RESPONSE  = "error_response"
	SIGN_NODE_NAME  = "sign"

	SIGN_TYPE_RSA2 = "RSA2"
	SIGN_TYPE_RSA  = "RSA"
)
