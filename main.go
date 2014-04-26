package main

import (
	"github.com/astaxie/beego"
	_ "github.com/dyzdyz010/Blog/routers"
	"strings"
)

func StrToUpper(in string) (out string) {
	out = strings.ToUpper(in)
	return
}

func main() {
	beego.AddFuncMap("upper", StrToUpper)

	beego.Run()
}
