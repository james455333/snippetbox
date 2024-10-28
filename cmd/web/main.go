package main

import (
	"log"
	"net/http"
)

var (
	StaticSrcRootPath string = "./ui/static/"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", staticSubTreeGet("/static", StaticSrcRootPath))

	mux.HandleFunc("GET /{$}", HomeGet)
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	mux.HandleFunc("GET /snippet/create", snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on :4000")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}
