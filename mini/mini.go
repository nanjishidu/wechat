package mini

import (
	"encoding/json"
	"fmt"
	"sync"
)

type WeMini struct {
	lock       *sync.RWMutex
	miniConfig map[string]map[string]string
}

func NewWeMini(mc map[string]map[string]string) *WeMini {
	return &WeMini{
		lock:       new(sync.RWMutex),
		miniConfig: mc,
	}
}
func (srv *WeMini) SetAppConfig(appId string, appConfig map[string]string) {
	srv.lock.Lock()
	defer srv.lock.Unlock()
	srv.miniConfig[appId] = appConfig
}
func (srv *WeMini) GetAppConfig(appId string) (map[string]string, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appConfig, ok := srv.miniConfig[appId]; ok {
		return appConfig, nil
	}
	return nil, fmt.Errorf("appConfig is not exist,appId:%s", appId)
}

//"https://api.weixin.qq.com/sns/jscode2session?appid=APPID&secret=SECRET&js_code=JSCODE&grant_type=authorization_code
//根据code获取SessionInfo
func (srv *WeMini) GetSessionInfo(appId string, code string) (sesstionInfo SesstionInfo, err error) {
	appConfig, err := srv.GetAppConfig(appId)
	if err != nil {
		return sesstionInfo, err
	}
	appSecret := appConfig["appSecret"]
	if appSecret == "" {
		return sesstionInfo, fmt.Errorf("appSecret is not set,appId:%s", appId)
	}
	uri := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code",
		appId, appSecret, code)
	err = HttpGet(uri).ToJson(&sesstionInfo)
	return sesstionInfo, err
}

//检验signature是否相同
func CheckSignature(signature, session_key, rawData string) bool {
	if signature == Sha1(rawData+session_key) {
		return true
	}
	return false
}

//根据seesion_key,加密数据encryptedData和向量偏移量iv获取微信用户信息 主要是 敏感信息
func GetUserInfo(session_key, encryptedData, iv string) (userInfo UserInfo, err error) {
	plaintext, err := AesCBCDecrypt(session_key, encryptedData, iv)
	if err != nil {
		return
	}
	err = json.Unmarshal(plaintext, &userInfo)
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
