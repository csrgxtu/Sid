package controllers

import (
	"github.com/astaxie/beego"
  "Sid/models"
  // "Sid/services"
)

type SidController struct {
	beego.Controller
}

func (this *SidController) Welcome() {
	var rt models.Result

	rt.Msg = "^_^"
	this.Ctx.ResponseWriter.WriteHeader(200)

	this.Data["json"] = &rt
	this.ServeJSON()
}
