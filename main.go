package main

import (
	_ "entrance/routers"
	"github.com/astaxie/beego"
	_ "entrance/models"
	_ "entrance/help"
)

func main() {
	beego.Run()
}

