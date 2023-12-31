package router

import "net/http"

type Route struct {
	Path    string
	Handler func(w http.ResponseWriter, r *http.Request)
	Method  string
}
