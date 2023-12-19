package handler

import (
	network "link-shortener/internal/pkg/network"

	"github.com/gorilla/mux"
)

func vOneRoutes(router *mux.Router) {
	vOneZeroRouter := router.PathPrefix("/v1").Subrouter()
	vOneZeroRouter.Handle("/links", network.HTTPHandler(getAllLinks)).Methods("GET")
	vOneZeroRouter.Handle("/slug/{id}", network.HTTPHandler(getBySlugId)).Methods("GET")
	vOneZeroRouter.Handle("/slug", network.HTTPHandler(createSlug)).Methods("POST")
	vOneZeroRouter.Handle("/slug/{id}", network.HTTPHandler(updateSlug)).Methods("PUT")
	vOneZeroRouter.Handle("/users", network.HTTPHandler(createUser)).Methods("POST")
}

// SetupRoutes is function to setup API (and Admin) routes
func SetupRoutes(router *mux.Router) {
	vOneRoutes(router)
}
