# paysdk

### 初始化客户端
```
    aliPayClient:=NewAliPayClient(appId,method,sign_type,notify_url,privateKey,publicKey)
```
### 业务参数
```
    bodyParam :=BizContent()
    bodyParam.Add("product_code","FAST_INSTANT_TRADE_PAY")
    bodyParam.Add("total_amount","0.01")
    bodyParam.Add("subject","1")
    bodyParam.Add("body","我是测试数据")
    bodyParam.Add("out_trade_no","IQJZSRC1YMQB5HU")
```

### 生成签名后的请求参数
```
    urlvalue,_:=aliPayClient.ProcessUrlValue(bodyParam)
```

### 请求下单
```
    data,err:=aliPayClient.sdkExcute(urlvalue)
    if err!=nil {
        t.Error(err)
    }
    fmt.Println(string(data))
```