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

	//Blocking call
	startServer()

}

func startServer() {

	r := mux.NewRouter()
	r.HandleFunc("/api/term-search", api.HandleTermSearch)
	r.HandleFunc("/api/word-lookup", api.HandleWordLookup)
	r.HandleFunc("/api/chapter-text", api.HandleChapterText)
	r.HandleFunc("/api/node-text", api.HandleNodeText)
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
