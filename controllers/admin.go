package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/models"
)

type AdminController struct {
	beego.Controller
}

func checkLogin(ac *AdminController) {
	name := ac.GetSession("user")
	if name == nil {
		ac.Redirect("/admin/login", 302)
	}
}

func (this *AdminController) Dashboard() {
	checkLogin(this)

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

func (this *AdminController) PostLogin() {
	this.TplNames = "admin/login.tpl"

	author := models.Author{}
	err := this.ParseForm(&author)
	if err != nil {
		panic(err)
	}

	a, err := models.AuthorByName(author.Name)
	if err != nil {
		this.Data["Message"] = err.Error()
	} else if models.Hash(author.Password) != a.Password {
		this.Data["Message"] = "Wrong password."
	} else {
		this.SetSession("user", author.Name)
		this.Redirect("/admin", 302)
	}
}

// Entry Operations

func (this *AdminController) Entries() {
	checkLogin(this)
	this.TplNames = "admin/entry-list.tpl"

	entries := models.PublishedEntries()

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["Entries"] = entries
}

func (this *AdminController) Entry() {
	checkLogin(this)
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
	entry.Author = this.GetSession("user").(string)
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
	checkLogin(this)
	this.TplNames = "admin/entry.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["PostId"] = "new"
	this.Data["EntryActive"] = true
	this.Data["MarkdownEnabled"] = true
	collections, err := models.CollectionsByUser(this.GetSession("user").(string))
	if err != nil {
		panic(err)
	}
	this.Data["Collections"] = collections
}

func (this *AdminController) PostNewEntry() {
	entry := models.Entry{}
	err := this.ParseForm(&entry)
	if err != nil {
		panic(err)
	}
	entry.Author = this.GetSession("user").(string)
	nid, err := models.PostNewEntry(entry)
	this.TplNames = "admin/entry.tpl"
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	if err != nil {
		this.Data["PostId"] = "new"
		this.Data["Message"] = err.Error()
	} else {
		this.Data["PostId"] = nid + "/update"
		this.Data["Message"] = "Post Successful"
	}
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true
}

// Collection Operations

func (this *AdminController) Collections() {
	checkLogin(this)
	this.TplNames = "admin/collection-list.tpl"

	collections, err := models.CollectionsByUser(this.GetSession("user").(string))
	if err != nil {
		panic(err)
	}

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	this.Data["Collections"] = collections
}

func (this *AdminController) Collection() {
	checkLogin(this)
	this.TplNames = "admin/collection.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
}

func (this *AdminController) UpdateCollection() {

}

func (this *AdminController) NewCollection() {
	checkLogin(this)
	this.TplNames = "admin/collection.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	this.Data["PostId"] = "new"
}

func (this *AdminController) CreateCollection() {
	this.TplNames = "admin/collection.tpl"

	collection := models.Collection{}
	err := this.ParseForm(&collection)
	if err != nil {
		panic(err)
	}
	fmt.Println(this.GetSession("user"))
	collection.Author = this.GetSession("user").(string)
	cid, err := models.CreateCollection(collection)
	if err != nil {
		this.Data["PostId"] = "new"
		this.Data["Message"] = err.Error()
	} else {
		this.Data["PostId"] = cid + "/update"
		this.Data["Message"] = "Post Successful"
	}
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["Collection"] = collection
	this.Data["CollectionActive"] = true
}
