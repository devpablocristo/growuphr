package api

import (
	"github.com/go-chi/chi"

	chihandler "github.com/devpablocristo/growuphr/reserve-number/infrastructure/driver-adapter/handler/chi"
)

func ChiRouter(handler *chihandler.ChiHandler) *chi.Mux {
	router := chi.NewRouter()
	//chiMux.Use("cors")
	//chiMux.Use(middleware.Logger)

	router.Route("/api/v1", func(r chi.Router) {
		r.Route("/number", func(r chi.Router) {
			r.Post("/add-number", handler.AddNumber)
		})
	})

	return router
}
