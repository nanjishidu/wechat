package mch

import (
	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// Reverse 撤销订单.
// NOTE: 请求需要双向证书.
func Reverse2(mchTLSClient *mchcore.Client, outTradeNo string, params ...string) (resp *mchpay.ReverseResponse, err error) {
	var transactionId string
	if len(params) >= 1 {
		transactionId = params[0]
	}
	req := &mchpay.ReverseRequest{
		TransactionId: transactionId,
		OutTradeNo:    outTradeNo,
	}
	resp, err = mchpay.Reverse2(mchTLSClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
