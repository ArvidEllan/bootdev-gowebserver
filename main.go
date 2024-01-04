package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi/v5"
)
type apiConfig struct {
	fileserverHits int
}
func main()  {
	const filepathRoot = "."
	const port = "8080"

	apiCfg := apiConfig{
		fileserverHits: 0,
	}

	r := chi.NewRouter()
	fsHandler := apiCfg.middlewareMetricsInc(http.StripPrefix("/app",http.FileServer(http.Dir(filepathRoot))))
	r.Handle("/app",fsHandler)
	r.Handle("/app/*",fsHandler)
	r.Get("/healthz",handlerReadiness)
	r.Get("/metrics",apiCfg.handlerMetrics)
	r.Get("/reset",apiCfg.handlerReset)


	// fs := http.FileServer(http.Dir("."))
	// http.Handle("/",http.StripPrefix("/",fs))
	// http.Handle("/assets/",http.StripPrefix("/assets",http.FileServer(http.Dir("assets"))))

    

	// mux := http.NewServeMux()
	// mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	// mux.HandleFunc("/health", handlerReadiness)
	// mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	// mux.HandleFunc("/reset", apiCfg.handlerReset)

	// mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	corsMux := middlewareCors(r)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: corsMux,
	}
	log.Printf("Serving files from %s on port : %s\n" ,filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

