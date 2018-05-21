package mmpaymkttransfers

import (
	"errors"

	mchmmpaymkttransfers "gopkg.in/chanxuehong/wechat.v2/mch/mmpaymkttransfers"
	wechatutil "gopkg.in/chanxuehong/wechat.v2/util"
	"gopkg.in/nanjishidu/wechat.v1/mch"
)

// 发送裂变红包红包
// 请求需要双向证书
// mch_billno 商户订单号
// send_name 商品名称 红包发送者名称
// re_openid 用户openid
// wishing 红包祝福语
// act_name 活动名称
// remark 备注
// total_amount 金额
// total_num 红包发放总人数
// scene_id 场景id
func SendGroupRedPack(mch_billno, send_name, re_openid, wishing, act_name, remark string,
	total_amount, total_num int64, scene_id ...int) (resp map[string]string, err error) {
	if mch.MchConfig == nil {
		return nil, errors.New("not init MchConfig")
	}
	if send_name == "" || re_openid == "" || wishing == "" || act_name == "" ||
		remark == "" || total_amount <= 0 {
		return nil, errors.New("parameter is incorrect")
	}
	if total_amount > 200*100 && len(scene_id) == 0 {
		return nil, errors.New("scene_id is null")
	}
	if total_num == 0 {
		total_num = 1
	}
	var scene_id_str string
	if len(scene_id) > 0 && scene_id[0] > 0 && scene_id[0] < 9 {
		scene_id_str = "PRODUCT_" + mch.GetIntStr(scene_id[0])
	}
	var req = map[string]string{
		"nonce_str":    wechatutil.NonceStr(),
		"mch_billno":   mch_billno,
		"client_ip":    mch.GetLocalIp(),
		"send_name":    send_name,
		"re_openid":    re_openid,
		"total_amount": mch.GetInt64Str(total_amount),
		"total_num":    mch.GetInt64Str(total_num),
		"wishing":      wishing,
		"act_name":     act_name,
		"remark":       remark,
		"amt_type":     "ALL_RAND",
		"scene_id":     scene_id_str,
	}
	if mch.MchConfig.SubAppId != "" {
		req["wxappid"] = mch.MchConfig.SubAppId
		req["msgappid"] = mch.MchConfig.AppId
		req["mch_id"] = mch.MchConfig.MchId
		req["sub_mch_id"] = mch.MchConfig.subMchId
	} else {
		req["mch_id"] = mch.MchConfig.MchId
		req["wxappid"] = mch.MchConfig.AppId
	}
	return mchmmpaymkttransfers.SendGroupRedPack(mch.MchTLSClient, req)
}
