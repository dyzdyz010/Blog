package main

import (
	_ "Blog/routers"
	"github.com/astaxie/beego"
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
