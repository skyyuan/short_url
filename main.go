package main

import (
	"flag"
	"log"
	"net/http"
	"short_url/lib/mux"
	"short_url/controllers"
)

var port string

func init() {
	flag.StringVar(&port, "port", ":8080", "port to listen")
	flag.Parse()
}

func main() {
	router()
}

func router() {
	router := mux.NewMux()
	router.Handle("/", &controllers.BaseController{})
	router.Handle("/api/url", &controllers.UrlController{})
	log.Println("ShortURL server will start at port " + port)
	log.Fatalln(http.ListenAndServe(port, router))
}
