package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

//  Refund 申请退款.
//  NOTE: 请求需要双向证书.
func Refund2(totalFee, refundFee int64, outTradeNo, outRefundNo string, params ...string) (resp *pay.RefundResponse,
	err error) {
	if MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	var transactionId string
	if len(params) >= 1 {
		transactionId = params[0]
	}
	req := &pay.RefundRequest{
		TransactionId: transactionId,
		OutTradeNo:    outTradeNo,
		OutRefundNo:   outRefundNo,
		TotalFee:      totalFee,
		RefundFee:     refundFee,
	}
	resp, err = pay.Refund2(MchTLSClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
