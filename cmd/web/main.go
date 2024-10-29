package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
)

var (
	StaticSrcRootPath string = "./ui/static/"
)

func main() {

	loggerHandler := slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelDebug,
		AddSource: true,
	})
	logger := slog.New(loggerHandler)

	cfg := InitConfig()
	app := &Application{
		Logger: logger,
		Cfg:    cfg,
	}

	logger.Info("starting server on ", slog.String("addr", cfg.addr))
	logger.Info(fmt.Sprint(cfg))
	err := http.ListenAndServe(cfg.addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
