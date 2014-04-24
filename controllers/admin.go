package controllers

import (
	"Blog/models"
	"fmt"
	"github.com/astaxie/beego"
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
	entry := models.EntryById(eid)
	fmt.Println(entry)

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["MarkdownEnabled"] = true

	if entry != nil {
		this.Data["Entry"] = entry
	}
}

func (this *AdminController) PostEntry() {
	entry := models.Entry{}
	entry.Title = this.GetString("title")
	fmt.Println(entry)

	this.ServeJson()
}
