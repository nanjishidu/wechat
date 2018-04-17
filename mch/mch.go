package mch

import (
	"errors"
	"fmt"
	"sync"

	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
)

type WeMch struct {
	lock      *sync.RWMutex
	mchCore   map[string]*MchCore
	mchConfig map[string]map[string]string
}
type MchCore struct {
	mchClient             *mchcore.Client
	mchTLSClient          *mchcore.Client
	mchUnifiedOrderServer *mchcore.Server
	mchRefundServer       *mchcore.Server
}

func NewWeMch(mc map[string]map[string]string) *WeMch {
	return &WeMch{
		lock:      new(sync.RWMutex),
		mchConfig: mc,
		mchCore:   make(map[string]*MchCore),
	}
}

// appId    微信分配的公众号id
// mchId    微信支付分配的商户号
// apiKey   商户的签名 key
// subAppId 公众号的 sub_appid
// subMchId 商户号 sub_mch_id
// certFile 证书
// keyFile
// notifyUrl 		支付结果通知
// refundNotifyUrl  退款结果通知
// appFinalId 为最终商户appid 如果存在子账户 即子账户appid
func (srv *WeMch) SetAppConfig(appFinalId string, appConfig map[string]string) {
	srv.lock.Lock()
	defer srv.lock.Unlock()
	srv.mchConfig[appFinalId] = appConfig
	srv.mchCore[appFinalId] = nil
}
func (srv *WeMch) GetAppConfig(appFinalId string) (map[string]string, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appConfig, ok := srv.mchConfig[appFinalId]; ok {
		return appConfig, nil
	}
	return nil, fmt.Errorf("appConfig is not exist,appId:%s", appFinalId)
}
func (srv *WeMch) GetAppId(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["appId"], nil
}
func (srv *WeMch) GetMchId(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["mchId"], nil
}
func (srv *WeMch) GetApiKey(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["apiKey"], nil
}
func (srv *WeMch) GetSubAppId(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["subAppId"], nil
}
func (srv *WeMch) GetSubMchId(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["subMchId"], nil
}
func (srv *WeMch) GetNotifyUrl(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["notifyUrl"], nil
}
func (srv *WeMch) GetRefundNotifyUrl(appFinalId string) (string, error) {
	if appFinalId == "" {
		return "", errors.New("empty appId")
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return "", err
	}
	return appConfig["refundNotifyUrl"], nil
}

func (srv *WeMch) GetMchClient(appFinalId string) (*mchcore.Client, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appFinalId == "" {
		return nil, errors.New("empty appId")
	}
	if appCore, ok := srv.mchCore[appFinalId]; ok {
		if appCore.mchUnifiedOrderServer != nil {
			return appCore.mchClient, nil
		}
	} else {
		srv.mchCore[appFinalId] = new(MchCore)
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return nil, err
	}
	appId := appConfig["appId"]
	mchId := appConfig["mchId"]
	if appId == "" || mchId == "" {
		return nil, errors.New("appId or mchId is not set")
	}
	apiKey := appConfig["apiKey"]
	if apiKey == "" {
		return nil, errors.New("apiKey is not set")
	}
	var mchClient *mchcore.Client
	subAppId := appConfig["subAppId"]
	subMchId := appConfig["subMchId"]
	if subAppId != "" {
		mchClient = mchcore.NewSubMchClient(appId, mchId, apiKey, subAppId, subMchId, nil)
	} else {
		mchClient = mchcore.NewClient(appId, mchId, apiKey, nil)
	}
	srv.mchCore[appFinalId].mchClient = mchClient
	return mchClient, nil
}
func (srv *WeMch) GetMchTLSClient(appFinalId string) (*mchcore.Client, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appFinalId == "" {
		return nil, errors.New("empty appId")
	}
	if appCore, ok := srv.mchCore[appFinalId]; ok {
		if appCore.mchTLSClient != nil {
			return appCore.mchTLSClient, nil
		}
	} else {
		srv.mchCore[appFinalId] = new(MchCore)
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return nil, err
	}
	appId := appConfig["appId"]
	mchId := appConfig["mchId"]
	if appId == "" || mchId == "" {
		return nil, errors.New("appId or mchId is not set")
	}
	apiKey := appConfig["apiKey"]
	if apiKey == "" {
		return nil, errors.New("apiKey is not set")
	}
	certFile := appConfig["certFile"]
	keyFile := appConfig["keyFile"]
	if certFile == "" || keyFile == "" {
		return nil, errors.New("certFile or keyFile  is not set")
	}

	cli, err := mchcore.NewTLSHttpClient(certFile, keyFile)
	if err != nil {
		return nil, err
	}
	var mchTLSClient *mchcore.Client
	subAppId := appConfig["subAppId"]
	subMchId := appConfig["subMchId"]
	if subAppId != "" {
		mchTLSClient = mchcore.NewSubMchClient(appId, mchId, apiKey, subAppId, subMchId, cli)
	} else {
		mchTLSClient = mchcore.NewClient(appId, mchId, apiKey, cli)
	}
	srv.mchCore[appFinalId].mchTLSClient = mchTLSClient
	return mchTLSClient, nil
}
func (srv *WeMch) GetMchUnifiedOrderServer(appFinalId string, handler mchcore.Handler, errorHandler mchcore.ErrorHandler) (*mchcore.Server, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appFinalId == "" {
		return nil, errors.New("empty appId")
	}
	if appCore, ok := srv.mchCore[appFinalId]; ok {
		if appCore.mchUnifiedOrderServer != nil {
			return appCore.mchUnifiedOrderServer, nil
		}
	} else {
		srv.mchCore[appFinalId] = new(MchCore)
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return nil, err
	}
	appId := appConfig["appId"]
	mchId := appConfig["mchId"]
	if appId == "" || mchId == "" {
		return nil, errors.New("appId or mchId is not set")
	}
	apiKey := appConfig["apiKey"]
	if apiKey == "" {
		return nil, errors.New("apiKey is not set")
	}
	subAppId := appConfig["subAppId"]
	subMchId := appConfig["subMchId"]
	var mchUnifiedOrderServer *mchcore.Server
	if subAppId != "" {
		mchUnifiedOrderServer = mchcore.NewSubMchServer("", "", apiKey, subAppId, subMchId, handler, errorHandler)
	} else {
		mchUnifiedOrderServer = mchcore.NewServer(appId, mchId, apiKey, handler, errorHandler)
	}
	srv.mchCore[appFinalId].mchUnifiedOrderServer = mchUnifiedOrderServer
	return mchUnifiedOrderServer, nil
}
func (srv *WeMch) GetMchRefundServer(appFinalId string, handler mchcore.Handler, errorHandler mchcore.ErrorHandler) (*mchcore.Server, error) {
	srv.lock.RLock()
	defer srv.lock.RUnlock()
	if appFinalId == "" {
		return nil, errors.New("empty appId")
	}
	if appCore, ok := srv.mchCore[appFinalId]; ok {
		if appCore.mchUnifiedOrderServer != nil {
			return appCore.mchUnifiedOrderServer, nil
		}
	} else {
		srv.mchCore[appFinalId] = new(MchCore)
	}
	appConfig, err := srv.GetAppConfig(appFinalId)
	if err != nil {
		return nil, err
	}
	appId := appConfig["appId"]
	mchId := appConfig["mchId"]
	if appId == "" || mchId == "" {
		return nil, errors.New("appId or mchId is not set")
	}
	apiKey := appConfig["apiKey"]
	if apiKey == "" {
		return nil, errors.New("apiKey is not set")
	}
	subAppId := appConfig["subAppId"]
	subMchId := appConfig["subMchId"]
	var mchRefundServer *mchcore.Server
	if subAppId != "" {
		mchRefundServer = mchcore.NewSubMchServer("", "", apiKey, subAppId, subMchId, handler, errorHandler)
	} else {
		mchRefundServer = mchcore.NewServer(appId, mchId, apiKey, handler, errorHandler)
	}
	srv.mchCore[appFinalId].mchRefundServer = mchRefundServer
	return mchRefundServer, nil
}
