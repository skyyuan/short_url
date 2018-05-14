package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"short_url/lib/mux"
	"github.com/astaxie/beego"
	"short_url/controllers"
)

var port string

func main() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()

	beego.RESTRouter("/api/url", &controllers.UrlController{})

	beego.Router("/*", &controllers.RedirectController{})

	beego.Run(port)
}

func myRouter() {
	router := mux.NewMuxHandler()
	router.Handle("/hello/golang/", &BaseHandler{})
	router.HandleFunc("/hello/world", func(resp http.ResponseWriter, req *http.Request) {
		resp.Write([]byte("Hello World"))
	})
	log.Println("ShortURL server will start at port " + port)
	log.Fatalln(http.ListenAndServe(port, router))
}

type BaseHandler struct {
}

func (handler *BaseHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	fmt.Println("url path => ", req.URL.Path)
	fmt.Println("url param a => ", req.URL.Query().Get("a"))

	resp.Write([]byte("Hello Golang"))
}
