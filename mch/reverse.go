package mch

import (
	"errors"

	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// Reverse 撤销订单.
// NOTE: 请求需要双向证书.
func Reverse2(outTradeNo string, params ...string) (resp *mchpay.ReverseResponse, err error) {
	if MchConfig == nil {
		return nil, errors.New("not init MchConfig")
	}
	var transactionId string
	if len(params) >= 1 {
		transactionId = params[0]
	}
	req := &mchpay.ReverseRequest{
		TransactionId: transactionId,
		OutTradeNo:    outTradeNo,
	}
	resp, err = mchpay.Reverse2(MchTLSClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
