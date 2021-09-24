package routes

import (
	"gochi/apps/middlewares"
	"net/http"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
)

// A completely separate router for administrator routes
func adminRouter(l *zap.Logger) chi.Router {
	r := chi.NewRouter()

	r.Use(middlewares.AdminOnly)
	r.Use(middlewares.ContentTypeJson)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		l.Info("Call my logger",
			// Structured context as strongly typed Field values.
			zap.String("url", "/admin"),
			zap.Int("attempt", 3),
			zap.Duration("backoff", 1),
		)
		w.Write([]byte("admin: index"))
	})
	r.Post("/accounts", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("admin: list accounts.."))
	})

	return r
}
