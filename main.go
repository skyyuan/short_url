package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"short_url/lib/mux"
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
	router.Handle("/hello/golang/", &BaseHandler{})
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
