package routers

import (
	"bulaoge/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dbm", &controllers.DbmController{}, "*:Get")

	beego.Router("/dbm/list/:tableName:string", &controllers.DbmController{}, "*:List")

	beego.Router("/dbm/move", &controllers.DbmController{}, "get:Move")
}
