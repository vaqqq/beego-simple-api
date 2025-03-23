package controllers

import (
	"github.com/beego/beego/v2/server/web"
)

type PostController struct {
	web.Controller
}

func (c *PostController) PostMessage() {
	data := c.GetString("message")
	c.Ctx.WriteString("You have sent: " + data)
}
