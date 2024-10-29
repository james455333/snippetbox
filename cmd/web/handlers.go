package main

import (
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"path/filepath"
	appConfig "snippetbox.james455333.github.com/cmd/web/config"
	"strconv"
	"text/template"
)

type neuteredFileSystem struct {
	fs  http.FileSystem
	app *appConfig.Application
}

func snippetCreatePost(app *appConfig.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "Application/json; charset=utf-8")
		writer.WriteHeader(http.StatusCreated)
		writer.Write([]byte(`{"name" : "Alex"}`))
	}
}

func snippetCreateGet(app *appConfig.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("Display snippet create"))
	}
}

func snippetViewGet(app *appConfig.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
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
}

func HomeGet(app *appConfig.Application) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Server", "Go")

		files := []string{
			"./ui/html/partials/nav.tmpl.html",
			"./ui/html/pages/base.tmpl.html",
			"./ui/html/pages/home.tmpl.html",
		}

		homeTemplate, err := template.ParseFiles(files...)
		if err != nil {
			app.Logger.Error(err.Error(), slog.String("method", request.Method), slog.String("uri", request.URL.RequestURI()))
			http.Error(writer, "Internal Sever Error", http.StatusInternalServerError)
			return
		}

		err = homeTemplate.ExecuteTemplate(writer, "base", nil)
		if err != nil {
			app.Logger.Error(err.Error(), slog.String("method", request.Method), slog.String("uri", request.URL.RequestURI()))
			http.Error(writer, "Internal Sever Error", http.StatusInternalServerError)
			return
		}
	}
}

func staticSubTreeGet(url, staticSrcRootPath string, app *appConfig.Application) http.Handler {
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
