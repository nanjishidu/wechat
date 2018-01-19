package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// RefundQuery 查询退款.
// out_trade_no   // 商户订单号
// transaction_id // 微信订单号
// out_refund_no  // 商户退款单号
// refund_id      // 微信退款单号
func RefundQuery2(outTradeNo string, params ...string) (resp *pay.RefundQueryResponse,
	err error) {
	if MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	var (
		transactionId, outRefundNo, refundId string
	)
	switch len(params) {
	case 1:
		transactionId = params[0]
	case 2:
		transactionId = params[0]
		outRefundNo = params[1]
	case 3:
		transactionId = params[0]
		outRefundNo = params[1]
		refundId = params[2]
	}
	req := &pay.RefundQueryRequest{
		OutTradeNo:    outTradeNo,
		TransactionId: transactionId,
		OutRefundNo:   outRefundNo,
		RefundId:      refundId,
	}
	resp, err = pay.RefundQuery2(MchClient, req)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
