package mch

import (
	"errors"

	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
)

var (
	MchConfig             *WecMch
	MchClient             *mchcore.Client
	MchTLSClient          *mchcore.Client
	MchUnifiedOrderServer *mchcore.Server
	MchRefundServer       *mchcore.Server
)

type WecMch struct {
	AppId    string //微信分配的公众号id
	MchId    string //微信支付分配的商户号
	ApiKey   string //商户的签名 key
	SubAppId string //可选; 公众号的 sub_appid
	SubMchId string //必选; 商户号 sub_mch_id
}

func NewWeMchClient(appId, mchId, apiKey, subAppId, subMchId string, certKeys ...string) error {
	if appId == "" || mchId == "" || apiKey == "" {
		return errors.New("appId or mchId or apiKey is not set")
	}
	MchConfig = &WecMch{
		AppId:    appId,
		MchId:    mchId,
		ApiKey:   apiKey,
		SubAppId: subAppId,
		SubMchId: subMchId,
	}
	if subAppId != "" {
		MchClient = mchcore.NewSubMchClient(appId, mchId, apiKey, subAppId, subMchId, nil)
	} else {
		MchClient = mchcore.NewClient(appId, mchId, apiKey, nil)
	}
	if len(certKeys) > 0 {
		cli, err := mchcore.NewTLSHttpClient(certKeys[0], certKeys[1])
		if err == nil {
			if subAppId != "" {
				MchTLSClient = mchcore.NewSubMchClient(appId, mchId, apiKey, subAppId, subMchId, cli)
			} else {
				MchTLSClient = mchcore.NewClient(appId, mchId, apiKey, cli)
			}
			return nil
		}
		return err
	} else if len(certKeys) < 2 {
		return errors.New("certKeys verification failed")
	}
	return nil
}
func NewWeMchUnifiedOrderServer(appId, mchId, apiKey, subAppId, subMchId string, handler mchcore.Handler,
	errorHandler mchcore.ErrorHandler) error {
	if subMchId != "" {
		MchUnifiedOrderServer = mchcore.NewSubMchServer("", "", apiKey, subAppId, subMchId, handler, errorHandler)
	} else {
		MchUnifiedOrderServer = mchcore.NewServer(appId, mchId, apiKey, handler, errorHandler)
	}
	return nil
}
func NewWeMchRefundServer(appId, mchId, apiKey, subAppId, subMchId string, handler mchcore.Handler,
	errorHandler mchcore.ErrorHandler) error {
	if subMchId != "" {
		MchRefundServer = mchcore.NewSubMchServer("", "", apiKey, subAppId, subMchId, handler, errorHandler)
	} else {
		MchRefundServer = mchcore.NewServer(appId, mchId, apiKey, handler, errorHandler)
	}
	return nil
}
