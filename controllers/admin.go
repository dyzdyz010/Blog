package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/models"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) Dashboard() {
	this.TplNames = "admin/dashboard.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["DashboardActive"] = true
}

func (this *AdminController) Entries() {
	this.TplNames = "admin/entry-list.tpl"

	entries := models.PublishedEntries()

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["Entries"] = entries
}

func (this *AdminController) Entry() {
	this.TplNames = "admin/entry.tpl"

	eid := this.Ctx.Input.Param(":id")
	fmt.Println(eid)
	if eid != "new" {
		entry := models.EntryById(eid)
		if entry == nil {
			this.Abort("404")
			return
		} else {
			this.Data["Entry"] = entry
		}
	}

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["MarkdownEnabled"] = true
}

func (this *AdminController) PostEntry() {
	entry := models.Entry{}
	err := this.ParseForm(&entry)
	if err != nil {
		panic(err)
	}

	entry.Id = this.Ctx.Input.Param(":id")
	fmt.Println(entry)

	this.ServeJson()
}

func (this *AdminController) Collections() {
	this.TplNames = "admin/collection-list.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	// this.Data["Collections"] = entries
}

func (this *AdminController) Collection() {
	this.TplNames = "admin/collection.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
}

func (this *AdminController) CreateCollection() {
	this.TplNames = "admin/collection.tpl"
}
