package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", app.staticSubTreeGet("/static", app.Cfg.staticSrcDir))

	mux.HandleFunc("GET /{$}", app.HomeGet)
	mux.HandleFunc("GET /snippet/view/{id}", app.snippetViewGet)
	mux.HandleFunc("GET /snippet/create", app.snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", app.snippetCreatePost)
	return mux
}
