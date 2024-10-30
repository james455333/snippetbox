package main

import (
	"database/sql"
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"snippetbox.james455333.github.com/internal/models"

	_ "github.com/lib/pq"
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

	db, err := openDB("web", "your_password", "snippetbox")
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer db.Close()

	cfg := InitConfig()
	app := &Application{
		Logger: logger,
		Cfg:    cfg,
		snippets: &models.SnippetModel{
			DB: db,
		},
	}

	logger.Info("starting server on ", slog.String("addr", cfg.addr))
	logger.Info(fmt.Sprint(cfg))
	err = http.ListenAndServe(cfg.addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDB(userName, pwd, dbName string) (*sql.DB, error) {
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable port=3301", userName, pwd, dbName)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}
