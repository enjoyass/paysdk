package paysdk

import (
	"bytes"
	"net/http"
	"crypto/tls"
	"io/ioutil"
)

type AliPayClient struct {
	url         string
	app_id      string
	method      string
	charset     string
	version     string
	format      string
	PrivateKey  []byte
	PublicKey   []byte
	sign_type   string
	notify_url  string
	biz_content string
	Clinet *http.Client
}

func NewAliPayClient(url,appId,charset,format,sign_type string,privateKey,publicKey []byte) *AliPayClient {
	return &AliPayClient{
		url:url,
		app_id:appId,
		charset:charset,
		format:format,
		sign_type:sign_type,
		PrivateKey:privateKey,
		PublicKey:publicKey,
		version:VERSION,

	}
}
func (apc *AliPayClient)sdkExcute(params *bytes.Buffer)([]byte,error){
	req, err := http.NewRequest("POST", apc.url, params)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
	}
	apc.Clinet = &http.Client{Transport: tr}

	resp, err := apc.Clinet.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	respData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return respData, nil
}
func AlipayTradeAppPayRequest() {

}
