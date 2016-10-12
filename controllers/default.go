package controllers

import (
	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.Data["Website"] = "www.dbm.com"
	c.Data["Email"] = "lmw@dbm.com"
	c.TplName = "index.html"
}
