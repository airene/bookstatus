package main

import (
	"github.com/airene/bookstatus/controllers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	//for print sql
	if beego.AppConfig.String("Runmode") == "dev" {
		orm.Debug = true
	}

	//route config
	beego.Router("/", &controllers.MainController{}, "get:Get")
	beego.Router("/add", &controllers.MainController{}, "get:Add")
	beego.Router("/add", &controllers.MainController{}, "post:Save")
	beego.Router("/borrow", &controllers.MainController{}, "get:Borrow")
	beego.Router("/borrow", &controllers.MainController{}, "post:BorrowSave")
	beego.Router("/back", &controllers.MainController{}, "get:Back")

	beego.Run()
}
