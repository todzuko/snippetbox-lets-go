package main

import (
	"flag"
	"github.com/joho/godotenv"
	"log/slog"
	"net/http"
	"os"
)

type application struct {
	logger *slog.Logger
}

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		AddSource: true,
	}))

	app := &application{
		logger: logger,
	}
	err := godotenv.Load(".env")
	if err != nil {
		logger.Error(err.Error())
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
