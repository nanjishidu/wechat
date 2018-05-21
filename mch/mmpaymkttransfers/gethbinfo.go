package mmpaymkttransfers

import (
	"errors"

	mchmmpaymkttransfers "gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers"
	wechatutil "gopkg.in/chanxuehong/wechat.v2/util"
	"gopkg.in/nanjishidu/wechat.v1/mch"
)

//查询红包记录
// mch_billno  商户发放红包的商户订单号
func GetRedPackInfo(mch_billno string) (resp map[string]string, err error) {
	if mch.MchConfig == nil {
		return nil, errors.New("not init MchConfig")
	}
	if mch_billno == "" {
		return nil, errors.New("parameter is incorrect")
	}
	var req = map[string]string{
		"nonce_str":  wechatutil.NonceStr(),
		"mch_billno": mch_billno,
		"bill_type":  "MCHT",
	}
	return mchmmpaymkttransfers.GetRedPackInfo(mch.MchTLSClient, req)
}
