package rets

import (
	"net/http"
	"regexp"
	"strings"
)

const routeBlock int = 128

// Route associates a basic URL route based on the request method and a matching
// RegExp to a specified HTTP handler function.
type Route struct {
	Methods string
	RegExp  *regexp.Regexp
	Handler http.HandlerFunc
}

// Router contains the slice that can contain registered routes.
type Router struct {
	routes []Route
}

// NewRouter creates and initializes a new Router.
func NewRouter() (router *Router) {
	router = new(Router)
	router.routes = make([]Route, 0, routeBlock)
	return
}

func (router *Router) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	// We preemptively set the response content type to JSON, as ReTS APIs
	// don't output anything else, even when there is a HTTP error.
	writer.Header().Set("Content-Type", "application/json")

	// Try to route the request, returns 404 if no route matched the path
	for _, route := range router.routes {
		if strings.Contains(route.Methods, request.Method) &&
			route.RegExp.MatchString(request.URL.Path) {
			route.Handler(writer, request)
			return
		}
	}
	HandlerBasicHTTP(writer, 404)
}

// AddRoute adds a given Route to a Router.
func (router *Router) AddRoute(route Route) {
	// Custom slice growing management. We don't really need exponential
	// growth for storing routes, so we use fixed-size blocks of 128 routes.
	if len(router.routes) >= cap(router.routes) {
		tmp := make([]Route, len(router.routes), cap(router.routes)+routeBlock)
		copy(tmp, router.routes)
		router.routes = tmp
	}
	router.routes = router.routes[:len(router.routes)+1]
	router.routes[len(router.routes)-1] = route
}
