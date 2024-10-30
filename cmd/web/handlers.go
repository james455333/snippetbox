package main

import (
	"errors"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

type neuteredFileSystem struct {
	fs  http.FileSystem
	app *Application
}

func (app *Application) snippetCreatePost(writer http.ResponseWriter, request *http.Request) {
	title := "O snail"
	content := "O snail\nClimb Mount Fuji,\nBut slowly, slowly!\n\n– Kobayashi Issa"
	expires := 7

	id, err := app.snippets.Insert(title, content, expires)
	if err != nil {
		app.ServerError(writer, request, err)
		return
	}

	http.Redirect(writer, request, fmt.Sprintf("/snippet/view/%d", id), http.StatusSeeOther)
}

func (app *Application) snippetCreateGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display snippet create"))
}

func (app *Application) snippetViewGet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil && errors.Is(err, strconv.ErrSyntax) {
		writer.WriteHeader(http.StatusBadRequest)
		app.Logger.Error(fmt.Sprintf("invalid syntax with : %s", request.PathValue("id")))
		return
	} else if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		app.Logger.Error(fmt.Sprintf("invalid syntax with : %s", request.PathValue("id")))
		return
	}
	fmt.Fprintf(writer, "Display a snippet view ID: %d", id)
}

func (app *Application) HomeGet(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	homeTemplate, err := template.ParseFiles(files...)
	if err != nil {
		app.ServerError(writer, request, err)
		return
	}

	err = homeTemplate.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		app.ServerError(writer, request, err)
		return
	}
}

func (app *Application) staticSubTreeGet(url, staticSrcRootPath string) http.Handler {
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(staticSrcRootPath), app})
	return http.StripPrefix(url, fileServer)
}

func (neuteredFileSystem neuteredFileSystem) Open(path string) (http.File, error) {
	neuteredFileSystem.app.Logger.Info(fmt.Sprintf("neuteredFileSystem Open start on : %s", path))
	file, err := neuteredFileSystem.fs.Open(path)
	if err != nil {
		neuteredFileSystem.app.Logger.Error(err.Error())
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		neuteredFileSystem.app.Logger.Error(err.Error())
		return nil, err
	}

	if stat.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := neuteredFileSystem.fs.Open(index); err != nil {
			closeErr := file.Close()
			if closeErr != nil {
				neuteredFileSystem.app.Logger.Error(closeErr.Error())
				return nil, closeErr
			}
			neuteredFileSystem.app.Logger.Error(err.Error())
			return nil, err
		}
	}

	neuteredFileSystem.app.Logger.Info(fmt.Sprintf("find file: %s", stat.Name()))
	return file, nil
}
