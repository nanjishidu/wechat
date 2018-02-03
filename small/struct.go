package small

//微信加密数据结构
type WxUserInfo struct {
	OpenId    string     `json:"openId"`
	NickName  string     `json:"nickName"`
	Gender    int        `json:"gender"`
	City      string     `json:"city"`
	Province  string     `json:"province"`
	Country   string     `json:"country"`
	AvatarUrl string     `json:"avatarUrl"`
	UnionId   string     `json:"unionId"`
	Watermark *Watermark `json:"watermark"` //数据水印( watermark )
}
type Watermark struct {
	Appid     string `json:"appid"`
	Timestamp int64  `json:"timestamp"`
}
type WxSesstion struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	ErrInfo
}
type ErrInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type WxPhoneNumber struct {
	PhoneNumber     string     `json:"phoneNumber"`     //用户绑定的手机号（国外手机号会有区号）
	PurePhoneNumber string     `json:"purePhoneNumber"` //没有区号的手机号
	CountryCode     string     `json:"countryCode"`     //区号
	Watermark       *Watermark `json:"watermark"`       //数据水印( watermark )
}
type WeRunData struct {
	StepInfoList []*StepInfo `json:"stepInfoList"`
}
type StepInfo struct {
	Step      int64 `json:"step"`
	Timestamp int64 `json:"timestamp"`
}
