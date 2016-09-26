package controllers

import (
	"github.com/astaxie/beego"
  "Sid/models"
  "Sid/services"
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

func (this *SidController) Propagate() {
	var rt models.Result
	var vid = this.GetString(":vid")

	// get virus info according the vid
	err, virus := services.GetVirus(vid)
	if err != nil {
		rt.Msg = "o_o"
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Data["json"] = &rt
		this.ServeJSON()
	}
	beego.Info(virus)

	// get the raw infected data
	err, infects := services.GetInfects(vid)
	if err != nil {
		rt.Msg = "o_o"
		this.Ctx.ResponseWriter.WriteHeader(404)
		this.Data["json"] = &rt
		this.ServeJSON()
	}

	// get all related user info
	var UserIds = make([]string, len(infects))
	for i := 0; i < len(infects); i++ {
		UserIds = append(UserIds, infects[i].CarryId)
		UserIds = append(UserIds, infects[i].InfectedId)
	}
	beego.Info(UserIds)

	// assemble the data


	this.Data["json"] = &rt
	this.ServeJSON()
}
