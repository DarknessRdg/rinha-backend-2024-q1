package endpoints

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
)

func writeError(err error, response http.ResponseWriter) {
	apiErr, ok := err.(errs.ApiError)

	if ok {
		response.WriteHeader(int(apiErr.StatusCode))
	} else {
		response.WriteHeader(http.StatusInternalServerError)
	}

	errDetail := struct {
		Detail string `json:"detalhe"`
	}{
		Detail: err.Error(),
	}

	errMessage, err := json.Marshal(errDetail)
	if err != nil {
		slog.Error("unable to marshal err detail json", "err", err.Error())
		return
	}
	response.Write(errMessage)
}
