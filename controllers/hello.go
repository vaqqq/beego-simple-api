package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type HelloController struct {
	web.Controller
}

func (c *HelloController) GetHello() {
	c.Ctx.WriteString("Hello World!")
}
