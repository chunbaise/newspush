package controllers

import (
	"newspush/models"

	"github.com/astaxie/beego"
)

// Operations about object
type ChannelController struct {
	beego.Controller
}

// @router /all [get]
func (this *ChannelController) ReturnMsg() {
	this.Data["json"] = models.OneNews
	this.ServeJSON()
}

// @router /post [post]
func (this *ChannelController) ReturnPost() {
	this.Data["json"] = models.OneNews
	this.ServeJSON()
}
