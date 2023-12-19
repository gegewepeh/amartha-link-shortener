package server

import (
	"fmt"
	"log"
	"net/http"
	"os"

	linkHandler "link-shortener/internal/app/link_shortener/main/handler/v1"
	mainServices "link-shortener/internal/app/link_shortener/main/services/v1"
	network "link-shortener/internal/pkg/network"

	"github.com/gorilla/mux"
)

func Start() {
	port := os.Getenv("PORT")

	if port := os.Getenv("PORT"); port == "" {
		log.Fatalf("Missing port configuration")
	}

	// start redis and db pool
	pool := mainServices.GetPool()
	go pool.Start()

	router := setupRoutes(port)

	network.Serve(port, router)
}

func setupRoutes(port string) *mux.Router {
	router := mux.NewRouter()
	linkShortenerRouter := router.PathPrefix("/link-shortener").Subrouter()

	linkHandler.SetupRoutes(linkShortenerRouter)

	// default route
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Link Shortener service running on port: %s", port)
	})

	return router
}
