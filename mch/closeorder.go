package mch

import (
	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// CloseOrder2 关闭订单.
func CloseOrder2(mchClient *mchcore.Client, outTradeNo string) (err error) {
	req := &mchpay.CloseOrderRequest{
		OutTradeNo: outTradeNo,
	}
	err = mchpay.CloseOrder2(mchClient, req)
	if err != nil {
		return err
	}
	return nil
}
