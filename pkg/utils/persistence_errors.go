package utils

import "errors"

var (
	ErrorInternal         = errors.New("internal server error")
	ErrorBadRequest       = errors.New("bad request")
	ErrorResourceNotFound = errors.New("resource not found")
	ErrorResourceConflict = errors.New("resource conflict")
)

func HandleWebStatusErrors(err error) int {
	status := 200
	if err != nil {
		if err == ErrorResourceNotFound {
			status = 404
		}
		if err == ErrorBadRequest {
			status = 400
		}
		if err == ErrorResourceConflict {
			status = 409
		}

		if err == ErrorInternal {
			status = 500
		}
	}

	return status
}
