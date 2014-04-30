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

// Login

func (this *AdminController) Login() {
	this.TplNames = "admin/login.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
}

// Entry Operations

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
			this.Data["PostId"] = entry.Id + "/update"
		}
	}

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["MarkdownEnabled"] = true
}

func (this *AdminController) UpdateEntry() {
	entry := models.Entry{}
	err := this.ParseForm(&entry)
	if err != nil {
		panic(err)
	}

	entry.Id = this.Ctx.Input.Param(":id")
	fmt.Println(entry)
	nid, err := models.UpdateEntry(entry)

	this.TplNames = "admin/entry.tpl"
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	if err == nil {
		entry.Id = nid
	}
	this.Data["PostId"] = entry.Id + "/update"
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true
	this.Data["Message"] = "Update Successful"
}

func (this *AdminController) NewEntry() {
	this.TplNames = "admin/entry.tpl"
	fmt.Println(this.GetSession("user"))

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["PostId"] = "new"
	this.Data["EntryActive"] = true
	this.Data["MarkdownEnabled"] = true
}

func (this *AdminController) PostNewEntry() {
	entry := models.Entry{}
	err := this.ParseForm(&entry)
	if err != nil {
		panic(err)
	}

	nid, err := models.PostNewEntry(entry)
	this.TplNames = "admin/entry.tpl"
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	if err != nil {
		entry.Id = "new"
	} else {
		entry.Id = nid
	}
	this.Data["PostId"] = entry.Id + "/update"
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true
	this.Data["Message"] = "Post Successful"
}

// Collection Operations

func (this *AdminController) Collections() {
	this.TplNames = "admin/collection-list.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	// this.Data["Collections"] = entries
}

func (this *AdminController) Collection() {
	fmt.Println("Collection")
	this.TplNames = "admin/collection.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
}

func (this *AdminController) UpdateCollection() {

}

func (this *AdminController) NewCollection() {
	fmt.Println("NewCollection")
	this.TplNames = "admin/collection.tpl"
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	this.Data["PostId"] = "new"
}

func (this *AdminController) CreateCollection() {
	this.TplNames = "admin/collection-list.tpl"

	collection := models.Collection{}
	err := this.ParseForm(&collection)
	if err != nil {
		panic(err)
	}
}
