package main

import (
	"database/sql"
	"flag"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/todzuko/snippetbox-lets-go/internal/models"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger   *slog.Logger
	snippets *models.SnippetModel
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
	}
	dsn := flag.String("dsn", os.Getenv("DSN"), "MySQL data source name")
	db, err := openDb(*dsn)
	if err != nil {
		logger.Error(err.Error())
	}

	defer db.Close()

	app := &application{
		logger:   logger,
		snippets: &models.SnippetModel{DB: db},
	}

	addr := flag.String("addr", os.Getenv("HTTP_ADDR"), "HTTP address")
	flag.Parse()

	logger.Info("Starting server", slog.String("addr", *addr))

	err = http.ListenAndServe(*addr, app.routes())
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}

func openDb(dsn string) (*sql.DB, error) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}
