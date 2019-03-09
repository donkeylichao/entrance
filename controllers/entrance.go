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

/**
 网关入口
 */
func (c *MainController) Entrance() {

	requestParams := c.parseRequestParameters()

	if requestParams["method"] == "OPTIONS" {
		c.Data["json"] = map[string]interface{}{"status": 1,"message": "success"}
	} else {

		matchRoute := route.GetGatewayService(requestParams)

		c.Data["json"] = map[string]interface{}{"status": 1, "message": "success", "data":matchRoute}
	}
	c.ServeJSON()
}

/**
 解析请求参数
 */
func (c *MainController) parseRequestParameters() map[string]interface{} {

	/* 获取请求参数 */
	method := c.Ctx.Request.Method
	header := c.Ctx.Request.Header
	body := c.Ctx.Request.Body
	form := c.Ctx.Request.Form
	multipart := c.Ctx.Request.MultipartForm
	path := c.Ctx.Request.RequestURI

	requestParams := make(map[string]interface{})

	requestParams["method"] = method
	requestParams["header"] = header
	requestParams["body"] = body
	requestParams["form"] = form
	requestParams["multipart"] = multipart
	requestParams["path"] = path

	return requestParams
}