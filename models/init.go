package models

import (
	// "fmt"
	// "github.com/astaxie/beego/config"
	"errors"
	"github.com/dyzdyz010/Blog/ssdb"
)

var db *ssdb.Client

// Hash Map Names
var db_prefix = "blog_"
var h_entry = db_prefix + "entry"
var h_collection = db_prefix + "collection"
var h_author = db_prefix + "author"

func init() {
	err := errors.New("")
	db, err = ssdb.Connect("127.0.0.1", 8888)
	if err != nil {
		panic(err)
	}
}
