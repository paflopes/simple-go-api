package main

import (
	"net/http"

	"github.com/paflopes/simple-go-api/src/routes"
)

func main() {
	router := routes.NewRouter()
	http.ListenAndServe(":8000", router)
}
