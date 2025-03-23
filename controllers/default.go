package controllers

import (
	beego "github.com/beego/beego/v2/server/web"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
	c.Data["GitHub"] = "https://github.com/vaqqq"
	c.Data["Modules"] = []map[string]string{
		{"Name": "Beego v2", "Link": "https://github.com/beego/beego"},
		{"Name": "Go 1.20+", "Link": "https://golang.org/dl/"},
		{"Name": "Custom Rate Limiting", "Link": "https://pkg.go.dev/golang.org/x/time/rate"},
		{"Name": "XSRF Protection", "Link": "https://beego.vip/docs/mvc/controller/xsrf.md"},
	}
}
