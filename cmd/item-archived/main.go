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

	connectcors "connectrpc.com/cors"
	"github.com/lmittmann/tint"
	"github.com/rs/cors"
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

func withCORS(connectHandler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"http://localhost:5173"}, // replace with your domain
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
		MaxAge:         7200, // 2 hours in seconds
	})
	return c.Handler(connectHandler)
}

func main() {
	reldir := flag.String("dir", ".", "The item archive directory to serve.")
	verbose := flag.Bool("v", false, "Enable verbose logging.")
	flag.Parse()

	logLevel := slog.LevelInfo
	if *verbose {
		logLevel = slog.LevelDebug
	}
	slog.SetDefault(slog.New(
		tint.NewHandler(os.Stderr, &tint.Options{
			Level: logLevel,
		}),
	))

	wd, err := os.Getwd()
	if err != nil {
		slog.Error("could not get current working directory", "err", err)
		os.Exit(1)
	}
	dir := *reldir
	if !filepath.IsAbs(*reldir) {
		dir = filepath.Join(wd, *reldir)
	}

	slog.Info("item archive directory", "dir", dir)

	service, err := service.NewService(dir)
	if err != nil {
		slog.Error("failed to create service", "err", err)
		os.Exit(1)
	}
	mux := http.NewServeMux()

	path, connecthandler := v1connect.NewArchiveServiceHandler(service)
	mux.Handle(path, withCORS(connecthandler))

	// mux.Handle("/", http.StripPrefix("/", http.FileServer(http.FS(web))))

	startHttpServer(8330, mux)
}
