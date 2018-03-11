package api

import (
	"github.com/go-chi/chi"
	"net/http"
	"context"
	"time"
)

var routes = chi.NewRouter()
var server *http.Server

func StartServer() error {
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
	h.Group(func(r chi.Router) {
		r.Use(recovery)

		r.Get("/start", startStream)
		r.Get("/stop", stopStream)
		r.Handle("/cam", serveStream())
	})
	return h
}

func StopServer() {
	ctx, _ := context.WithTimeout(context.Background(), 5*time.Second)
	server.Shutdown(ctx)
}
