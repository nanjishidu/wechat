package mmpaymkttransfers

import (
	"errors"

	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
	mchmmpaymkttransfers "gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers"
	wechatutil "gopkg.in/chanxuehong/wechat.v2/util"
)

//查询红包记录
// mch_billno  商户发放红包的商户订单号
func GetRedPackInfo(mchTLSClient *mchcore.Client, mch_billno string) (resp map[string]string, err error) {
	if mch_billno == "" {
		return nil, errors.New("parameter is incorrect")
	}
	var req = map[string]string{
		"nonce_str":  wechatutil.NonceStr(),
		"mch_billno": mch_billno,
		"bill_type":  "MCHT",
	}
	return mchmmpaymkttransfers.GetRedPackInfo(mchTLSClient, req)
}
