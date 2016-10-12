package routers

import (
	"bulaoge/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/dbm", &controllers.DbmController{})
	beego.Router("/dbm/list/*", &controllers.DbmController{})
}
