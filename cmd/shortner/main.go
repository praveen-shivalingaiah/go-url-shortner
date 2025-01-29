package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	myhttp "github.com/praveen-shivalingaiah/go-url-shortner/adapter/http"
	"github.com/praveen-shivalingaiah/go-url-shortner/adapter/storage"
	"github.com/praveen-shivalingaiah/go-url-shortner/app"
)

func main() {

	repo := storage.NewInMemoryURLRepository()

	service := app.NewShortnerService(repo)

	handler := myhttp.NewHandlerService(service)

	r := mux.NewRouter()
	handler.RegisterRoutes(r)

	log.Println("URL Shortner server is running on the port 8080")
	if err := http.ListenAndServe(":8080", r); err != nil {
		log.Fatalf("Failed to start the server: %v", err)
	}
}
