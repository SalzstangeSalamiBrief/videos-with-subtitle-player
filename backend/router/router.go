package router

import (
	"fmt"
	"net/http"
	"regexp"
	"slices"
)

type routes []Route

var acceptedMethods = []string{http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodOptions}
var Routes = make(routes, 0)

func Router(w http.ResponseWriter, r *http.Request) {
	if !slices.Contains(acceptedMethods, r.Method) {
		ErrorHandler(w, fmt.Sprintf("The method '%v' is not acceptable", r.Method), http.StatusBadRequest)
		return
	}

	hasMatched := false
	for path, route := range Routes {
		isPathMatching, pathMatchingErr := regexp.MatchString(route.Path, r.URL.Path)
		isMethodMatching := route.Method == r.Method

		if pathMatchingErr != nil {
			ErrorHandler(w, fmt.Sprintf("Could not get resource '%v'", path), http.StatusBadRequest)
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
		ErrorHandler(w, fmt.Sprintf("Could not get resource '%v' with method '%v'", r.URL.Path, r.Method), http.StatusBadRequest)
	}
}

func (r *routes) AddRoute(routeToAdd Route) {
	*r = append(*r, routeToAdd)
}
