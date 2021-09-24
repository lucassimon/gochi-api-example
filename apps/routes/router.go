package routes

import (
	"errors"
	"gochi/apps/middlewares"
	"gochi/apps/responses"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"go.uber.org/zap"
)

func CreateRoutes(log *zap.Logger) http.Handler {

	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.URLFormat)
	r.Use(middlewares.ContentTypeJson)

	if log != nil {
		r.Use(middlewares.SetLogger(log))
	}

	r.Get("/ping", func(w http.ResponseWriter, r *http.Request) {
		errInvalid := responses.ErrInvalidRequest(errors.New("teste"))
		responses.RespondWithError(w, errInvalid)
	})

	// Mount the admin sub-router, which btw is the same as:
	// r.Route("/admin", func(r chi.Router) { admin routes here })
	r.Mount("/admin", adminRouter(log))

	return r
}
