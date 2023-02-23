package api

import (
	"github.com/go-chi/chi"

	handler "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driver-adapter/handler"
)

func Router(handler *handler.Handler) *chi.Mux {
	router := chi.NewRouter()
	//chiMux.Use("cors")
	//chiMux.Use(middleware.Logger)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/number-service", func(r chi.Router) {
			r.Post("/reserve/{username}", handler.ReserveNumber)
			r.Post("/reserved-numbers", handler.ReservedNumbers)
		})
	})

	return router
}
