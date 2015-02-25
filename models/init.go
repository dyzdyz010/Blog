package models

import (
	"errors"
	"fmt"
	"github.com/astaxie/beego/config"
	"github.com/dyzdyz010/Blog/ssdb"
)

var db *ssdb.Client
var Appconf config.ConfigContainer

// Hash Map Names
var db_prefix = "blog_"
var h_entry = db_prefix + "entry"
var h_collection = db_prefix + "collection"
var h_author = db_prefix + "author"

var page_size = 10

func init() {
	err := errors.New("")

	Appconf, err = config.NewConfig("json", "conf/blog.json")
	if err != nil {
		panic(err)
	}

	fmt.Println(Appconf.Int("database::port"))
	host := Appconf.String("database::host")
	port, _ := Appconf.Int("database::port")

	db, err = ssdb.Connect(host, port)
	if err != nil {
		panic(err)
	}
}
