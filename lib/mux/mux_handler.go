package mux

import "net/http"

type muxHandler struct {
	handlers    map[string]http.Handler
	handleFuncs map[string]func(resp http.ResponseWriter, req *http.Request)
}

func NewMuxHandler() *muxHandler {
	return &muxHandler{
		handlers:    make(map[string]http.Handler),
		handleFuncs: make(map[string]func(resp http.ResponseWriter, req *http.Request)),
	}
}

func (mux *muxHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	// 分发请求
	// 精确匹配
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

func (mux *muxHandler) Handle(pattern string, handler http.Handler) {
	mux.handlers[pattern] = handler
}

func (mux *muxHandler) HandleFunc(pattern string, fn func(resp http.ResponseWriter,
	req *http.Request)) {
	mux.handleFuncs[pattern] = fn
}
