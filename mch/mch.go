package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/core"
)

var (
	MchCommonConfig *CommonConfig
	MchClient       *core.Client
	MchTLSClient    *core.Client
)

type CommonConfig struct {
	AppId  string `json:"app_id"`  //微信分配的公众号id
	MchId  string `json:"mch_id"`  //微信支付分配的商户号
	ApiKey string `json:"api_key"` //微信支付分配的商户号
}

func (m *CommonConfig) SetAppId(app_id string) {
	m.AppId = app_id
}
func (m *CommonConfig) SetMchId(mch_id string) {
	m.MchId = mch_id
}
func (m *CommonConfig) SetApiKey(api_key string) {
	m.ApiKey = api_key
}

//构建通用配置
func InitWechatMch(app_id, mch_id, api_key string, certKeys ...string) error {
	if app_id == "" || mch_id == "" {
		return errors.New("appid or mch_id is null")
	}
	MchCommonConfig = &CommonConfig{
		AppId:  app_id,
		MchId:  mch_id,
		ApiKey: api_key,
	}
	initMchClient()
	if len(certKeys) > 0 {
		if err := initMchTLSClient(certKeys...); err != nil {
			return err
		}
	}
	return nil
}

func initMchClient() {
	MchClient = core.NewClient(MchCommonConfig.AppId, MchCommonConfig.MchId, MchCommonConfig.ApiKey, nil)
}
func initMchTLSClient(certKeys ...string) error {
	if len(certKeys) < 2 {
		return errors.New("certKeys verification failed")
	}
	cli, err := core.NewTLSHttpClient(certKeys[0], certKeys[1])
	if err == nil {
		MchTLSClient = core.NewClient(MchCommonConfig.AppId, MchCommonConfig.MchId, MchCommonConfig.ApiKey, cli)
		return nil
	}
	return err

}
