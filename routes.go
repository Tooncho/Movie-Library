package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name       string
	Method     string
	Pattern    string
	HandleFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.
			Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(route.HandleFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"MovieList",
		"GET",
		"/movies",
		MovieList,
	},
	Route{
		"MovieShow",
		"GET",
		"/movies/{id}",
		MovieShow,
	},
	Route{
		"MovieAdd",
		"POST",
		"/movie",
		MovieAdd,
	},
	Route{
		"MovieUpdate",
		"PUT",
		"/movie/{id}",
		MovieUpdate,
	},
	Route{
		"MovieRemove",
		"DELETE",
		"/movie/{id}",
		MovieRemove,
	},
}
