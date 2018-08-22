package mini

type SesstionInfo struct {
	Openid     string `json:"openid"`
	SessionKey string `json:"session_key"`
	Unionid    string `json:"unionid"`
	ErrInfo
}

//微信加密数据结构
type UserInfo struct {
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

type ErrInfo struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}
type PhoneNumber struct {
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

const (
	WxAcodeTypeA string = "A"
	WxAcodeTypeB string = "B"
	WxAcodeTypeC string = "C"
)

type WxAcode struct {
	Scene     string     `json:"scene,omitempty"`
	Path      string     `json:"path,omitempty"`
	Width     int        `json:"width,omitempty"`
	AutoColor bool       `json:"auto_color,omitempty"`
	LineColor *LineColor `json:"line_color,omitempty"`
	IsHyaline bool       `json:"is_hyaline,omitempty"`
}
type LineColor struct {
	R string `json:"r"`
	G string `json:"g"`
	B string `json:"b"`
}
