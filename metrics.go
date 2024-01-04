package main

import (
	"fmt"
	"net/http"

	//"golang.org/x/tools/go/cfg"
)

func (cfg *apiConfig) handlerMetrics(w http.ResponseWriter, r *http.Request)  {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits: %d\n", cfg.fileserverHits)))
	
}

func (cfg *apiConfig)middlewareMetricsInc(next http.Handler) http.Handler  {
	return http.HandlerFunc(func (w http.ResponseWriter, r *http.Request)  {
		cfg.fileserverHits++
		next.ServeHTTP(w, r)
		
	})
	
}