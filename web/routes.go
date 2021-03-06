package web

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}

	return router
}

var routes = Routes{
	Route{"Index", "GET", "/", Index },
	Route{"NewBuild", "POST", "/builds", NewBuild },
	Route{"Build", "GET", "/builds/{id}", Build },
	Route{"Output", "GET", "/output/{build_id}", Output },
}