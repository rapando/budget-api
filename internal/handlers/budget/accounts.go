package budget

import (
	"github.com/rapando/budget-api/internal/entities"
	"github.com/rapando/budget-api/internal/middleware"
	"net/http"
)

func (h handler) ListAccounts(w http.ResponseWriter, r *http.Request) {
	var accounts []entities.Account
	var err error

	accounts, err = h.service.ListAccounts()
	if err != nil {
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}
	middleware.Response(w, http.StatusOK, accounts)
}
