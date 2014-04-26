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

	this.TplNames = "index.tpl"
}

func (this *FrontController) Collections() {

}

func (this *FrontController) Entry() {
	eid := this.Ctx.Input.Param(":id")
	entry := models.EntryById(eid)
	// fmt.Println(entry)

	this.TplNames = "entry.tpl"

	this.Data["Title"] = "Moonlightter"
	this.Data["Subtitle"] = "My Programming Life"
	this.Data["Entry"] = entry
}
