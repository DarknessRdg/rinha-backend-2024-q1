package endpoints

import (
	"net/http"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
)

func writeError(err error, resposne http.ResponseWriter) {
	apiErr, ok := err.(errs.ApiError)

	if ok {
		resposne.WriteHeader(int(apiErr.StatusCode))
	}

	resposne.Write([]byte(err.Error()))
}
