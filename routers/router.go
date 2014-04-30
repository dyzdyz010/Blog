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

	// Admin -> Login
	beego.Router("/admin/login", &controllers.AdminController{}, "get:Login")

	// Admin -> Entry
	beego.Router("/admin/entries", &controllers.AdminController{}, "get:Entries")
	beego.Router("/admin/entries/:id", &controllers.AdminController{}, "get:Entry")
	beego.Router("/admin/entries/:id/update", &controllers.AdminController{}, "post:UpdateEntry")
	beego.Router("/admin/entries/new", &controllers.AdminController{}, "get:NewEntry;post:PostNewEntry")

	// Admin -> Collection
	beego.Router("/admin/collections", &controllers.AdminController{}, "get:Collections")
	beego.Router("/admin/collections/:id", &controllers.AdminController{}, "get:Collection")
	beego.Router("/admin/collections/:id/update", &controllers.AdminController{}, "post:UpdateCollection")
	beego.Router("/admin/collections/new", &controllers.AdminController{}, "get:NewCollection;post:CreateCollection")
}
