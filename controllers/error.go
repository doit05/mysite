package controllers

import (
	"github.com/astaxie/beego"
)

type ErrorController struct {
	beego.Controller
}

func (c *ErrorController) Error404() {
	c.Data["content"] = "Not Found"
	c.TplName = "404.tpl"
	c.Render()
}

func (c *ErrorController) Error501() {
	c.Data["content"] = "server error"
	c.TplName = "501.tpl"
	c.Render()
}

func (c *ErrorController) ErrorDb() {
	c.Data["content"] = "database is now down"
	c.TplName = "dberror.tpl"
	c.Render()
}
