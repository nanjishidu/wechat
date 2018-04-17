package mp

import (
	"errors"
	"fmt"
	"sync"

	mpcore "gopkg.in/chanxuehong/wechat.v2/mp/core"
	mpjssdk "gopkg.in/chanxuehong/wechat.v2/mp/jssdk"
	mpmedia "gopkg.in/chanxuehong/wechat.v2/mp/media"
	mpmenu "gopkg.in/chanxuehong/wechat.v2/mp/menu"
	mpcustom "gopkg.in/chanxuehong/wechat.v2/mp/message/custom"
	mptemplate "gopkg.in/chanxuehong/wechat.v2/mp/message/template"
	mpoauth2 "gopkg.in/chanxuehong/wechat.v2/mp/oauth2"
	mpqrcode "gopkg.in/chanxuehong/wechat.v2/mp/qrcode"
	mpuser "gopkg.in/chanxuehong/wechat.v2/mp/user"
	"gopkg.in/chanxuehong/wechat.v2/oauth2"
)

//定义微信公众号结构
type WeMp struct {
	lock     *sync.RWMutex
	mpServer map[string]*MpServer
	mpConfig map[string]map[string]string
}

//定义每个微信公众号结构
type MpServer struct {
	server            *mpcore.Server
	accessTokenServer *mpcore.DefaultAccessTokenServer
	ticketServer      *mpjssdk.DefaultTicketServer
}

func NewWeMp(mc map[string]map[string]string) *WeMp {
	return &WeMp{
		lock:     new(sync.RWMutex),
		mpConfig: mc,
		mpServer: make(map[string]*MpServer),
	}
}
func (srv *WeMp) SetAppConfig(appId string, appConfig map[string]string) {
	srv.lock.Lock()
	defer srv.lock.Unlock()
	srv.mpConfig[appId] = appConfig
	srv.mpServer[appId] = nil
}

//获取公众号配置
func (srv *WeMp) GetAppConfig(appId string) (map[string]string, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appConfig, ok := srv.mpConfig[appId]; ok {
		return appConfig, nil
	}
	return nil, fmt.Errorf("appConfig is not exist,appId:%s", appId)
}

// Server 用于处理微信服务器的回调请求, 并发安全!
func (srv *WeMp) GetServer(appId string, handler mpcore.Handler) (*mpcore.Server, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appId == "" {
		return nil, errors.New("empty appId")
	}
	if appServer, ok := srv.mpServer[appId]; ok {
		if appServer.server != nil {
			return appServer.server, nil
		}
	} else {
		srv.mpServer[appId] = new(MpServer)
	}
	appConfig, err := srv.GetAppConfig(appId)
	if err != nil {
		return nil, err
	}
	token := appConfig["token"]
	if token == "" {
		return nil, fmt.Errorf("token is not set,appId:%s", appId)
	}
	base64AESKey := appConfig["base64AESKey"]
	if base64AESKey == "" {
		return nil, fmt.Errorf("base64AESKey is not set,appId:%s", appId)
	}

	srv.mpServer[appId].server = mpcore.NewServer("", appId, token, base64AESKey, handler, nil)
	return srv.mpServer[appId].server, nil
}

// access_token 中控服务器
func (srv *WeMp) GetAccessTokenServer(appId string) (*mpcore.DefaultAccessTokenServer, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appId == "" {
		return nil, errors.New("empty appId")
	}
	if appServer, ok := srv.mpServer[appId]; ok {
		if appServer.accessTokenServer != nil {
			return appServer.accessTokenServer, nil
		}
	} else {
		srv.mpServer[appId] = new(MpServer)
	}
	appConfig, err := srv.GetAppConfig(appId)
	if err != nil {
		return nil, err
	}
	appSecret := appConfig["appSecret"]
	if appSecret == "" {
		return nil, fmt.Errorf("appSecret is not set,appId:%s", appId)
	}

	srv.mpServer[appId].accessTokenServer = mpcore.NewDefaultAccessTokenServer(appId, appSecret, nil)
	return srv.mpServer[appId].accessTokenServer, nil
}

//jsapi_ticket 中控服务器
func (srv *WeMp) GetTicketServer(appId string) (*mpjssdk.DefaultTicketServer, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appId == "" {
		return nil, errors.New("empty appId")
	}
	if appServer, ok := srv.mpServer[appId]; ok {
		if appServer.ticketServer != nil {
			return appServer.ticketServer, nil
		}
	} else {
		srv.mpServer[appId] = new(MpServer)
	}
	accessTokenServer, err := srv.GetAccessTokenServer(appId)
	if err != nil {
		return nil, err
	}
	srv.mpServer[appId].ticketServer = mpjssdk.NewDefaultTicketServer(mpcore.NewClient(accessTokenServer, nil))
	return srv.mpServer[appId].ticketServer, nil
}

//创建临时微信二维码
func CreateTempQrcode(accessTokenSrv *mpcore.DefaultAccessTokenServer, sceneId int32, expireSeconds int) (*mpqrcode.TempQrcode, error) {
	tqrcode, err := mpqrcode.CreateTempQrcode(mpcore.NewClient(accessTokenSrv, nil), sceneId, expireSeconds)
	if err != nil {
		return nil, err
	}
	return tqrcode, nil
}

//根据openid 获取用户信息
func GetUserInfo(accessTokenSrv *mpcore.DefaultAccessTokenServer, openid string) (uir *mpuser.UserInfo, err error) {
	return mpuser.Get(mpcore.NewClient(accessTokenSrv, nil), openid, "")
}

//下载临时素材图片
func Download(accessTokenSrv *mpcore.DefaultAccessTokenServer, mediaId, filepath string) (written int64, err error) {
	return mpmedia.Download(mpcore.NewClient(accessTokenSrv, nil), mediaId, filepath)
}

//创建菜单支持二级菜单
func CreateMenu(accessTokenSrv *mpcore.DefaultAccessTokenServer, menu *mpmenu.Menu) error {
	return mpmenu.Create(mpcore.NewClient(accessTokenSrv, nil), menu)
}

//删除菜单
func DelelteMenu(accessTokenSrv *mpcore.DefaultAccessTokenServer) error {
	return mpmenu.Delete(mpcore.NewClient(accessTokenSrv, nil))
}

//推送文章信息
func SendNews(accessTokenSrv *mpcore.DefaultAccessTokenServer, touser string, articles ...mpcustom.Article) error {
	return mpcustom.Send(mpcore.NewClient(accessTokenSrv, nil), mpcustom.NewNews(touser, articles, ""))
}

//发送模版消息
func SendTemplateNews(accessTokenSrv *mpcore.DefaultAccessTokenServer, touser, templateId, uri string, data map[string]mptemplate.DataItem) (int64, error) {
	return mptemplate.Send(mpcore.NewClient(accessTokenSrv, nil),
		mptemplate.TemplateMessage2{
			ToUser:     touser,
			TemplateId: templateId,
			URL:        uri,
			Data:       data,
		})
}

//生成网页授权地址.
func GetAuthCodeUrl(appId, uri string) string {
	return mpoauth2.AuthCodeURL(appId, uri, "snsapi_userinfo", "STATE")
}

// 微信网页授权
// 根据code 获取微信用户信息
func GetUserInfoByCode(appId, appSecret, code string) (info *mpoauth2.UserInfo, err error) {
	token, err := getWxOauth2TokenByCode(appId, appSecret, code)
	if err != nil {
		return nil, errors.New("get mp oauth2 token is failed")
	}
	valid, err := mpoauth2.Auth(token.AccessToken, token.OpenId, nil)
	if err != nil || !valid {
		return nil, errors.New("token is invalid")
	}
	return mpoauth2.GetUserInfo(token.AccessToken, token.OpenId, mpoauth2.LanguageZhCN, nil)
}

//根据code 获取access_token
func getWxOauth2TokenByCode(appId, appSecret, code string) (token *oauth2.Token, err error) {
	var cli = &oauth2.Client{
		Endpoint: mpoauth2.NewEndpoint(appId, appSecret),
	}
	token, err = cli.ExchangeToken(code)
	if err != nil {
		return nil, err
	}
	return token, nil
}

//jssdk sign
func WxConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	return mpjssdk.WXConfigSign(jsapiTicket, nonceStr, timestamp, url)
}
