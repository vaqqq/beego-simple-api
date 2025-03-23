package controllers

import (
	"time"

	"github.com/beego/beego/v2/server/web"
)

type TimeController struct {
	web.Controller
}

func (c *TimeController) GetTime() {
	c.Ctx.WriteString("Time: " + time.Now().Format(time.RFC1123))
}
