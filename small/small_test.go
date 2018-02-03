// small_test.go
package small

import (
	// "encoding/base64"
	"fmt"
	"testing"
	"time"
)

var (
	encryptedData  = "6LNzLzbVnIpvMiNvpVrERKtCxiRtIG7ev4BNh1sFHQ5yC78RUkmNSwBPT0hvrMMUVsovhI6klS1FqRQn9w2qMP/jT4/Jx0DYTTqVLxgP/Rs5vDt9ceblI36m6CppaofcZzaj7uttwRTIbgIfRCZuaXT3O7OuT0jMCWVgnwR6XTb4eQIExLVOfiGOUPbSkeGlbcHJVGuK3UF2mdi0C50GQyTP2Iwb9l8BTkeY+wV4L67Hc5NUEgrN8lp0AZQKYOOduwFAh0e64vR4M4IxZU6hQRAnt6GM04TffLixPMYgWDD9D0bq/qPjXmdUy58bfFp4yYdPF4UxlaGT5Luf7Q6cNIEoE936ReHthhEk6SsvbDScgAmDPx2hVxZ8trj1TsfYF8lPpIdkkh4zYD5eiFvsc1A9r0liQUK8A/fb/xDipKbhNg513QnJ4aApPPxzpYe+UPXyXWIT8+wzlfzFnu20rX8WB4XwVa8TBU8SVTM4HiY="
	iv             = "YqO15JMdn/PTRRflnwT/7A=="
	sessionKey     = "3a6dWz/lMsi+eEw8LgBn5Q=="
	signature      = "766145434ad810e9bea254beb0daf13a0dc8ef89"
	rawData        = `{"nickName":"nanjishidu","gender":1,"language":"zh_CN","city":"Jinan","province":"Shandong","country":"China","avatarUrl":"https://wx.qlogo.cn/mmopen/vi_32/Q0j4TwGTfTK7jgM4moDqiaAx2JeGUSFPx59w78dS4eA3vbKc7vYicfeAzxEHKibnclhTy9uX8IhTx463VrRAnib5Ig/0"}`
	encryptedData2 = ""
	iv2            = ""
	sessionKey2    = ""
	encryptedData3 = "RQI1jOX+E+fetggVFpYVk0wHKYel6wcUM57B7Djis9BqDvp4wcjD2ndtnLeGFR4K3TUFUwvP8B1cRSn3AcfUG9Yig+NSDzalzkk9UY22rW8BWXEVnSa4zV19OiKKs6gR/S+TQXLuex3kQiS7t8JB6PEfPc0DTuCJPlOOuwurxsIDSILCp53anSO/sPOnnFmhPGzXREfCpjPDZ6mbzzCDbfyNNR1ivWuB37l5xzbs8bjKWGXOCLCI/Ho0kGgLWDXEsa6pi7NU2N5if16ZDpENkaNrISjBth0tAjVekyvaIMggKNQvBHFzbc/6rhYZdJSbCjVwAjOTQLFyJ0SZqwSuC2Qo8Lu4AXYwPIm8dS/d2II+yIpXtgQTpIVufhSO7iLxKCtY/Yjeu0JPkvaMMnRdQWjvHXZWr8VYywepNhzzvX1azlxnOObSHuc57aQ16LLdZ1YrWt/5n34yKfkxMHr/Z0IsswMEeiPxG1sFbRf6znjkTI+ObSZDYa+ADqNCtgiibpxnGU3TLqkrdZQzJHzqgXTPAyLcZB1NfEYeqav6QwhmBKLUwwNNGqO4sxSA5PyxCsJeMOHc2FWQ207pIirpubwMI+ADDqgqyE5gMIdYN44ackXamNQj2mB18MnWBKv8lD+GYQVGQWMh47dhrzTRUjEgJpzF2cL3jYajzkuT0RvGhzmIjy5hWP21DYtj2CiAsW2vWhYV/24TtMFEDr2MAezBx/yCF0JQKW+0T+phFneGXinVqVWdx5/x4Vrsh4FQZnMOAKmryZGp46z1OrDqfOEV2osBQuSRbe7XOljuht34/v2GvtyBlKZLuJYOr42T7gsGcD29MWJzTidx6v9I2urTAaN7Lc0nsOHcBsg73mzr0bm49LG+Ew5iFkDQhg+Zs4ZZdQSqPaVwGm+fUZxoB8HV/5Lnd3lUqmbtQT6SANIl9s4tHbw5tLfLzFolBC11H5RLGXAVO4yf35j6JkSsUxtAb5YybCF/oVaIrNAsXRZC2BillCu3vdEla9Aaf9X+NrvcTCYg+1TMGJkSXP1FPtM21YEMMMp9MQ94SkPYU4PKP7CAhe7YglGG31hJnoqBgbN571ae5Bj16jnNqFRhH9N2wVUA/D/E/2zx3OcgMp80t0pqAsGgG5w9pRdoqvgfMjKLfrdewH4HE97hH6Z+wpdzjFvMxNPehtoST9VjM8thI4oKbor/zQL9WBC2LUBENINrj/hvbteBnJvn2zxnfPrDg/yum9TObpVdJWFTqRhsZL1fdiGic+OXcTQV+lbvy9pijjr+bfsUNG7Z6Z6/tTEUmQMVNsuSmQBtnBck+/xuPL28rwQNy0whvHz9liwxV16+f56JYV0afZP8jC2YOasvAWVLmPR5/INcvN29VxJ0jNQiB7rT3o0PN/Ot6cXRg2uDeV6XFOyluPfxmOyOnlDJl7wUEmdeWo8ZpRgcyT5GnZhRg38SRtx/WjbB4pz5j7KTHIUGf5MfNUZJKqmilpsdYN7xagpgJBDWLaVediW6Vj0yOd/SqYMjZQ24nTuusfEA6FtuBZ4RwbrSp7Lt5Eyo5/NDVx6hLIGYRLVi3z1sm4Llqa1hekfu1Y1Am+i2fDo7uE03FzHT7rDfFxrSM7T2G9tdWSblkpRqVa2kH9g3d3cgnipI69hE7/7NhobJ"
	iv3            = "mEArGQh2DoQEhBZtt4S+aA=="
	sessionKey3    = "pZddSukZ8UPYBFcmzGuyBA=="
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
func TestGetPhoneNumber(t *testing.T) {
	wpn, err := GetPhoneNumber(sessionKey2, encryptedData2, iv2)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	fmt.Println(wpn)
	fmt.Println("用户绑定的手机号（国外手机号会有区号）", wpn.PhoneNumber)
	fmt.Println("没有区号的手机号", wpn.PurePhoneNumber)
	fmt.Println("区号", wpn.CountryCode)
	fmt.Println("appid", wpn.Watermark.Appid)
	fmt.Println("timestamp", wpn.Watermark.Timestamp)
}
func TestGetWeRunData(t *testing.T) {
	wrd, err := GetWeRunData(sessionKey3, encryptedData3, iv3)
	if err != nil {
		t.Fatal(err)
		t.FailNow()
	}
	for _, v := range wrd.StepInfoList {
		fmt.Println("时间", time.Unix(v.Timestamp, 0).Format("2006-01-02 15:04:05"))
		fmt.Println("步数", v.Step)

	}
}
