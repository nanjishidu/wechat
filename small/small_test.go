// small_test.go
package small

import (
	// "encoding/base64"
	"fmt"
	"testing"
)

var (
	encryptedData = ""
	iv            = ""
	sessionKey    = ""
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
	// if CheckSignature(signature, sessionKey, rawData) != true {
	// 	t.Fatal("CheckSignature failed")
	// 	t.FailNow()
	// }
	u, err := GetWxUserInfo(sessionKey, encryptedData, iv)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	fmt.Println(u)

}
