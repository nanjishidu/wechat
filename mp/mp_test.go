package mp

import (
	"fmt"
	"testing"
	"time"

	mptemplate "gopkg.in/chanxuehong/wechat.v2/mp/message/template"
)

var (
	Srv *WeMp
)

func InitWeMp() {
	Srv = NewWeMp(map[string]map[string]string{
		"wxa98b6372189983ea": map[string]string{
			"appId":        "wxa98b6372189983ea",
			"appSecret":    "204b13f456ad051c94d9bbb84d1020c8",
			"token":        "",
			"base64AESKey": "",
		},
	})
}

func TestGetAccessTokenServer(t *testing.T) {
	InitWeMp()
	a, err := Srv.GetAccessTokenServer("wxa98b6372189983ea")
	if err != nil {
		fmt.Println(err)
	}
	aa, err := CreateTempQrcode(a, 1, 1)
	if err != nil {
		fmt.Println(aa, err)
	} else {
		fmt.Println("aa", aa)
	}
}
func TestSendTemplateNews(t *testing.T) {
	InitWeMp()
	a, err := Srv.GetAccessTokenServer("wxa98b6372189983ea")
	if err != nil {
		fmt.Println(err)
	} else {
		for {
			accessToken, _ := a.Token()
			fmt.Println("accessToken", accessToken)
			_, err = SendTemplateNews(a, "o3UAiw1vwLB81n480ZhpTmsesulo", "SMqZEWqUZ9CMFOd60g7gGE8QUjXcAJSdReNUqITt9LA",
				"", map[string]mptemplate.DataItem{
					"first":    mptemplate.DataItem{Value: "hello world 2018"},
					"keyword1": mptemplate.DataItem{Value: "1000元"},
					"keyword2": mptemplate.DataItem{Value: time.Now().Format("2006-01-02 15:04:05")},
					"remark":   mptemplate.DataItem{Value: ""},
				})
			if err != nil {
				fmt.Println("err", err)
			}
			time.Sleep(1000 * time.Second)
		}

	}

}
