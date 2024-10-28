package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"text/template"
)

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
	fileServer := http.FileServer(http.Dir(staticSrcRootPath))
	return http.StripPrefix(url, fileServer)
}
