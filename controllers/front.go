package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/dyzdyz010/Blog/models"
	"time"
)

type FrontController struct {
	beego.Controller
}

func (this *FrontController) Home() {
	value := "2014-04-20T15:13:09+08:00"
	t, _ := time.Parse(time.RFC3339, value)
	timeStr := fmt.Sprintf("%s %d, %d", t.Month().String(), t.Day(), t.Year())
	fmt.Println(timeStr)

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	// Main Nav
	this.Data["HomeActive"] = "active"

	// Data Source
	this.Data["Entries"] = models.PublishedEntries()

	// Pagination
	this.Data["PageNav"] = "true"
	this.Data["LeftPage"] = "disabled"

	this.TplNames = "entry-list.tpl"
}

func (this *FrontController) Collections() {
	this.TplNames = "collection-list.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
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
}

func (this *FrontController) Collection() {
	cid := this.Ctx.Input.Param(":id")
	collection, err := models.CollectionById(cid)
	if err != nil {
		panic(err)
	}

	entries, _ := models.EntriesByCollection(cid)

	this.TplNames = "entry-list.tpl"
	this.Data["Title"] = collection.Title
	this.Data["Subtitle"] = collection.Subtitle
	this.Data["Entries"] = entries
}

func (this *FrontController) Entry() {
	eid := this.Ctx.Input.Param(":id")
	entry := models.EntryById(eid)
	// fmt.Println(entry)

	this.TplNames = "entry.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	fmt.Println(entry.Content)
	this.Data["Entry"] = entry
	this.Data["MarkdownEnabled"] = true
}
