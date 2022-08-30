package autoloading

import (
	"strings"

	"gopkg.in/ini.v1"
)

var err error
var conf *ini.File

func GetEnv(s string, p string, config interface{}) {
	if conf, err = ini.Load(p); err != nil {
		panic("支付env配置文件读取失败" + err.Error())
	}
	if conf.Section(strings.ToUpper(s)).MapTo(config); err != nil {
		panic("结构体绑定失败" + err.Error())
	}
}
