package controllers

type RedirectController struct {
	BaseController
}

// 302 redirect to "/api/url"
func (r RedirectController) Get() {
	//r.Ctx.ResponseWriter.Header().Set("Location", "/api/url")
	//r.Ctx.ResponseWriter.WriteHeader(302)
	r.Redirect("/api/url", 302)
}
