package main

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/models"
	_ "github.com/dyzdyz010/Blog/routers"
	. "github.com/qiniu/api/conf"
	"strings"
)

func StrToUpper(in string) (out string) {
	out = strings.ToUpper(in)
	return
}

func main() {
	beego.AddFuncMap("upper", StrToUpper)

	// Qiniu
	ACCESS_KEY = models.Appconf.String("qiniu::ak")
	SECRET_KEY = models.Appconf.String("qiniu::sk")

	beego.Run()
}
