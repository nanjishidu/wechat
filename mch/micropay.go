package mch

import (
	"errors"

	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

func MicroPay2(totalFee int64, outTradeNo, body, spbillCreateIP, authCode string) (resp *mchpay.MicroPayResponse, err error) {
	if MchConfig == nil {
		return nil, errors.New("not init MchConfig")
	}
	resp, err = mchpay.MicroPay2(MchClient, &mchpay.MicroPayRequest{
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
