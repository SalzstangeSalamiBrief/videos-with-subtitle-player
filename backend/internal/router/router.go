package router

import (
	"net/http"
	"regexp"
	"slices"
	"strings"
)

var acceptedMethods = []string{http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodOptions}

type Middleware func(handlerFunc http.HandlerFunc) http.HandlerFunc

type RouterBase struct {
	routes      []Route
	middlewares []Middleware
}

type Router interface {
	AddRoute(route Route)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewRouter() *RouterBase {
	return &RouterBase{
		routes:      make([]Route, 0),
		middlewares: make([]Middleware, 0),
	}
}

func (routerBase *RouterBase) Build() http.HandlerFunc {
	if len(routerBase.middlewares) == 0 {
		return routerBase.handleRouting
	}

	middlewareParent := routerBase.handleRouting
	for _, middleware := range routerBase.middlewares {
		middlewareParent = middleware(middlewareParent)
	}

	return middlewareParent
}

func (routerBase *RouterBase) handleRouting(w http.ResponseWriter, r *http.Request) {
	if !slices.Contains(acceptedMethods, r.Method) {
		http.NotFound(w, r)
		return
	}

	pathWithoutPrefix := strings.TrimPrefix(r.URL.Path, "/api")
	hasMatched := false
	for _, route := range routerBase.routes {
		isPathMatching, pathMatchingErr := regexp.MatchString(route.Path, pathWithoutPrefix)
		isMethodMatching := route.Method == r.Method

		if pathMatchingErr != nil {
			http.NotFound(w, r)
			break
		}

		if !isPathMatching || !isMethodMatching {
			continue
		}

		if route.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			break
		}

		route.Handler(w, r)
		hasMatched = true
	}

	if !hasMatched {
		http.NotFound(w, r)
	}
}

func (routerBase *RouterBase) RegisterRoute(routeToAdd Route) *RouterBase {
	routerBase.routes = append(routerBase.routes, routeToAdd)
	return routerBase

}

func (routerBase *RouterBase) RegisterMiddleware(middlewareToAdd Middleware) *RouterBase {
	routerBase.middlewares = append(routerBase.middlewares, middlewareToAdd)
	return routerBase
}
