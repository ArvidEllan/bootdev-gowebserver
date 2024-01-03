package main

import (
	"log"
	"net/http"
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


	fs := http.FileServer(http.Dir("."))
	http.Handle("/",http.StripPrefix("/",fs))
	http.Handle("/assets/",http.StripPrefix("/assets",http.FileServer(http.Dir("assets"))))

    

	mux := http.NewServeMux()
	mux.Handle("/app/", apiCfg.middlewareMetricsInc(http.StripPrefix("/app", http.FileServer(http.Dir(filepathRoot)))))
	mux.HandleFunc("/healthz", handlerReadiness)
	mux.HandleFunc("/metrics", apiCfg.handlerMetrics)
	mux.HandleFunc("/reset", apiCfg.handlerReset)

	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: corsMux,
	}
	log.Printf("Serving files from %s on port : %s\n" ,filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func (cfg *apiConfig) handlerMetrics( next http.Handler) http.Handler {
	return http.HandlerFunc(func (w http.ResponseWriter, r * http.Request){
		cfg.fileserverHits++
		next.ServeHTTP(w,r)
	})
}

func (cfg *apiConfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
	})
}