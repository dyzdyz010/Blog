package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/models"
)

type FrontController struct {
	beego.Controller
}

func (this *FrontController) Prepare() {
	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
}

func (this *FrontController) Home() {
	// Main Nav
	this.Data["HomeActive"] = "active"

	// Data Source
	this.Data["Entries"] = models.PublishedEntries()

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["LeftPage"] = "disabled"

	this.TplNames = "entry-list.tpl"

	renderTemplate(this.Ctx, "views/entry-list.amber", this.Data)
}

func (this *FrontController) Collections() {
	this.TplNames = "collection-list.tpl"

	// Main Nav
	this.Data["CollectionActive"] = "active"

	// Data Source
	collections, err := models.AllCollections()
	if err != nil {
		panic(err)
	}
	this.Data["Collections"] = collections

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["LeftPage"] = "disabled"

	renderTemplate(this.Ctx, "views/collection-list.amber", this.Data)
}

func (this *FrontController) Collection() {
	cid := this.Ctx.Input.Param(":id")
	collection, err := models.CollectionById(cid)
	if err != nil {
		panic(err)
	}

	entries, err := models.EntriesByCollection(cid)
	if err != nil {
		// panic(err)
	}

	this.TplNames = "entry-list.tpl"
	this.Data["Title"] = collection.Title
	this.Data["Subtitle"] = collection.Subtitle
	this.Data["Entries"] = entries

	renderTemplate(this.Ctx, "views/entry-list.amber", this.Data)
}

func (this *FrontController) Entry() {
	eid := this.Ctx.Input.Param(":id")
	entry, _ := models.EntryById(eid)
	// fmt.Println(entry)

	this.TplNames = "entry.tpl"

	fmt.Println(entry.Content)
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true

	renderTemplate(this.Ctx, "views/entry.amber", this.Data)
}
