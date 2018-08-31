package paysdk

const (
	TIME_FORMAT = "2006-01-02 15:04:05"

	ALI_PAY_TRADE_STATUS_WAIT_BUYER_PAY = "WAIT_BUYER_PAY"
	ALI_PAY_TRADE_STATUS_TRADE_CLOSED   = "TRADE_CLOSED"
	ALI_PAY_TRADE_STATUS_TRADE_SUCCESS  = "TRADE_SUCCESS" 
	ALI_PAY_TRADE_STATUS_TRADE_FINISHED = "TRADE_FINISHED"

	ALI_PAY_SANDBOX_API_URL     = "https://openapi.alipaydev.com/gateway.do"
	ALI_PAY_API_URL  = "https://openapi.alipay.com/gateway.do?charset=utf-8"
	ALI_PAY_MAPI_URL = "https://mapi.alipay.com/gateway.do"

	FORMAT  = "JSON"
	CHARSET = "utf-8"
	VERSION = "1.0"
	

	RESPONSE_SUFFIX = "_response"
	ERROR_RESPONSE  = "error_response"
	SIGN_NODE_NAME  = "sign"

	SIGN_TYPE_RSA2 = "RSA2"
	SIGN_TYPE_RSA  = "RSA"
)