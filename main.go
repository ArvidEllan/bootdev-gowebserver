package main

import (
	"log"
	"net/http"
)

func main() {
	const port = "8080"

	// Serve static files from the root directory
	fs := http.FileServer(http.Dir("."))
	http.Handle("/", http.StripPrefix("/", fs))

	// Serve static files from the "assets" directory
	http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))

	// Apply CORS middleware
	handler := middlewareCors(http.DefaultServeMux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: handler,
	}

	log.Printf("Serving on port: %s\n", port)
	log.Fatal(srv.ListenAndServe())
}

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")

		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		next.ServeHTTP(w, r)
	})
}
