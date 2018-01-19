package mmpaymkttransfers

import (
	"errors"

	"github.com/nanjishidu/wechat/mch"
	"gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers"
	wechatutil "gopkg.in/chanxuehong/wechat.v2/util"
)

// 查询企业付款.
// 请求需要双向证书
// 商户号 mch_id
// Appid appid
// 签名   sign
// 以上参数调用接口时自动追加
// partner_trade_no  商户调用企业付款API时使用的商户订单号
func GetTransferInfo(partner_trade_no string) (resp map[string]string, err error) {
	if mch.MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	if partner_trade_no == "" {
		return nil, errors.New("parameter is incorrect")
	}
	var req = map[string]string{
		"nonce_str":        wechatutil.NonceStr(),
		"partner_trade_no": partner_trade_no,
	}
	return mmpaymkttransfers.GetTransferInfo(mch.MchTLSClient, req)
}
