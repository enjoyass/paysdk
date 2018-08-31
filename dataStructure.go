package paysdk

import (
	"net/http"
)

type AliPayClient struct {
	app_id     string //商户ID
	method     string //支付方法，app,wap
	notify_url string //异步通知地址
	PrivateKey []byte //pck1 私钥
	PublicKey  []byte //pck1 公钥
	sign_type  string //签名方式RSA或者RSA2,默认RSA2
	Client     *http.Client
}

type NotificationParam struct {
	AuthAppId         string `json:"auth_app_id"`         // App Id
	NotifyTime        string `json:"notify_time"`         // 通知时间
	NotifyType        string `json:"notify_type"`         // 通知类型
	NotifyId          string `json:"notify_id"`           // 通知校验ID
	AppId             string `json:"app_id"`              // 开发者的app_id
	Charset           string `json:"charset"`             // 编码格式
	Version           string `json:"version"`             // 接口版本
	SignType          string `json:"sign_type"`           // 签名类型
	Sign              string `json:"sign"`                // 签名
	TradeNo           string `json:"trade_no"`            // 支付宝交易号
	OutTradeNo        string `json:"out_trade_no"`        // 商户订单号
	OutBizNo          string `json:"out_biz_no"`          // 商户业务号
	BuyerId           string `json:"buyer_id"`            // 买家支付宝用户号
	BuyerLogonId      string `json:"buyer_logon_id"`      // 买家支付宝账号
	SellerId          string `json:"seller_id"`           // 卖家支付宝用户号
	SellerEmail       string `json:"seller_email"`        // 卖家支付宝账号
	TradeStatus       string `json:"trade_status"`        // 交易状态
	TotalAmount       string `json:"total_amount"`        // 订单金额
	ReceiptAmount     string `json:"receipt_amount"`      // 实收金额
	InvoiceAmount     string `json:"invoice_amount"`      // 开票金额
	BuyerPayAmount    string `json:"buyer_pay_amount"`    // 付款金额
	PointAmount       string `json:"point_amount"`        // 集分宝金额
	RefundFee         string `json:"refund_fee"`          // 总退款金额
	Subject           string `json:"subject"`             // 总退款金额
	Body              string `json:"body"`                // 商品描述
	GmtCreate         string `json:"gmt_create"`          // 交易创建时间
	GmtPayment        string `json:"gmt_payment"`         // 交易付款时间
	GmtRefund         string `json:"gmt_refund"`          // 交易退款时间
	GmtClose          string `json:"gmt_close"`           // 交易结束时间
	FundBillList      string `json:"fund_bill_list"`      // 支付金额信息
	PassbackParams    string `json:"passback_params"`     // 回传参数
	VoucherDetailList string `json:"voucher_detail_list"` // 优惠券信息
}
type FundBill struct {
	FundChannel string  `json:"fund_channel"`       // 交易使用的资金渠道，详见 支付渠道列表
	Amount      string  `json:"amount"`             // 该支付工具类型所使用的金额
	RealAmount  float64 `json:"real_amount,string"` // 渠道实际付款金额
}
type VoucherDetail struct{
	Id                 string `json:"id"`                  // 券id
	Name               string `json:"name"`                // 券名称
	Type               string `json:"type"`                // 当前有三种类型： ALIPAY_FIX_VOUCHER - 全场代金券, ALIPAY_DISCOUNT_VOUCHER - 折扣券, ALIPAY_ITEM_VOUCHER - 单品优惠
	Amount             string `json:"amount"`              // 优惠券面额，它应该会等于商家出资加上其他出资方出资
	MerchantContribute string `json:"merchant_contribute"` // 商家出资（特指发起交易的商家出资金额）
	OtherContribute    string `json:"other_contribute"`    // 其他出资方出资金额，可能是支付宝，可能是品牌商，或者其他方，也可能是他们的一起出资
	Memo               string `json:"memo"`                // 优惠券备注信息
}
type  AliPayTradeQueryResponse  struct{
	AliPayTradeQuery struct {
		Code             string `json:"code"`
		Msg              string `json:"msg"`
		SubCode          string `json:"sub_code"`
		SubMsg           string `json:"sub_msg"`
		AuthTradePayMode string `json:"auth_trade_pay_mode"` // 预授权支付模式，该参数仅在信用预授权支付场景下返回。信用预授权支付：CREDIT_PREAUTH_PAY
		BuyerLogonId     string `json:"buyer_logon_id"`      // 买家支付宝账号
		BuyerPayAmount   string `json:"buyer_pay_amount"`    // 买家实付金额，单位为元，两位小数。
		BuyerUserId      string `json:"buyer_user_id"`       // 买家在支付宝的用户id
		BuyerUserType    string `json:"buyer_user_type"`     // 买家用户类型。CORPORATE:企业用户；PRIVATE:个人用户。
		InvoiceAmount    string `json:"invoice_amount"`      // 交易中用户支付的可开具发票的金额，单位为元，两位小数。
		OutTradeNo       string `json:"out_trade_no"`        // 商家订单号
		PointAmount      string `json:"point_amount"`        // 积分支付的金额，单位为元，两位小数。
		ReceiptAmount    string `json:"receipt_amount"`      // 实收金额，单位为元，两位小数
		SendPayDate      string `json:"send_pay_date"`       // 本次交易打款给卖家的时间
		TotalAmount      string `json:"total_amount"`        // 交易的订单金额
		TradeNo          string `json:"trade_no"`            // 支付宝交易号
		TradeStatus      string `json:"trade_status"`        // 交易状态

		DiscountAmount      string           `json:"discount_amount"`               // 平台优惠金额
		FundBillList        []*FundBill      `json:"fund_bill_list,omitempty"`      // 交易支付使用的资金渠道
		MdiscountAmount     string           `json:"mdiscount_amount"`              // 商家优惠金额
		PayAmount           string           `json:"pay_amount"`                    // 支付币种订单金额
		PayCurrency         string           `json:"pay_currency"`                  // 订单支付币种
		SettleAmount        string           `json:"settle_amount"`                 // 结算币种订单金额
		SettleCurrency      string           `json:"settle_currency"`               // 订单结算币种
		SettleTransRate     string           `json:"settle_trans_rate"`             // 结算币种兑换标价币种汇率
		StoreId             string           `json:"store_id"`                      // 商户门店编号
		StoreName           string           `json:"store_name"`                    // 请求交易支付中的商户店铺的名称
		TerminalId          string           `json:"terminal_id"`                   // 商户机具终端编号
		TransCurrency       string           `json:"trans_currency"`                // 标价币种
		TransPayRate        string           `json:"trans_pay_rate"`                // 标价币种兑换支付币种汇率
		DiscountGoodsDetail string           `json:"discount_goods_detail"`         // 本次交易支付所使用的单品券优惠的商品优惠信息
		IndustrySepcDetail  string           `json:"industry_sepc_detail"`          // 行业特殊信息（例如在医保卡支付业务中，向用户返回医疗信息）。
		VoucherDetailList   []*VoucherDetail `json:"voucher_detail_list,omitempty"` // 本交易支付时使用的所有优惠券信息
	} `json:"alipay_trade_query_response"`
	Sign string `json:"sign"`
}
