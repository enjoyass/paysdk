package paysdk

import (
	"net/url"
	"net/http"
	"io/ioutil"
	"time"
	"crypto"
	"strings"
	"sort"
	"encoding/base64"
)


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
func NewBizContent() BizContent{
	var bizContent= BizContent{}
	return bizContent
}
func (apc *AliPayClient)ProcessUrlValue(bizContent BizContent) (requestBody string, err error){
	var param =url.Values{}
	param.Add("app_id",apc.app_id)
	param.Add("method", apc.method)
	param.Add("notify_url", apc.notify_url)
	param.Add("format", FORMAT)
	param.Add("charset", CHARSET)
	param.Add("sign_type", apc.sign_type)
	param.Add("timestamp", time.Now().Format(TIME_FORMAT))
	param.Add("version", VERSION)
	param.Add("biz_content",bizContent.ToString())

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

func signWithPKCS1v15(param url.Values, privateKey []byte, hash crypto.Hash) (string, error) {
	if param == nil {
		param = make(url.Values, 0)
	}

	var paramList = make([]string, 0, 0)
	for key := range param {
		var value = strings.TrimSpace(param.Get(key))
		if len(value) > 0 {
			paramList = append(paramList, key+"="+value)
		}
	}
	sort.Strings(paramList)
	var src = strings.Join(paramList, "&")
	sign, err := SignPKCS1v15ByPemByte([]byte(src), privateKey, hash)
	if err != nil {
		return "", err
	}
	signed := base64.StdEncoding.EncodeToString(sign)
	return signed, nil
}

func (apc *AliPayClient) VerifySign(data url.Values) (bool, error) {
	return verifySign(data, apc.PublicKey)
}

func verifySign(data url.Values, key []byte) (bool, error) {
	sign := data.Get("sign")
	signType := data.Get("sign_type")

	var keys = make([]string, 0, 0)
	for key, value := range data {
		if key == "sign" || key == "sign_type" {
			continue
		}
		if len(value) > 0 {
			keys = append(keys, key)
		}
	}

	sort.Strings(keys)

	var pList = make([]string, 0, 0)
	for _, key := range keys {
		var value = strings.TrimSpace(data.Get(key))
		if len(value) > 0 {
			pList = append(pList, key+"="+value)
		}
	}
	var s = strings.Join(pList, "&")

	return verifyData([]byte(s), signType, sign, key)
}

func verifyData(data []byte, signType, signed string, key []byte) (bool, error) {
	sign, err := base64.StdEncoding.DecodeString(signed)
	if err != nil {
		return false, err
	}

	if signType == SIGN_TYPE_RSA {
		err = VerifyPKCS1v15ByPemByte(data, sign, key, crypto.SHA1)
	} else {
		err = VerifyPKCS1v15ByPemByte(data, sign, key, crypto.SHA256)
	}
	if err != nil {
		return false, err
	}
	return true, nil
}