package httpserver

import "net/http"

type HTTPMethod string

const (
	MethodGet     HTTPMethod = "GET"
	MethodHead    HTTPMethod = "HEAD"
	MethodPost    HTTPMethod = "POST"
	MethodPut     HTTPMethod = "PUT"
	MethodPatch   HTTPMethod = "PATCH"
	MethodDelete  HTTPMethod = "DELETE"
	MethodConnect HTTPMethod = "CONNECT"
	MethodOptions HTTPMethod = "OPTIONS"
	MethodTrace   HTTPMethod = "TRACE"
)

type Route struct {
	Method  HTTPMethod
	Path    string
	Handler http.HandlerFunc
}

func NewRoute(
	method HTTPMethod,
	path string,
	handler http.HandlerFunc,
) *Route {
	return &Route{
		Method:  method,
		Path:    path,
		Handler: handler,
	}
}
