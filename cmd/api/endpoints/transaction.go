package endpoints

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DarknessRdg/rinha-backend-2024-q1/cmd/api/decoder"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/service"
	"github.com/go-chi/chi/v5"
)

type TransactionEndpoint struct {
	TransactionService service.ITransactionService
}

func NewTransactionEndpoint(s service.ITransactionService) *TransactionEndpoint {
	return &TransactionEndpoint{TransactionService: s}
}

func (e *TransactionEndpoint) Router(r chi.Router) chi.Router {
	r.Post("/clientes/{clientId}/transacoes", e.PostTransaction)
	return r
}

func (e *TransactionEndpoint) PostTransaction(w http.ResponseWriter, request *http.Request) {
	decoder := decoder.NewJsonDecoder[dto.TransactionDto]()
	clientId, err := strconv.Atoi(chi.URLParam(request, "clientId"))
	if err != nil {
		writeError(err, w)
		return
	}

	transactionDto, err := decoder.Decode(request.Body)
	if err != nil {
		writeError(err, w)
		return
	}

	result, err := e.TransactionService.PostTransaction(clientId, transactionDto)
	if err != nil {
		writeError(err, w)
		return
	}
	response, err := json.Marshal(result)
	if err != nil {
		writeError(err, w)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(response)
}
