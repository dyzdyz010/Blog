package models

import (
	// "fmt"
	// "github.com/astaxie/beego/config"
	"errors"
	"github.com/dyzdyz010/Blog/ssdb"
)

var db *ssdb.Client

func init() {
	err := errors.New("")
	db, err = ssdb.Connect("127.0.0.1", 8888)
	if err != nil {
		panic(err)
	}
}
