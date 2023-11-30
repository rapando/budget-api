package summary

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
	h.router.Route("/summary", func(r chi.Router) {
		r.Use(middleware.AuthMW)

		r.Get("/{period:(day|week|month|year)}", h.Summary)
		r.Get("/transactions/{period:(day|week|month|year)}", h.TransactionsSummary)
		r.Get("/account/{period:(day|week|month|year)}", h.AccountSummary)
		r.Get("/category/{period:(day|week|month|year)}", h.CategorySummary)

	})
}
