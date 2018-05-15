package controllers

import (
	"net/url"
	"short_url/models"
	"short_url/utils"
)

type UrlController struct {
	BaseController
}

func (c *UrlController) Get() {
	sourceUrl := c.Ctx.Input.Query("sourceUrl")
	_, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		c.error("Url is not a valid url: " + err.Error())
		return
	}
	u := &models.Url{SourceUrl: sourceUrl,}
	u.GenId()
	u.ShortUrl = utils.IdToString(u.Id)
	u.Save()
	c.success(u)
}
