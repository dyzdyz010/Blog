package routers

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/controllers"
)

func init() {
	beego.Router("/", &controllers.FrontController{}, "get:Home")
	beego.Router("/collections", &controllers.FrontController{}, "get:Collections")
	beego.Router("/entry/:id", &controllers.FrontController{}, "get:Entry")

	// Admin
	beego.Router("/admin", &controllers.AdminController{}, "get:Dashboard")
	beego.Router("/admin/entries", &controllers.AdminController{}, "get:Entries")
	beego.Router("/admin/entries/:id", &controllers.AdminController{}, "get:Entry")
	beego.Router("/admin/entries/:id", &controllers.AdminController{}, "post:PostEntry")
}
