package paysdk

type AliPayClient struct {
	app_id      string
	method      string
	charset     string
	version     string
	format      string
	sign_type   string
	notify_url  string
	biz_content string
}

func NewAliPayClient() *AliPayClient {
	return &AliPayClient{}
}
