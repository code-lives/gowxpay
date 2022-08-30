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
	MchId     string
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
type FindData struct {
	ReturnCode     string `xml:"return_code"`
	ReturnMsg      string `xml:"return_msg"`
	ResultCode     string `xml:"result_code"`
	Appid          string `xml:"appid"`
	NonceStr       string `xml:"nonce_str"`
	Sign           string `xml:"sign"`
	TradeType      string `xml:"trade_type"`
	MchId          string `xml:"mch_id"`
	Openid         string `xml:"openid"`
	IsSubscribe    string `xml:"is_subscribe"`
	TradeState     string `xml:"trade_state"`
	BankType       string `xml:"bank_type"`
	TotalFee       string `xml:"total_fee"`
	FeeType        string `xml:"fee_type"`
	CashFee        string `xml:"cash_fee"`
	CashFeeType    string `xml:"cash_fee_type"`
	TransactionId  string `xml:"transaction_id"`
	OutTradeNo     string `xml:"out_trade_no"`
	Attach         string `xml:"attach"`
	TimeEnd        string `xml:"time_end"`
	TradeStateDesc string `xml:"trade_state_desc"`
}
type PayOrder struct {
	AppId     string `json:"appId"`
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"package"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
type Openid struct {
	SessionKey string `json:"session_key"`
	ExpiresIn  int64  `json:"expires_in"`
	Openid     string `json:"openid"`
	Unionid    string `json:"unionid"`
}
