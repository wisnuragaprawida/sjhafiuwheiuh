package api

import (
	"net/http"

	"github.com/go-chi/chi"
	mdl "github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/jmoiron/sqlx"
	"github.com/wisnuragaprawida/project/internal/api/auth"
	"github.com/wisnuragaprawida/project/internal/api/response"
)

type ServerConfig struct {
	EncKey string
}

func tryHandler(w http.ResponseWriter, r *http.Request) {
	response.Yay(w, r, "caca cacing", http.StatusOK)
}

func NewServer(db *sqlx.DB, cfg ServerConfig) *chi.Mux {

	var (
		r = chi.NewRouter()

		authHandler = auth.NewAuthHandler(db)
	)

	r.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		Debug:            false}).Handler)

	r.Use(mdl.RequestID)
	r.Use(mdl.RealIP)
	r.Use(mdl.Logger)
	r.Use(mdl.Recoverer)
	r.Use(mdl.Heartbeat("/ping"))
	r.Use(mdl.Recoverer)

	r.Route("/api", func(r chi.Router) {
		r.Post("/v1/register", authHandler.Register)
		r.Post("/v1/login", authHandler.Login)
	})

	r.Route("/v1", func(r chi.Router) {
		//for auth
		// r.Use(midleware.Auth(cfg.EncKey))
		r.Get("/test", tryHandler)
	})

	return r
}
