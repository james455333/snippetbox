package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var (
	StaticSrcRootPath string = "./ui/static/"
	logger            slog.Logger
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

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	logger := slog.New(loggerHandler)

	mux := http.NewServeMux()

	mux.Handle("GET /static/", staticSubTreeGet("/static", cfg.staticSrcDir))

	mux.HandleFunc("GET /{$}", HomeGet)
	mux.HandleFunc("GET /snippet/view/{id}", snippetViewGet)
	mux.HandleFunc("GET /snippet/create", snippetCreateGet)
	mux.HandleFunc("POST /snippet/create", snippetCreatePost)

	logger.Info("starting server on ", slog.String("addr", cfg.addr))
	logger.Info(fmt.Sprint(cfg))
	err := http.ListenAndServe(cfg.addr, mux)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
