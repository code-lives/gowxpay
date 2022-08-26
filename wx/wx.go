package wx

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"net"
	"sort"
	"strings"
	"time"

	"github.com/code-lives/gowxpay/autoloading"
	"github.com/code-lives/gowxpay/common"
)

var (
	config     = &Config{}
	OrderParam interface{}
	Order      = make(map[string]interface{})
	err        error
)

func Init(s string, p string) *Config {

	autoloading.GetEnv(s, p, &config)
	if config.Appid == "" || config.Secret == "" {
		panic("Appid Is Null Or Appid Is Null")
	}
	return config
}

// GetOrderParam 设置订单
func (w *Config) GetOrderParam(o string, m int32, d string, openid string) (interface{}, error) {
	Order["appid"] = w.Appid
	Order["mch_id"] = w.MchId
	Order["trade_type"] = w.TradeType
	Order["nonce_str"] = nonceStr()
	Order["body"] = d
	Order["out_trade_no"] = o
	Order["total_fee"] = m
	Order["notify_url"] = w.NotifyUrl
	if w.TradeType == "JSAPI" {
		Order["openid"] = openid
	}
	Order["spbill_create_ip"] = GetRemoteClientIp()
	Order["sign"] = sign(Order)
	data := MapXml(Order)
	xmls := common.HttpPost(PayUrl, "POST", "application/xml", data)
	x := ReturnData{}
	if err = xml.Unmarshal([]byte(xmls), &x); err != nil {
		return nil, fmt.Errorf("err %s", err.Error())
	}
	u, err := json.Marshal(x)
	if err != nil {
		return nil, fmt.Errorf("err %s", err.Error())
	}
	if w.TradeType == "MWEB" {
		return string(u), nil
	}
	var Xcx = make(map[string]interface{})
	Xcx["appId"] = w.Appid
	Xcx["timeStamp"] = fmt.Sprintf("%d", time.Now().Unix())
	Xcx["nonceStr"] = nonceStr()
	Xcx["package"] = "prepay_id=" + x.PrepayId
	Xcx["signType"] = "MD5"
	Xcx["paySign"] = sign(Xcx)
	pay, err := json.Marshal(Xcx)
	if err != nil {
		return nil, fmt.Errorf("err %s", err.Error())
	}
	return string(pay), nil
}
func (w *Config) GetOpenid(c string) *Openid {
	openidString := common.HttpGet(CodedUrl + "appid=" + w.Appid + "&secret=" + w.Secret + "&js_code=" + c + "&grant_type=authorization_code")
	openid := &Openid{}
	err = json.Unmarshal([]byte(openidString), openid)
	if err != nil {
		panic(err)
	}
	return openid
}
func GetRemoteClientIp() string {
	adders, err := net.InterfaceAddrs()
	if err != nil {
		return "127.0.0.1"
	}
	for _, address := range adders {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return "127.0.0.1"
}
func nonceStr() string {
	chars := "abcdefghijklmnopqrstuvwxyz0123456789"
	bytes := []byte(chars)
	var result []byte
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 32; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//签名
func sign(o map[string]interface{}) string {
	joinString := formatBizQueryParaMap(o)
	return strings.ToUpper(common.Setmd5(joinString + "&key=" + config.MchKey))
}

// MapXml Map转换为XML
func MapXml(o map[string]interface{}) string {
	keys := make([]string, 0, len(o))
	for k := range o {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	xml := "<xml>"
	for _, k := range keys {
		xml += "<" + fmt.Sprint(k) + ">" + fmt.Sprint(o[k]) + "</" + fmt.Sprint(k) + ">"
	}
	xml += "</xml>"
	return xml
}

// 转换成字符串
func formatBizQueryParaMap(o map[string]interface{}) string {
	keys := make([]string, 0, len(o))
	for k := range o {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	buff := ""
	for _, k := range keys {
		buff += fmt.Sprint(k) + "=" + fmt.Sprint(o[k]) + "&"
	}
	reqpar := ""
	if len(buff) > 0 {
		reqpar = buff[:len(buff)-1]
	}
	return reqpar
}
