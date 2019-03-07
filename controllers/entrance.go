package controllers

import (
	"github.com/astaxie/beego"
	"entrance/service/route"
)

type Message struct {
	Name string
	Age int
	Phone string
}
type EntranceController struct {
	beego.Controller
}

func (c *MainController) Entrance() {
	c.handle()
	c.ServeJSON()
}

func (c *MainController) handle() {

	apidata := route.GetRouteConfig()

	c.Data["json"] = map[string]interface{}{"success": 0, "message": "111","data":apidata}
}