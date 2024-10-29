package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	StaticSrcRootPath string = "./ui/static/"
)

type config struct {
	addr         string
	staticSrcDir string
	flag1        bool
}

func main() {
	var cfg config
	//addr := flag.String("addr", ":4000", "HTTP network address")
	//addr := os.Getenv("SNIPPETBOX_ADDR")
	flag.StringVar(&cfg.addr, "addr", ":4000", "HTTP network address")
	flag.StringVar(&cfg.staticSrcDir, "staticSrcDir", StaticSrcRootPath, "Path to static assets")
	//boolFlag1 := flag.Bool("flag1", true, "test boolean flag")
	flag.BoolVar(&cfg.flag1, "flag1", true, "test boolean flag")
	flag.Parse()

	//fmt.Printf("boolFlag1 : %b\n", *boolFlag1)
	fmt.Printf("boolFlag1 : %b\n", &cfg.flag1)

	mux := http.NewServeMux()

	mux.Handle("GET /static/", staticSubTreeGet("/static", cfg.staticSrcDir))

	mux.HandleFunc("GET /{$}", HomeGet)
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	mux.HandleFunc("GET /snippet/create", snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	log.Print("starting server on ", cfg.addr)
	log.Print(cfg)
	err := http.ListenAndServe(cfg.addr, mux)
	log.Fatal(err)
}
