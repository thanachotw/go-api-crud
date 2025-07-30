package response

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	ErrNotFound            = errors.New("resource not found")
	ErrInvalidInput        = errors.New("invalid input provided")
	ErrInternalServer      = errors.New("internal server error")
	ErrUnauthorized        = errors.New("unauthorized access")
	ErrOwnerIDInvalid      = errors.New("owner_id must be positive")
	ErrUnsupportedCurrency = errors.New("unsupported currency")
)

var (
	SuccessMessage = "successfully"
)
var (
	// Errs is a map of error messages to their corresponding error codes
	Errs = map[error]int{
		ErrNotFound:            http.StatusNotFound,
		ErrInvalidInput:        http.StatusBadRequest,
		ErrUnsupportedCurrency: http.StatusBadRequest,
		ErrOwnerIDInvalid:      http.StatusBadRequest,
		ErrInternalServer:      http.StatusInternalServerError,
		ErrUnauthorized:        http.StatusUnauthorized,
	}
)

func HandlerErrorResponse(c *gin.Context, err error) {
	errCode, ok := Errs[err]
	if !ok {
		errCode = http.StatusInternalServerError
	}
	ReponseError(c, errCode, err.Error())
}
