package main

import (
	"log"
	"net/http"
)
func main()  {
	const filepathRoot = "."
	const port = "8080"


	fs := http.FileServer(http.Dir("."))
	http.Handle("/",http.StripPrefix("/",fs))
	http.Handle("/assets/",http.StripPrefix("/assets",http.FileServer(http.Dir("assets"))))

    

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr: ":" + port,
		Handler: corsMux,
	}
	log.Printf("Serving files from %s on port : %s\n" ,filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(http.StatusText(http.StatusOK)))
}
