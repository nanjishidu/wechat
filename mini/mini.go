package mini

import (
	"encoding/json"
	"fmt"
)

type WeMini struct {
	AppId     string
	AppSecret string
}

func NewWeMini(appid, secret string) *WeMini {
	return &WeMini{
		AppId:     appid,
		AppSecret: secret,
	}
}

//"https://api.weixin.qq.com/sns/jscode2session?
//appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
//根据code获取WxSesstion
func (wm *WeMini) GetSessionKey(code string) (sesstionInfo SesstionInfo, err error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		wm.AppId, wm.AppSecret, code)
	err = HttpGet(uri).ToJson(&sesstionInfo)
	return
}

//检验signature是否相同
func CheckSignature(signature, session_key, rawData string) bool {
	if signature == Sha1(rawData+session_key) {
		return true
	}
	return false
}

//根据seesion_key,加密数据encryptedData和向量偏移量iv获取微信用户信息 主要是 敏感信息
func GetUserInfo(session_key, encryptedData, iv string) (wui UserInfo, err error) {
	plaintext, err := AesCBCDecrypt(session_key, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &wui)
	return
}

//获取微信用户绑定的手机号
func GetPhoneNumber(session_key, encryptedData, iv string) (phoneNumber PhoneNumber, err error) {
	plaintext, err := AesCBCDecrypt(session_key, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &phoneNumber)
	return
}

//获取用户过去三十天微信运动步数
func GetWeRunData(session_key, encryptedData, iv string) (weRunData WeRunData, err error) {
	plaintext, err := AesCBCDecrypt(session_key, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &weRunData)
	return
}
