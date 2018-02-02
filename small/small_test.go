// small_test.go
package small

import (
	// "encoding/base64"
	"fmt"
	"testing"
)

var (
	encryptedData = "6LNzLzbVnIpvMiNvpVrERKtCxiRtIG7ev4BNh1sFHQ5yC78RUkmNSwBPT0hvrMMUVsovhI6klS1FqRQn9w2qMP/jT4/Jx0DYTTqVLxgP/Rs5vDt9ceblI36m6CppaofcZzaj7uttwRTIbgIfRCZuaXT3O7OuT0jMCWVgnwR6XTb4eQIExLVOfiGOUPbSkeGlbcHJVGuK3UF2mdi0C50GQyTP2Iwb9l8BTkeY+wV4L67Hc5NUEgrN8lp0AZQKYOOduwFAh0e64vR4M4IxZU6hQRAnt6GM04TffLixPMYgWDD9D0bq/qPjXmdUy58bfFp4yYdPF4UxlaGT5Luf7Q6cNIEoE936ReHthhEk6SsvbDScgAmDPx2hVxZ8trj1TsfYF8lPpIdkkh4zYD5eiFvsc1A9r0liQUK8A/fb/xDipKbhNg513QnJ4aApPPxzpYe+UPXyXWIT8+wzlfzFnu20rX8WB4XwVa8TBU8SVTM4HiY="
	iv            = "YqO15JMdn/PTRRflnwT/7A=="
	sessionKey    = "3a6dWz/lMsi+eEw8LgBn5Q=="
	signature     = "766145434ad810e9bea254beb0daf13a0dc8ef89"
	rawData       = `{"nickName":"nanjishidu","gender":1,"language":"zh_CN","city":"Jinan","province":"Shandong","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTK7jgM4moDqiaAx2JeGUSFPx59w78dS4eA3vbKc7vYicfeAzxEHKibnclhTy9uX8IhTx463VrRAnib5Ig/0"}`
)

func TestGetWxSessionKey(t *testing.T) {
	// ws, err := GetWxSessionKey(appid, secret, code)
	// if err != nil {
	// 	t.Fatal(err)
	// }
	// fmt.Println(ws)
	// if ws.ErrCode != 0 {
	// 	t.Log(ws.ErrMsg)
	// 	t.FailNow()
	// }
	// fmt.Println(ws.Openid)
	// fmt.Println(ws.SessionKey)
	if CheckSignature(signature, sessionKey, rawData) != true {
		t.Fatal("CheckSignature failed")
		t.FailNow()
	}
	u, err := GetWxUserInfo(sessionKey, encryptedData, iv)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	fmt.Println(u)

}
