package paysdk

import (
	"net/http"
	"errors"
)
func (apc *AliPayClient)GetNotificationParam(req *http.Request)(*NotificationParam, error){
	return getNotificationParam(req,apc.PublicKey)
}
func getNotificationParam(req *http.Request,publicKey []byte)(*NotificationParam, error){
	if req == nil {
		return nil, errors.New("request 参数不能为空")
	}
	np := &NotificationParam{}
	np.AppId = req.FormValue("app_id")
	np.AuthAppId = req.FormValue("auth_app_id")
	np.NotifyId = req.FormValue("notify_id")
	np.NotifyType = req.FormValue("notify_type")
	np.NotifyTime = req.FormValue("notify_time")
	np.TradeNo = req.FormValue("trade_no")
	np.TradeStatus = req.FormValue("trade_status")
	np.TotalAmount = req.FormValue("total_amount")
	np.ReceiptAmount = req.FormValue("receipt_amount")
	np.InvoiceAmount = req.FormValue("invoice_amount")
	np.BuyerPayAmount = req.FormValue("buyer_pay_amount")
	np.SellerId = req.FormValue("seller_id")
	np.SellerEmail = req.FormValue("seller_email")
	np.BuyerId = req.FormValue("buyer_id")
	np.BuyerLogonId = req.FormValue("buyer_logon_id")
	np.FundBillList = req.FormValue("fund_bill_list")
	np.Charset = req.FormValue("charset")
	np.PointAmount = req.FormValue("point_amount")
	np.OutTradeNo = req.FormValue("out_trade_no")
	np.OutBizNo = req.FormValue("out_biz_no")
	np.GmtCreate = req.FormValue("gmt_create")
	np.GmtPayment = req.FormValue("gmt_payment")
	np.GmtRefund = req.FormValue("gmt_refund")
	np.GmtClose = req.FormValue("gmt_close")
	np.Subject = req.FormValue("subject")
	np.Body = req.FormValue("body")
	np.RefundFee = req.FormValue("refund_fee")
	np.Version = req.FormValue("version")
	np.SignType = req.FormValue("sign_type")
	np.Sign = req.FormValue("sign")
	np.PassbackParams = req.FormValue("passback_params")
	np.VoucherDetailList = req.FormValue("voucher_detail_list")

	if len(np.NotifyId) == 0 {
		return nil, errors.New("不是有效的 Notify")
	}

	ok, err := verifySign(req.Form, publicKey)
	if ok == false {
		return nil, err
	}
	return np, err
}