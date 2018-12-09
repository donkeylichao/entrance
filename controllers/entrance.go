package controllers

import (
	"github.com/astaxie/beego"
)

type EntranceController struct {
	beego.Controller
}

func (c *MainController) Entrance() {
	c.Data["Website"] = "beego.me"
	c.Data["Email"] = "astaxie@gmail.com"
	c.TplName = "index.tpl"
}
