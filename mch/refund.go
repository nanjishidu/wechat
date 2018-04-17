package mch

import (
	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

//  Refund 申请退款.
//  NOTE: 请求需要双向证书.
func Refund2(mchTLSClient *mchcore.Client, totalFee, refundFee int64, outTradeNo, outRefundNo string,
	params ...string) (resp *mchpay.RefundResponse, err error) {
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
	resp, err = mchpay.Refund2(mchTLSClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
