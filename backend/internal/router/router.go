package router

import (
	"net/http"
	"regexp"
	"slices"
	"strings"
)

var acceptedMethods = []string{http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodOptions}

type RouterBase struct {
	routes []Route
}

type Router interface {
	AddRoute(route Route)
	ServeHTTP(http.ResponseWriter, *http.Request)
}

func NewRouter() *RouterBase {
	return &RouterBase{
		routes: make([]Route, 0),
	}
}

func (routerBase *RouterBase) ServeHTTP(w http.ResponseWriter, r *http.Request) {
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
		http.NotFound(w, r) // TODO TEST THIS REFACTOR
	}
}

func (routerBase *RouterBase) RegisterRoute(routeToAdd Route) {
	routerBase.routes = append(routerBase.routes, routeToAdd)
}
