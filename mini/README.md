## 微信小程序登陆接口
	
+ 根据 appid secret 和 code 获取session_key 和 openid

## 微信小程序校验接口
	
+ 获取前端数据rawData和signature
+ rawData和session_key通过sha1加密后与signature进行合法性校验 

## 微信小程序数据解密

+ 获取前端数据encryptedData和iv
+ session_key为aeskey
+ 通过base64解码后，进行aes CBC模式PKCS7解密