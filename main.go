package main

import (
	"github.com/gorilla/mux"
	"github.com/just1689/pb-go-server/api"
	"log"
	"net/http"
	"time"
)

func main() {

	log.Println("Oh hai 🚀 Lets Go 🎠")

	r := mux.NewRouter()
	r.HandleFunc("/api", api.HandleAPICall)
	r.Handle("*", http.FileServer(http.Dir("./")))
	http.Handle("/", r)

	srv := &http.Server{
		Handler:      r,
		Addr:         "127.0.0.1:8000",
		WriteTimeout: 5 * time.Second,
		ReadTimeout:  5 * time.Second,
	}

	log.Fatal(srv.ListenAndServe())

}
