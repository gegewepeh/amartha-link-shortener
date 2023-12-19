package network

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Serve(port string, router *mux.Router) {
	log.Printf("Starting server at port %v\n", port)

	log.Fatal(http.ListenAndServe(":"+port, router))
}
