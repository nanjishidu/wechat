package mmpaymkttransfers

import (
	"errors"

	mchmmpaymkttransfers "gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers"
	wechatutil "gopkg.in/chanxuehong/wechat.v2/util"
	"gopkg.in/nanjishidu/wechat.v1/mch"
)

// 查询企业付款.
// 请求需要双向证书
// 商户号 mch_id
// Appid appid
// 签名   sign
// 以上参数调用接口时自动追加
// partner_trade_no  商户调用企业付款API时使用的商户订单号
func GetTransferInfo(partner_trade_no string) (resp map[string]string, err error) {
	if mch.MchConfig == nil {
		return nil, errors.New("not init MchConfig")
	}
	if partner_trade_no == "" {
		return nil, errors.New("parameter is incorrect")
	}
	var req = map[string]string{
		"nonce_str":        wechatutil.NonceStr(),
		"partner_trade_no": partner_trade_no,
	}
	return mchmmpaymkttransfers.GetTransferInfo(mch.MchTLSClient, req)
}
