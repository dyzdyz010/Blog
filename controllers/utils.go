package controllers

import (
	"bytes"
	// "fmt"
	"github.com/astaxie/beego/context"
	"github.com/eknkc/amber"
)

func renderTemplate(ctx *context.Context, tplName string, data map[interface{}]interface{}) {
	compiler := amber.New()
	err := compiler.ParseFile(tplName)
	if err != nil {
		panic(err)
	}
	tpl, err := compiler.Compile()
	if err != nil {
		panic(err)
	}

	var content bytes.Buffer
	err = tpl.Execute(&content, data)
	if err != nil {
		panic(err)
	}
	ctx.WriteString(content.String())
}
