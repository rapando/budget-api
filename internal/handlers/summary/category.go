package summary

import (
	"github.com/go-chi/chi"
	"github.com/rapando/budget-api/internal/entities"
	"github.com/rapando/budget-api/internal/middleware"
	"net/http"
)

func (h handler) CategorySummary(w http.ResponseWriter, r *http.Request) {
	var summary entities.CategorySummary
	var err error
	var period = chi.URLParam(r, "period")

	summary, err = h.service.CategorySummary(period)
	if err != nil {
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}
	middleware.Response(w, http.StatusOK, summary)
}
