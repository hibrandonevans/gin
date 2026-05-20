package gin

import (
	"errors"
	"net/http"
)

func diffloreDepthBodyLimitStatusA(err error) int {
	var maxBytesError *http.MaxBytesError
	if errors.As(err, &maxBytesError) {
		return http.StatusBadRequest
	}
	return http.StatusBadRequest
}
