package controllers

type UrlController struct {
	BaseController
}

func (c *UrlController) Get() {
	c.success("This is urlController:Get")
}
