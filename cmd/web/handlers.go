package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"path/filepath"
	"strconv"
	"text/template"
)

type neuteredFileSystem struct {
	fs http.FileSystem
}

func snippetCreatePost(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Content-Type", "application/json; charset=utf-8")
	writer.WriteHeader(http.StatusCreated)
	writer.Write([]byte(`{"name" : "Alex"}`))
}

func snippetCreateGet(writer http.ResponseWriter, request *http.Request) {
	writer.Write([]byte("Display snippet create"))
}

func snippetViewGet(writer http.ResponseWriter, request *http.Request) {
	id, err := strconv.Atoi(request.PathValue("id"))
	if err != nil && errors.Is(err, strconv.ErrSyntax) {
		writer.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid syntax with : %s", request.PathValue("id"))
		return
	} else if err != nil {
		writer.WriteHeader(http.StatusBadRequest)
		log.Printf("invalid syntax with : %s", request.PathValue("id"))
		return
	}
	msg := fmt.Sprintf("Display a snippet view ID: %d", id)
	fmt.Fprintf(writer, msg)
}

func HomeGet(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Add("Server", "Go")

	files := []string{
		"./ui/html/partials/nav.tmpl.html",
		"./ui/html/pages/base.tmpl.html",
		"./ui/html/pages/home.tmpl.html",
	}

	homeTemplate, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Sever Error", http.StatusInternalServerError)
	}

	err = homeTemplate.ExecuteTemplate(writer, "base", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(writer, "Internal Sever Error", http.StatusInternalServerError)
	}
}

func staticSubTreeGet(url, staticSrcRootPath string) http.Handler {
	fileServer := http.FileServer(neuteredFileSystem{http.Dir(staticSrcRootPath)})
	return http.StripPrefix(url, fileServer)
}

func (neuteredFileSystem neuteredFileSystem) Open(path string) (http.File, error) {
	log.Printf("neuteredFileSystem Open start on : %s", path)
	file, err := neuteredFileSystem.fs.Open(path)
	if err != nil {
		log.Println("error :", err.Error())
		return nil, err
	}

	stat, err := file.Stat()
	if err != nil {
		log.Println("error", err.Error())
		return nil, err
	}

	if stat.IsDir() {
		index := filepath.Join(path, "index.html")
		if _, err := neuteredFileSystem.fs.Open(index); err != nil {
			closeErr := file.Close()
			if closeErr != nil {
				log.Println("error :", closeErr.Error())
				return nil, closeErr
			}
			log.Println("error :", err.Error())
			return nil, err
		}
	}

	log.Println("find file:", stat.Name())
	return file, nil
}
