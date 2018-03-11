package api

import (
	"github.com/go-chi/chi"
	"net/http"
	"context"
	"time"
	"github.com/go-chi/chi/middleware"
)

var routes = chi.NewRouter()
var server *http.Server

func StartServer() error {
	routes.Use(middleware.Recoverer)
	routes.Mount("/", defaultRoutes())

	server = &http.Server{
		Addr:    ":9900",
		Handler: routes,
	}
	err := server.ListenAndServe()
	if err != nil {
		return err
	}
	return nil
}

func defaultRoutes() http.Handler {
	h := chi.NewRouter()
	h.Get("/start", startStream)
	h.Get("/stop", stopStream)
	h.Handle("/cam", serveStream())
	return h
}

func StopServer() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
}
