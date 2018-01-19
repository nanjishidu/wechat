package mch

import (
	"crypto/md5"
	"encoding/hex"
	"net"
	"strconv"
	"strings"

	uuid "github.com/satori/go.uuid"
)

//生成商品订单号
func GetOutTradeNo() string {
	u := uuid.Must(uuid.NewV4()).String()
	return strings.Replace(u, "-", "", -1)
}

//生成退款订单号
func GetOutRefundNo() string {
	u := uuid.Must(uuid.NewV1()).String()
	return strings.Replace(u, "-", "", -1)
}

//微信红包 商户订单号
func GetMchBillno() string {
	u := uuid.Must(uuid.NewV1()).String()
	return strings.Replace(u, "-", "", -1)
}

//企业付款 商户订单号
func GetPartnerRefundNo() string {
	u := uuid.Must(uuid.NewV1()).String()
	return strings.Replace(u, "-", "", -1)
}

func Md5(key string) string {
	h := md5.New()
	h.Write([]byte(key))
	return hex.EncodeToString(h.Sum(nil))
}
func GetInt64Str(d int64) string {
	return strconv.FormatInt(d, 10)
}
func GetIntStr(d int) string {
	return strconv.Itoa(d)
}

func GetLocalIp() string {
	addrs, err := net.InterfaceAddrs()
	if err == nil {
		for _, addr := range addrs {
			if ipnet, ok := addr.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
				if ipnet.IP.To4() != nil {
					return ipnet.IP.String()
				}
			}
		}
	}
	return "127.0.0.1"
}
