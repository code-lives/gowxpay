package wx

const (
	CodedUrl   = "https://api.weixin.qq.com/sns/jscode2session?"
	AppCodeUrl = "https://api.weixin.qq.com/sns/oauth2/access_token?"
	TokenUrl   = "https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid="
	PayUrl     = "https://api.mch.weixin.qq.com/pay/unifiedorder"
	Query      = "https://api.mch.weixin.qq.com/pay/orderquery"
	RefundUrl  = "https://api.mch.weixin.qq.com/secapi/pay/refund"
)

type Config struct {
	Appid     string
	Secret    string
	MchId     int64
	MchKey    string
	NotifyUrl string
	TradeType string
	KeyPem    string
	CertPem   string
}

type ReturnData struct {
	ReturnCode string `xml:"return_code"`
	ReturnMsg  string `xml:"return_msg"`
	ResultCode string `xml:"result_code"`
	Appid      string `xml:"appid"`
	NonceStr   string `xml:"nonce_str"`
	Sign       string `xml:"sign"`
	PrepayId   string `xml:"prepay_id"`
	TradeType  string `xml:"trade_type"`
}
type Openid struct {
	SessionKey string `json:"session_key"`
	ExpiresIn  int64  `json:"expires_in"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
}
