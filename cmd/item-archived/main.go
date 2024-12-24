package main

import (
	"flag"
	"fmt"
	"item-archived/api/v1/v1connect"
	"item-archived/internal/service"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"

	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

// //go:embed web/dist/*
// var web embed.FS

func startHttpServer(port int, mux *http.ServeMux) {
	slog.Info("listening to gRPC...", "port", port)
	err := http.ListenAndServe(
		fmt.Sprintf("0.0.0.0:%d", port),
		h2c.NewHandler(mux, &http2.Server{}),
	)
	if err != nil {
		slog.Error("failed to listen to port", "port", port, "err", err)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	rel := flag.Arg(0)
	if rel == "" {
		rel = "."
	}
	wd, err := os.Getwd()
	if err != nil {
		slog.Error("could not get current working directory", "err", err)
		os.Exit(1)
	}
	dir := filepath.Join(wd, rel)

	service := service.NewService(dir)
	mux := http.NewServeMux()
	mux.Handle(v1connect.NewArchiveServiceHandler(service))
	// mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(web))))

	startHttpServer(8330, mux)
}
