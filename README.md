# 智付SDK V2

### 初始化

```
accessId := "3011"
accessPrivateKeyPath := "../config/Test_AccessPrivateKey.cer"
platformPublicKeyPath := "../config/Test_PlatformPublicKey.cer"
prodEnv := false // prodEnv 为true请求鲲鹏的生产环境，为false请求鲲鹏的测试环境，请悉知
devMode := false // devMode 为true会打印接口请求和响应内容, 依然是请求鲲鹏的生产环境，请悉知
client, err := globebill.NewClient(accessId, accessPrivateKeyPath, platformPublicKeyPath, prodEnv, devMode)
```

### 打赏

赞赏多少都是您的心意，感谢您的支持！打赏时烦请备注一下您的github账号，以便添加感谢名单

<img src="./image/微信收款码.jpg" height="300"> <img src="./image/支付宝收款码.jpg" height="300">

### 感谢名单

| Benefactor | Channel | Amount | Time |
|------------|---------|--------|------|
