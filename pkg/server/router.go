package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mijailr/go-learn/pkg/api"
	"log"
	"net/http"
	"os"
)

var port = ":9090"

func Run() error {
	if pt := os.Getenv("PORT"); pt != "" {
		port = fmt.Sprintf(":%s", pt)
	}

	log.Printf("Running server on port %s", port)
	return http.ListenAndServe(port, router())
}

func router() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.Handle("/api", api.Handler())

	return router
}
