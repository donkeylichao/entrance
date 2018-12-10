package main

import (
	_ "entrance/routers"
	"github.com/astaxie/beego"
	_ "entrance/models"
)

func main() {
	beego.Run()
}

