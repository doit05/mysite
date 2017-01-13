package controllers

import (
	"github.com/astaxie/beego"
)

type ADController struct {
	beego.Controller
}

func (c *ADController) Get() {
	c.Data["Website"] = "www.doit05.top"
	c.Data["Email"] = "768068275@qq.com"
	c.TplName = "ads.tpl"
	c.Render()
}
