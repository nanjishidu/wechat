package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// Reverse 撤销订单.
// NOTE: 请求需要双向证书.
func Reverse2(outTradeNo string, params ...string) (resp *pay.ReverseResponse, err error) {
	if MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	var transactionId string
	if len(params) >= 1 {
		transactionId = params[0]
	}
	req := &pay.ReverseRequest{
		TransactionId: transactionId,
		OutTradeNo:    outTradeNo,
	}
	resp, err = pay.Reverse2(MchTLSClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
