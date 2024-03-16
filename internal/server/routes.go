package server

import (
	"net/http"

	"emil/cmd/web"
	"emil/internal/static"
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", web.HomeHandler)
	fileServer := http.FileServer(http.FS(web.Files))
	r.Handle("/styles/*", http.FileServer(http.FS(static.StyleFiles)))
	r.Handle("/media/*", http.FileServer(http.FS(static.ImageFiles)))
	r.Handle("/js/*", fileServer)
	r.Get("/web", templ.Handler(web.HelloForm()).ServeHTTP)
	r.Post("/hello", web.HelloWebHandler)
	r.Get("/commands", web.CommandsHandler)

	return r
}
