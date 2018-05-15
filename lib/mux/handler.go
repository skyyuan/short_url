package mux

import "net/http"

type mux struct {
	handlers    map[string]http.Handler
	handleFuncs map[string]func(resp http.ResponseWriter, req *http.Request)
}

func NewMux() *mux {
	return &mux{
		handlers:    make(map[string]http.Handler),
		handleFuncs: make(map[string]func(resp http.ResponseWriter, req *http.Request)),
	}
}

func (mux *mux) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	urlPath := req.URL.Path
	if hl, ok := mux.handlers[urlPath]; ok {
		hl.ServeHTTP(resp, req)
		return
	}
	if fn, ok := mux.handleFuncs[urlPath]; ok {
		fn(resp, req)
		return
	}
	http.NotFound(resp, req)
}

func (mux *mux) Handle(pattern string, handler http.Handler) {
	mux.handlers[pattern] = handler
}

func (mux *mux) HandleFunc(pattern string, fn func(resp http.ResponseWriter,
	req *http.Request)) {
	mux.handleFuncs[pattern] = fn
}
