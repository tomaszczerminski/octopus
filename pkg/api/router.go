package api

import (
	"github.com/go-chi/chi"
	"octopus/pkg/api/endpoints"
)
import "github.com/go-chi/chi/middleware"

func Router() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/games", endpoints.Games())
	return r
}
