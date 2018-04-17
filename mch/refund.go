package mch

import (
	"errors"

	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

//  Refund 申请退款.
//  NOTE: 请求需要双向证书.
func Refund2(totalFee, refundFee int64, outTradeNo, outRefundNo string, params ...string) (resp *mchpay.RefundResponse,
	err error) {
	if MchConfig == nil {
		return nil, errors.New("not init MchConfig")
	}
	var transactionId string
	if len(params) >= 1 {
		transactionId = params[0]
	}
	req := &mchpay.RefundRequest{
		TransactionId: transactionId,
		OutTradeNo:    outTradeNo,
		OutRefundNo:   outRefundNo,
		TotalFee:      totalFee,
		RefundFee:     refundFee,
	}
	resp, err = mchpay.Refund2(MchTLSClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
