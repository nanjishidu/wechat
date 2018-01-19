package mch

import (
	"errors"

	"gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

func MicroPay2(totalFee int64, outTradeNo, body, spbillCreateIP, authCode string) (resp *pay.MicroPayResponse, err error) {
	if MchCommonConfig == nil {
		return nil, errors.New("not init MchCommonConfig")
	}
	resp, err = pay.MicroPay2(MchClient, &pay.MicroPayRequest{
		Body:           body,
		OutTradeNo:     outTradeNo,
		TotalFee:       totalFee,
		SpbillCreateIP: spbillCreateIP,
		AuthCode:       authCode,
	})
	if err != nil {
		return nil, err
	}
	return resp, nil
}
