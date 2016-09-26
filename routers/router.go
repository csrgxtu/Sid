package routers

import (
	"Sid/controllers"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/sid/welcome", &controllers.SidController{}, "get:Welcome")
	beego.Router("/sid/propagation/:vid", &controllers.SidController{}, "get:Propagate")
}
