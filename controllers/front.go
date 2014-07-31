package controllers

import (
	// "fmt"
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
	// fmt.Println(this.GetString("prev") == "")
	dir := ""
	dirId := ""

	// Configure direction
	if this.GetString("prev") != "" {
		dir = "prev"
		dirId = this.GetString("prev")
	}
	if this.GetString("next") != "" {
		dir = "next"
		dirId = this.GetString("next")
	}

	// Main Nav
	this.Data["HomeActive"] = "active"

	// Data Source
	entries, havePrev, haveNext, err := models.PublishedEntries(dir, dirId)
	if err != nil {
		panic(err)
	}
	// fmt.Println(len(entries))
	this.Data["Entries"] = entries
	this.Data["Pos"] = ""
	if len(entries) != 0 {
		this.Data["FirstId"] = entries[0].Id
		this.Data["LastId"] = entries[len(entries)-1].Id
	}

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["HavePrev"] = havePrev
	this.Data["HaveNext"] = haveNext

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
	dir := ""
	dirId := ""

	// Configure direction
	if this.GetString("prev") != "" {
		dir = "prev"
		dirId = this.GetString("prev")
	}
	if this.GetString("next") != "" {
		dir = "next"
		dirId = this.GetString("next")
	}

	cid := this.Ctx.Input.Param(":id")
	collection, err := models.CollectionById(cid)
	if err != nil {
		panic(err)
	}

	entries, havePrev, haveNext, err := models.EntriesByCollection(cid, dir, dirId)
	if err != nil {
		panic(err)
	}

	this.TplNames = "entry-list.tpl"
	this.Data["Title"] = collection.Title
	this.Data["Subtitle"] = collection.Subtitle
	this.Data["Entries"] = entries
	this.Data["Pos"] = this.UrlFor("FrontController.Collection", ":id", cid)

	if len(entries) != 0 {
		this.Data["FirstId"] = entries[0].Id
		this.Data["LastId"] = entries[len(entries)-1].Id
	}

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["HavePrev"] = havePrev
	this.Data["HaveNext"] = haveNext

	renderTemplate(this.Ctx, "views/entry-list.amber", this.Data)
}

func (this *FrontController) Entry() {
	eid := this.Ctx.Input.Param(":id")
	entry, _ := models.EntryById(eid)
	// fmt.Println(entry)

	this.TplNames = "entry.tpl"

	// fmt.Println(entry.Content)
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true

	renderTemplate(this.Ctx, "views/entry.amber", this.Data)
}
