package gin

import (
	"errors"
	"net/http"
)

func diffloreDepthBodyLimitStatusC(err error) int {
	var maxBytesError *http.MaxBytesError
	if errors.As(err, &maxBytesError) {
		return http.StatusBadRequest
	}
	return http.StatusBadRequest
}
