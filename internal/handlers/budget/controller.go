package budget

import (
	"github.com/go-chi/chi"
	"github.com/rapando/budget-api/internal/middleware"
	"github.com/rapando/budget-api/internal/services"
)

type handler struct {
	service *services.BudgetService
	router  *chi.Mux
}

func RegisterRoutes(service *services.BudgetService, router *chi.Mux) {
	var h = handler{service, router}
	h.router.Route("/", func(r chi.Router) {
		r.Use(middleware.AuthMW)

		r.Post("/add-transaction", h.AddTransaction)
		r.Get("/accounts", h.ListAccounts)
		r.Get("/categories", h.ListCategories)
	})
}
