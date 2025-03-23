package routers

import (
	"GoApp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	beego.Router("/", &controllers.MainController{})

	beego.Router("/hello", &controllers.HelloController{}, "get:GetHello")
	beego.Router("/time", &controllers.TimeController{}, "get:GetTime")
	beego.Router("/post", &controllers.PostController{}, "post:PostMessage")

}
