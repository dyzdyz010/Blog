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

	renderTemplate(this.Ctx, "views/admin/dashboard.amber", this.Data)
}

// Login

func (this *AdminController) Login() {
	this.TplNames = "admin/login.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"

	renderTemplate(this.Ctx, "views/admin/login.amber", this.Data)
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
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"

	renderTemplate(this.Ctx, "views/admin/login.amber", this.Data)
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

	renderTemplate(this.Ctx, "views/admin/entry-list.amber", this.Data)
}

func (this *AdminController) Entry() {
	checkLogin(this)
	this.TplNames = "admin/entry.tpl"

	eid := this.Ctx.Input.Param(":id")
	fmt.Println(eid)
	if eid != "new" {
		entry, _ := models.EntryById(eid)
		if entry == nil {
			this.Abort("404")
			return
		} else {
			this.Data["Entry"] = entry
			this.Data["PostId"] = "update/" + entry.Id
		}
	}

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["MarkdownEnabled"] = true

	collections, err := models.CollectionsByUser(this.GetSession("user").(string))
	if err != nil {
		panic(err)
	}
	this.Data["Collections"] = collections
	fmt.Println(this.Data["Entry"])

	renderTemplate(this.Ctx, "views/admin/entry.amber", this.Data)
}

func (this *AdminController) UpdateEntry() {
	entry := models.Entry{}
	err := this.ParseForm(&entry)
	if err != nil {
		panic(err)
	}

	entry.Id = this.Ctx.Input.Param(":id")
	entry.Author = this.GetSession("user").(string)
	entry.Collection = this.Input().Get("collection")
	err = models.UpdateEntry(entry)
	if err != nil {
		panic(err)
	}

	this.TplNames = "admin/entry.tpl"
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	this.Data["PostId"] = "update/" + entry.Id
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true
	collections, err := models.CollectionsByUser(this.GetSession("user").(string))
	if err != nil {
		panic(err)
	}
	this.Data["Collections"] = collections
	this.Data["Message"] = "Update Successful"

	renderTemplate(this.Ctx, "views/admin/entry.amber", this.Data)
}

func (this *AdminController) DeleteEntry() {
	this.TplNames = "admin/entry-list.tpl"

	id := this.GetString("id")
	fmt.Println("Entry Id: ", id)
	err := models.DeleteEntry(id)
	if err != nil {
		panic(err)
	}

	this.Redirect("/admin/entries", 302)
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
	entry.Collection = this.Input().Get("collection")
	nid, err := models.PostNewEntry(entry)
	this.TplNames = "admin/entry.tpl"
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["EntryActive"] = true
	if err != nil {
		this.Data["PostId"] = "new"
		this.Data["Message"] = err.Error()
	} else {
		this.Data["PostId"] = "update/" + nid
		this.Data["Message"] = "Post Successful"

		collections, err := models.CollectionsByUser(this.GetSession("user").(string))
		if err != nil {
			panic(err)
		}
		this.Data["Collections"] = collections
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

	renderTemplate(this.Ctx, "views/admin/collection-list.amber", this.Data)
}

func (this *AdminController) Collection() {
	fmt.Println("Collection")
	checkLogin(this)
	this.TplNames = "admin/collection.tpl"

	id := this.Ctx.Input.Param(":id")
	collection, err := models.CollectionById(id)
	if err != nil {
		panic(err)
	}

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	this.Data["Collection"] = collection
	this.Data["PostId"] = "update/" + id

	renderTemplate(this.Ctx, "views/admin/collection.amber", this.Data)
}

func (this *AdminController) UpdateCollection() {
	this.TplNames = "admin/collection.tpl"

	collection := models.Collection{}
	err := this.ParseForm(&collection)
	if err != nil {
		panic(err)
	}

	collection.Id = this.Ctx.Input.Param(":id")
	collection.Author = this.GetSession("user").(string)
	nid, err := models.UpdateCollection(collection)
	if err != nil {
		this.Data["PostId"] = "update/" + collection.Id
		this.Data["Message"] = err.Error()
	} else {
		this.Data["PostId"] = "update/" + nid
		this.Data["Message"] = "Update Successful"
	}
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["Collection"] = collection
	this.Data["CollectionActive"] = true

	renderTemplate(this.Ctx, "views/admin/collection.amber", this.Data)
}

func (this *AdminController) NewCollection() {
	fmt.Println("New Collection")
	checkLogin(this)
	this.TplNames = "admin/collection.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["CollectionActive"] = true
	this.Data["PostId"] = "new"

	renderTemplate(this.Ctx, "views/admin/collection.amber", this.Data)
}

func (this *AdminController) DeleteCollection() {
	fmt.Println("Delete Collection")
	this.TplNames = "admin/collection-list.tpl"

	id := this.GetString("id")
	err := models.DeleteCollection(id)
	if err != nil {
		panic(err)
	}

	this.Redirect("/admin/collections", 302)
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
		this.Data["PostId"] = "update/" + cid
		this.Data["Message"] = "Post Successful"
	}
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["Collection"] = collection
	this.Data["CollectionActive"] = true

	renderTemplate(this.Ctx, "views/admin/collection.amber", this.Data)
}
