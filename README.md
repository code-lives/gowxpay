# gowxpay
微信小程序支付（进行中）
# 安装
```
go get github.com/code-lives/gowxpay
```

# 自定义 env.ini 配置文件 例如如（env/pay.ini）
```
[WEIXIN]
Appid =
Secret =
MchId =
MchKey =
NotifyUrl = 回调地址
TradeType = JSAPI
```
# 全局调用  自定义(pay/pay.go)
```
package pay

import (
 "github.com/code-lives/gowxpay/wx"
)

var Config *wx.Config

func PayInit() {
 Config = wx.Init("weixin", "env/pay.ini")

}
func Pay() *wx.Config {
 return Config
}
```
# main.go 全局加载
```
pay.PayInit()
```
# 支付 记得(import pay/pay.go文件)
```

 data, err := pay.Pay().GetOrderParam("订单号", 10, "描述", "openid")
 if err != nil {
  panic(err)
 }
 fmt.Println(data)
```
# Openid 记得(import pay/pay.go文件)
```

 data, err := pay.Pay().GetOpenid("code")
 if err != nil {
  panic(err)
 }
 fmt.Println(data)
```

# 订单查询 (import pay/pay.go文件)
```
data, err := pay.FindOrder("订单号")
if err != nil {
    panic(err)
}
fmt.Println("data", data.IsSubscribe)
```

