package endpoints

import (
	"net/http"

	"github.com/DarknessRdg/rinha-backend-2024-q1/internal/errs"
)

func writeError(err error, response http.ResponseWriter) {
	apiErr, ok := err.(errs.ApiError)

	if ok {
		response.WriteHeader(int(apiErr.StatusCode))
	}

	response.Write([]byte(err.Error()))
}
