package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

// CloseOrder2 关闭订单.
func CloseOrder2(outTradeNo string) (err error) {
	if MchCommonConfig == nil {
		return errors.New("not init MchCommonConfig")
	}
	req := &pay.CloseOrderRequest{
		OutTradeNo: outTradeNo,
	}
	err = pay.CloseOrder2(MchClient, req)
	if err != nil {
		return err
	}
	return nil
}
