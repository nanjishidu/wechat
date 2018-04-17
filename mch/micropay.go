package mch

import (
	mchcore "gopkg.in/chanxuehong/wechat.v2/mch/core"
	mchpay "gopkg.in/chanxuehong/wechat.v2/mch/pay"
)

func MicroPay2(mchClient *mchcore.Client, totalFee int64, outTradeNo, body, spbillCreateIP,
	authCode string) (resp *mchpay.MicroPayResponse, err error) {
	resp, err = mchpay.MicroPay2(mchClient, &mchpay.MicroPayRequest{
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
