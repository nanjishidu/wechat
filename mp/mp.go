package mp

import (
	"errors"

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
	openoauth2 "gopkg.in/chanxuehong/wechat.v2/open/oauth2"
)

var (
	MpAccessTokenSrv *mpcore.DefaultAccessTokenServer
	MpTicketSrv      *mpjssdk.DefaultTicketServer
)

//初始化
func initMp(appId, appSecret string) {
	MpAccessTokenSrv = mpcore.NewDefaultAccessTokenServer(appId, appSecret, nil)
	MpTicketSrv = mpjssdk.NewDefaultTicketServer(mpcore.NewClient(MpAccessTokenSrv, nil))
}

//创建临时微信二维码
func CreateTempQrcode(sceneId int32, expireSeconds int) (*mpqrcode.TempQrcode, error) {
	tqrcode, err := mpqrcode.CreateTempQrcode(mpcore.NewClient(MpAccessTokenSrv, nil), sceneId, expireSeconds)
	if err != nil {
		return nil, err
	}
	return tqrcode, nil
}

//生成网页授权地址.
func GetAuthCodeUrl(appId, uri string) string {
	return mpoauth2.AuthCodeURL(appId, uri, "snsapi_userinfo", "STATE")
}

//根据code 获取微信用户信息
func GetUserInfoByCode(appId, appSecret, code string) (info *openoauth2.UserInfo, err error) {
	token, err := getWxOauth2TokenByCode(appId, appSecret, code)
	if err != nil {
		return nil, errors.New("get wx oauth2 token is failed")
	}
	valid, err := mpoauth2.Auth(token.AccessToken, token.OpenId, nil)
	if err != nil || !valid {
		return nil, errors.New("token is invalid")
	}
	return openoauth2.GetUserInfo(token.AccessToken, token.OpenId, "zh_CN", nil)
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

//推送文章信息
func SendNews(touser string, articles ...mpcustom.Article) error {
	return mpcustom.Send(mpcore.NewClient(MpAccessTokenSrv, nil),
		mpcustom.NewNews(touser, articles, ""))
}

//根据openid 获取用户信息
func GetUserInfo(openid string) (uir *mpuser.UserInfo, err error) {
	return mpuser.Get(mpcore.NewClient(MpAccessTokenSrv, nil), openid, "")
}

//jssdk sign
func WxConfigSign(jsapiTicket, nonceStr, timestamp, url string) (signature string) {
	return mpjssdk.WXConfigSign(jsapiTicket, nonceStr, timestamp, url)
}

//下载临时素材图片
func Download(mediaId, filepath string) (written int64, err error) {
	return mpmedia.Download(mpcore.NewClient(MpAccessTokenSrv, nil), mediaId, filepath)
}

//创建菜单支持二级菜单
func CreateMenu(menu *mpmenu.Menu) error {
	return mpmenu.Create(mpcore.NewClient(MpAccessTokenSrv, nil), menu)
}

//删除菜单
func DelelteMenu() error {
	return mpmenu.Delete(mpcore.NewClient(MpAccessTokenSrv, nil))
}

//发送模版消息
func SendTemplateNews(touser, templateId, uri string, data map[string]mptemplate.DataItem) (int64, error) {
	return mptemplate.Send(mpcore.NewClient(MpAccessTokenSrv, nil),
		mptemplate.TemplateMessage2{
			ToUser:     touser,
			TemplateId: templateId,
			URL:        uri,
			Data:       data,
		})
}
