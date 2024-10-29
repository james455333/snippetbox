package main

import "net/http"

// The routes() method returns a servemux containing our application routes.
func (app *Application) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.Handle("GET /static/", staticSubTreeGet("/static", app.Cfg.staticSrcDir, app))

	mux.HandleFunc("GET /{$}", HomeGet(app))
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewGet(app))
	mux.HandleFunc("GET /snippet/create", snippetCreateGet(app))
	mux.HandleFunc("POST /snippet/create", snippetCreatePost(app))
	return mux
}
