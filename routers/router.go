package routers

import (
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/controllers"
)

func init() {

	//-------------------------------------------------------------------------------
	//                Front
	//-------------------------------------------------------------------------------

	// Front -> Home
	beego.Router("/", &controllers.FrontController{}, "get:Home")

	// Front -> Entry
	beego.Router("/entry/:id", &controllers.FrontController{}, "get:Entry")

	// Front -> Collection List
	beego.Router("/collections", &controllers.FrontController{}, "get:Collections")

	// Front -> Collection's Entry List
	beego.Router("/collection/:id", &controllers.FrontController{}, "get:Collection")

	// Front ->

	//-------------------------------------------------------------------------------
	//                Admin
	//-------------------------------------------------------------------------------

	// Admin -> Dashboard
	beego.Router("/admin", &controllers.AdminController{}, "get:Dashboard")

	// Admin -> Login
	beego.Router("/admin/login", &controllers.AdminController{}, "get:Login;post:PostLogin")

	// Admin -> Entry
	beego.Router("/admin/entries", &controllers.AdminController{}, "get:Entries")
	beego.Router("/admin/entries/:id", &controllers.AdminController{}, "get:Entry")
	beego.Router("/admin/entries/update/:id", &controllers.AdminController{}, "post:UpdateEntry")
	beego.Router("/admin/entries/delete", &controllers.AdminController{}, "get:DeleteEntry")
	beego.Router("/admin/entries/new", &controllers.AdminController{}, "get:NewEntry;post:PostNewEntry")

	// Admin -> Collection
	beego.Router("/admin/collections", &controllers.AdminController{}, "get:Collections")
	beego.Router("/admin/collections/:id", &controllers.AdminController{}, "get:Collection")
	beego.Router("/admin/collections/update/:id", &controllers.AdminController{}, "post:UpdateCollection")
	beego.Router("/admin/collections/delete", &controllers.AdminController{}, "get:DeleteCollection")
	beego.Router("/admin/collections/new", &controllers.AdminController{}, "get:NewCollection;post:CreateCollection")
}
