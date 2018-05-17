package controllers

import (
	"net/url"
	"short_url/models"
	"short_url/utils"
	"net/http"
	"fmt"
)

type UrlController struct {
}

func (handler *UrlController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	sourceUrl := req.URL.Query().Get("sourceUrl")
	_, err := url.ParseRequestURI(sourceUrl)
	if err != nil {
		fmt.Fprintln(resp, "Url is not a valid u: "+err.Error())
		return
	}
	fmt.Println("sourceUrl: ", sourceUrl)
	u := &models.Url{
		SourceUrl: sourceUrl,
	}
	u.GenId()
	u.ShortUrl = utils.IdToString(u.Id)
	u.Save()
	fmt.Fprintln(resp, u.SourceUrl)
}
