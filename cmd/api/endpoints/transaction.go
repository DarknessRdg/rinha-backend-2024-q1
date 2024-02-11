package endpoints

import (
	"net/http"

	"github.com/DarknessRdg/rinha-backend-2024-q1/cmd/api/decoder"
	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/transaction/dto"
)

func PostTransaction(w http.ResponseWriter, request *http.Request) {
	decoder := decoder.NewJsonDecoder[dto.TransactionDto]()

	_, err := decoder.Decode(request.Body)
	if err != nil {
		panic(err)
	}
}
