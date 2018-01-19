package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// TradeState     交易状态
// TradeStateDesc 对当前查询订单状态的描述和下一步操作的指引
// OpenId          用户在商户appid下的唯一标识
// TransactionId   微信支付订单号
// OutTradeNo      商户系统的订单号，与请求一致。
// TradeType       调用接口提交的交易类型，取值如下：JSAPI，NATIVE，APP，MICROPAY，详细说明见参数规定
// BankType        银行类型，采用字符串类型的银行标识
// TotalFee        订单总金额，单位为分
// CashFee         现金支付金额订单现金支付金额，详见支付金额
// TimeEnd         订单支付时间，格式为yyyyMMddHHmmss，如2009年12月25日9点10分10秒表示为20091225091010。其他详见时间规则

//查询订单
func OrderQuery2(outTradeNo string, params ...string) (resp *pay.OrderQueryResponse, err error) {
	if MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	var transactionId string
	if len(params) > 0 {
		transactionId = params[0]
	}
	req := &pay.OrderQueryRequest{
		TransactionId: transactionId,
		OutTradeNo:    outTradeNo,
	}
	resp, err = pay.OrderQuery2(MchClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
