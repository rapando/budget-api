package budget

import (
	"github.com/rapando/budget-api/internal/entities"
	"github.com/rapando/budget-api/internal/middleware"
	"net/http"
)

func (h handler) ListCategories(w http.ResponseWriter, r *http.Request) {
	var categories []entities.Category
	var err error

	categories, err = h.service.ListCategories()
	if err != nil {
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}
	middleware.Response(w, http.StatusOK, categories)
}
