package promotion

import (
	"errors"

	"github.com/nanjishidu/wechat/mch"
	"gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers/promotion"
	wechatutil "gopkg.in/chanxuehong/wechat.v2/util"
)

// 企业付款.
// 请求需要双向证书
// 商户号 mch_id
// Appid appid
// 签名   sign
// 以上参数调用接口时自动追加
//partner_trade_no 		商户订单号
//openid 				商户appid下，某用户的openid
//re_user_name 			收款用户真实姓名,可以为空
//desc	 				企业付款操作说明信息,必填。
//amount 				企业付款金额，单位为分
func Transfers(partner_trade_no, openid, re_user_name, desc string,
	amount int64) (resp map[string]string, err error) {
	if mch.MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	if openid == "" || desc == "" || amount <= 0 {
		return nil, errors.New("parameter is incorrect")
	}
	var check_name string = "NO_CHECK"
	if re_user_name != "" {
		check_name = "FORCE_CHECK"
	}
	var req = map[string]string{
		"mch_appid":        mch.MchCommonConfig.AppId,
		"mchid":            mch.MchCommonConfig.MchId,
		"nonce_str":        wechatutil.NonceStr(),
		"partner_trade_no": partner_trade_no,
		"openid":           openid,
		"check_name":       check_name,
		"re_user_name":     check_name,
		"desc":             desc,
		"amount":           mch.GetInt64Str(amount),
		"spbill_create_ip": mch.GetLocalIp(),
	}
	return promotion.Transfers(mch.MchTLSClient, req)
}
