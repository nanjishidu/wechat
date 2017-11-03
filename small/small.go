// small.go
package small

import (
	"encoding/json"
	"fmt"
)

type Wx struct {
	AppId     string
	AppSecret string
}

func NewWx(appid, secret string) *Wx {
	return &Wx{
		AppId:     appid,
		AppSecret: secret,
	}
}

//"https://api.weixin.qq.com/sns/jscode2session?
//appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
//根据code获取WxSesstion
func (wx *Wx) GetWxSessionKey(code string) (ws WxSesstion, err error) {
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?"+
		"appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		wx.AppId, wx.AppSecret, code)
	err = HttpGet(uri).ToJson(&ws)
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
func GetWxUserInfo(session_key, encryptedData, iv string) (wui WxUserInfo, err error) {
	plaintext, err := AesCBCDecrypt(session_key, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &wui)
	return
}

//微信加密数据结构
type WxUserInfo struct {
	OpenId    string     `json:"openId"`
	NickName  string     `json:"nickName"`
	Gender    int        `json:"gender"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	Country   string     `json:"country"`
	AvatarUrl string     `json:"avatarUrl"`
	UnionId   string     `json:"unionId"`
	Watermark *Watermark `json:"watermark"` //数据水印( watermark )
}
type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}
type WxSesstion struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrInfo
}
type ErrInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
