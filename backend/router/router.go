package router

import (
	"fmt"
	"net/http"
	"regexp"
)

type routes []Route

var acceptedMethods = [6]string{http.MethodGet, http.MethodDelete, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodOptions}
var Routes = make(routes, 0)

func HandleRouting(w http.ResponseWriter, r *http.Request) {
	useCors(w, r)

	quitChannel := make(chan bool)

	if validateHttpMethod(w, r) == false {
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

		if isPathMatching == false || isMethodMatching == false {
			continue
		}

		if route.Method == http.MethodOptions {
			w.WriteHeader(http.StatusOK)
			break
		}

		go func() {
			route.Handler(w, r)
			hasMatched = true
			quitChannel <- true
		}()

		<-quitChannel // wait till the goroutine is completed and discard value
	}

	if hasMatched == false {
		ErrorHandler(w, fmt.Sprintf("Could not get resource '%v' with method '%v'", r.URL.Path, r.Method), http.StatusBadRequest)
	}
}
func (r *routes) AddRoute(routeToAdd Route) {
	*r = append(*r, routeToAdd)
}

func validateHttpMethod(w http.ResponseWriter, r *http.Request) bool {
	isAccepted := false

	for _, method := range acceptedMethods {
		if method == r.Method {
			isAccepted = true
			break
		}
	}

	if isAccepted == false {
		ErrorHandler(w, fmt.Sprintf("The method '%v' is not acceptable", r.Method), http.StatusBadRequest)
	}

	return isAccepted
}
