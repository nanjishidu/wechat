package mch

import (
	"errors"

	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// CloseOrder2 关闭订单.
func CloseOrder2(outTradeNo string) (err error) {
	if MchConfig == nil {
		return errors.New("not init MchConfig")
	}
	req := &mchpay.CloseOrderRequest{
		OutTradeNo: outTradeNo,
	}
	err = mchpay.CloseOrder2(MchClient, req)
	if err != nil {
		return err
	}
	return nil
}
