package budget

import (
	"encoding/json"
	"github.com/rapando/budget-api/internal/entities"
	"github.com/rapando/budget-api/internal/middleware"
	"github.com/rapando/budget-api/pkg/log"
	"io"
	"net/http"
)

func (h handler) AddTransaction(w http.ResponseWriter, r *http.Request) {
	var data entities.AddTransactionRequest
	var err error
	var dataBytes []byte

	dataBytes, err = io.ReadAll(r.Body)
	if err != nil {
		log.Warnf("failed to read add-transaction data: %v", err)
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}
	log.Debugf("add-transaction payload: %+v", string(dataBytes))

	err = json.Unmarshal(dataBytes, &data)
	if err != nil {
		log.Warnf("failed to unmarshal add-transaction data because %v", err)
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}

	err = data.Validate()
	if err != nil {
		log.Warnf("error in add-transaction payload: %v", err)
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}
	err = h.service.AddTransaction(data.ToTransaction())
	if err != nil {
		middleware.Response(w, http.StatusBadRequest, middleware.ErrorResponse)
		return
	}

	middleware.Response(w, http.StatusCreated, map[string]string{"message": "ok"})
}
