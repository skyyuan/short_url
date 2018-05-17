package controllers

import (
	"net/http"
	"html/template"
	"os"
)

type BaseController struct {
}

func (b *BaseController) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	t := template.New("index template")
	t, err := t.ParseFiles("views/index.html")
	if err != nil {
		panic(err)
	}
	t.Execute(os.Stdout, "Hello")
	//resp.Header().Set("Location", "http://www.baidu.com")
	//resp.WriteHeader(302)
}
