package mini

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"sync"

	mpcore "gopkg.in/chanxuehong/wechat.v2/mp/core"
)

type WeMini struct {
	lock       *sync.RWMutex
	miniConfig map[string]map[string]string
	miniServer map[string]*MiniServer
}

type MiniServer struct {
	accessTokenServer *mpcore.DefaultAccessTokenServer
}

func NewWeMini(mc map[string]map[string]string) *WeMini {
	return &WeMini{
		lock:       new(sync.RWMutex),
		miniConfig: mc,
		miniServer: make(map[string]*MiniServer),
	}
}
func (srv *WeMini) SetAppConfig(appId string, appConfig map[string]string) {
	srv.lock.Lock()
	defer srv.lock.Unlock()
	srv.miniConfig[appId] = appConfig
	srv.miniServer[appId] = nil
}
func (srv *WeMini) GetAppConfig(appId string) (map[string]string, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appConfig, ok := srv.miniConfig[appId]; ok {
		return appConfig, nil
	}
	return nil, fmt.Errorf("appConfig is not exist,appId:%s", appId)
}

// access_token 中控服务器
func (srv *WeMini) GetAccessTokenServer(appId string) (*mpcore.DefaultAccessTokenServer, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appId == "" {
		return nil, errors.New("empty appId")
	}
	if appServer, ok := srv.miniServer[appId]; ok {
		if appServer.accessTokenServer != nil {
			return appServer.accessTokenServer, nil
		}
	} else {
		srv.miniServer[appId] = new(MiniServer)
	}
	appConfig, err := srv.GetAppConfig(appId)
	if err != nil {
		return nil, err
	}
	appSecret := appConfig["appSecret"]
	if appSecret == "" {
		return nil, fmt.Errorf("appSecret is not set,appId:%s", appId)
	}

	srv.miniServer[appId].accessTokenServer = mpcore.NewDefaultAccessTokenServer(appId, appSecret, nil)
	return srv.miniServer[appId].accessTokenServer, nil
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

//发送模版消息
func SendTemplateNews(accessTokenSrv *mpcore.DefaultAccessTokenServer, touser, templateId, formId, page string,
	data map[string]interface{}) (resp ErrInfo, err error) {
	var incompleteURL = "https://api.weixin.qq.com/cgi-bin/message/wxopen/template/send?access_token="
	token, err := accessTokenSrv.Token()
	if err != nil {
		return resp, err
	}
	finalURL := incompleteURL + url.QueryEscape(token)
	sendData := map[string]interface{}{
		"touser":      touser,
		"template_id": templateId,
		"form_id":     formId,
		"page":        page,
		"data":        data,
	}

	b, err := json.Marshal(sendData)
	if err != nil {
		return resp, err
	}
	err = HttpPost(finalURL, bytes.NewReader(b)).ToJson(&resp)
	return
}

//获取二维码
//接口A: 适用于需要的码数量较少的业务场景
//接口B：适用于需要的码数量极多的业务场景
//接口C：适用于需要的码数量较少的业务场景
func GetWxAcode(accessTokenSrv *mpcore.DefaultAccessTokenServer, rtype string, wa *WxAcode) (resp []byte, err error) {
	var incompleteURL string
	switch rtype {
	case "A":
		incompleteURL = "https://api.weixin.qq.com/wxa/getwxacode?access_token="
	case "B":
		incompleteURL = "https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token="
	case "C":
		incompleteURL = "https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token="
	default:
		return nil, errors.New("type is not A,B,C")
	}
	token, err := accessTokenSrv.Token()
	if err != nil {
		return nil, err
	}
	finalURL := incompleteURL + url.QueryEscape(token)
	b, err := json.Marshal(wa)
	if err != nil {
		return nil, err
	}
	resp, err = HttpPost(finalURL, bytes.NewReader(b)).Bytes()
	return
}
