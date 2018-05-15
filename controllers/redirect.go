package controllers

import (
	"strings"
	"short_url/utils"
	"short_url/models"
	"html/template"
)

type RedirectController struct {
	BaseController
}

// 302 redirect to "/api/url"
func (r RedirectController) Get() {
	//r.Ctx.ResponseWriter.Header().Set("Location", "/api/url")
	//r.Ctx.ResponseWriter.WriteHeader(302)
	if r.Ctx.Request.URL.Path == "/" {
		t, err := template.ParseFiles("views/index.html")
		if err != nil {
			panic(err)
		}
		data := map[string]interface{}{
			"title_name": "短链接服务",
		}
		t.Execute(r.Ctx.ResponseWriter, data)
		//r.TplName = "index.html"
		return
	}
	urlPath := r.Ctx.Request.URL.Path
	urlPath = strings.Trim(urlPath, "/")

	id := utils.StringToId(urlPath)

	u := &models.Url{
		Id: id,
	}
	err := u.FindById()
	if err != nil {
		r.error("没有找到对应的短链接" + err.Error())
	}
	r.Redirect(u.SourceUrl, 302)
}
