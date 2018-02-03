package mch

import (
	"errors"

	"time"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

//统一下单接口 基于wechat.v2进行二次封装
/*
用户 前端  服务端  微信服务端
1. 前端调用网页生成支付订单
2. 服务端接收请求后，调用统一下单接口
3. 微信服务端接受请求后生成预付单信息返回服务端
4. 服务端接收信息后生成JSAPI页面跳用的支付参数并签名
5. 前端接受支付参数并发起JSAPI接口请求支付，用户发起支付后
6. 微信服务端进行参数合法性和授权域权限检查，返回结果给前端
7. 前端输入密码提交授权后，异步通知商户服务结果，通过微信服务端告知处理结果
9. 微信服务端返回支付结果到前端，前端展示支付消息给用户并跳转到信息反馈界面
*/

//简化 wechat.v2 中部分参数
/*
totalFee 订单总金额，单位为分，详见支付金额
openId   用户在商户appid下的唯一标识
outTradeNo 商户系统内部的订单号,32个字符内、可包含字母, 其他说明见商户订单号
body 商品或支付单简要描述
spbillCreateIP APP和网页支付提交用户端ip，Native支付填调用微信支付API的机器IP
notifyURL 接收微信支付异步通知回调地址，通知url必须为直接可访问的url，不能携带参数
detail 商品名称明细列表。
attach 附加数据，在查询API和支付通知中原样返回，该字段主要用于商户携带订单的自定义数据

UnifiedOrderResponse
PrepayId为微信生成预支付绘画标示，用于后续接口使用，有效期为两小时
TradeType为取值如下：JSAPI，NATIVE，APP，
JSAPI--公众号支付、NATIVE--原生扫码支付、APP--app支付，统一下单接口trade_type的传参可参考这里
MICROPAY--刷卡支付，刷卡支付有单独的支付接口，不调用统一下单接口
*/
//统一下单
func JsapiUnifiedOrder(totalFee int64, openId, outTradeNo, body, spbillCreateIP, notifyURL, detail, attach, goodsTag string) (resp *pay.UnifiedOrderResponse, err error) {
	req := &pay.UnifiedOrderRequest{
		Body:           body,
		OutTradeNo:     outTradeNo,
		TotalFee:       totalFee,
		SpbillCreateIP: spbillCreateIP,
		NotifyURL:      notifyURL,
		TradeType:      "JSAPI",
		DeviceInfo:     "web",
		Detail:         detail,
		Attach:         attach,
		GoodsTag:       goodsTag,
		FeeType:        "CNY",
		TimeStart:      time.Now(),
		TimeExpire:     time.Now().Add(600 * time.Second),
		OpenId:         openId,
	}
	return UnifiedOrder(req)
}
func UnifiedOrder(req *pay.UnifiedOrderRequest) (resp *pay.UnifiedOrderResponse, err error) {
	if MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}

	resp, err = pay.UnifiedOrder2(MchClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
