package web

import (
	"github.com/gin-gonic/gin"
)

type HttpMethod string

const (
	MethodGet     HttpMethod = "GET"
	MethodHead    HttpMethod = "HEAD"
	MethodPost    HttpMethod = "POST"
	MethodPut     HttpMethod = "PUT"
	MethodPatch   HttpMethod = "PATCH" // RFC 5789
	MethodDelete  HttpMethod = "DELETE"
	MethodConnect HttpMethod = "CONNECT"
	MethodOptions HttpMethod = "OPTIONS"
	MethodTrace   HttpMethod = "TRACE"
)

type Router struct {
	Method  HttpMethod
	Path    string
	Handler gin.HandlerFunc
}

func (r *Router) Register(rg *gin.RouterGroup) {
	switch r.Method {
	case MethodGet:
		rg.GET(r.Path, r.Handler)
	case MethodPost:
		rg.POST(r.Path, r.Handler)
	case MethodPatch:
		rg.PATCH(r.Path, r.Handler)
	case MethodDelete:
		rg.DELETE(r.Path, r.Handler)
	case MethodHead:
		rg.HEAD(r.Path, r.Handler)
	case MethodOptions:
		rg.OPTIONS(r.Path, r.Handler)
	case MethodPut:
		rg.PUT(r.Path, r.Handler)
	}
}
