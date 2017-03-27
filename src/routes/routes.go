package routes

import (
	"net/http"

	"github.com/paflopes/simple-go-api/src/services/todo"

	"github.com/gorilla/mux"
)

type Route struct {
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []*Route

var routes = Routes{
	&Route{"GET", "/", todo.Index},
	&Route{"GET", "/todos", todo.TodoIndex},
	&Route{"GET", "/todos/{todoId}", todo.TodoShow},
}

func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).Path(route.Pattern).Handler(route.HandlerFunc)
	}
	return router
}
