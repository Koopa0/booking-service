package web

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/koopa0/booking-service/pkg/config"
	"github.com/koopa0/booking-service/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := chi.NewRouter()

	mux.Use(middleware.Recoverer)
	mux.Use(NoSurf)
	mux.Use(SessionLoad)

	mux.Get("/", handlers.Repo.Home)
	mux.Get("/about", handlers.Repo.About)

	return mux
}
