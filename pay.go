package paysdk

import (
	"fmt"
	"net/url"
	"net/http"
	"io/ioutil"
	"time"
	"crypto"
	"strings"
	"sort"
	"encoding/base64"
)

type AliPayClient struct {
	app_id      string
	method      string
	notify_url  string
	PrivateKey  []byte
	PublicKey   []byte
	sign_type   string
	Client *http.Client
}

func NewAliPayClient(appId,method,sign_type,notify_url,privateKey,publicKey string) *AliPayClient {
	return &AliPayClient{
		app_id:appId,
		notify_url:notify_url,
		method:method,
		sign_type:sign_type,
		PrivateKey:parsePrivateKey(privateKey),
		PublicKey:parsePublicKey(publicKey),
		Client:http.DefaultClient,
	}
}
func NewModel() Model{
	var model= Model{}
	return model
}
func (apc *AliPayClient)ProcessUrlValue(model Model) (requestBody string, err error){
	fmt.Println(model.ToString())
	var param =url.Values{}
	param.Add("app_id",apc.app_id)
	param.Add("method", apc.method)
	param.Add("notify_url", apc.notify_url)
	param.Add("format", FORMAT)
	param.Add("charset", CHARSET)
	param.Add("sign_type", apc.sign_type)
	param.Add("timestamp", time.Now().Format(TIME_FORMAT))
	param.Add("version", VERSION)
	param.Add("biz_content",model.ToString())

	var hash crypto.Hash
	if apc.sign_type == SIGN_TYPE_RSA {
		hash = crypto.SHA1
	} else {
		hash = crypto.SHA256
	}
	sign, err := signWithPKCS1v15(param, apc.PrivateKey, hash)
	if err != nil {
		return "", err
	}
	param.Add("sign", sign)
	return param.Encode(), nil

}
func signWithPKCS1v15(param url.Values, privateKey []byte, hash crypto.Hash) (s string, err error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var pList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	sort.Strings(pList)
	var src = strings.Join(pList, "&")
	sign, err := SignPKCS1v15ByPemByte([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	signed := base64.StdEncoding.EncodeToString(sign)
	return signed, nil
}

func (apc *AliPayClient)sdkExcute(requestBody string)([]byte,error){
	buf := strings.NewReader(requestBody)
	req, err := http.NewRequest("POST", ALI_PAY_API_URL, buf)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-type", "application/x-www-form-urlencoded;charset=UTF-8")

	resp, err := apc.Client.Do(req)
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
