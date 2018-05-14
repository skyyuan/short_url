package controllers

import "github.com/astaxie/beego"

type BaseController struct {
	beego.Controller
}

func (b *BaseController) output(res string) {
	b.Ctx.WriteString(res)
}

func (b *BaseController) success(res interface{}) {
	b.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  res,
	}
	b.ServeJSON()
}

func (b *BaseController) error(errorMsg string) {
	b.Data["json"] = map[string]interface{}{
		"success": true,
		"result":  errorMsg,
	}
	b.ServeJSON()
}
