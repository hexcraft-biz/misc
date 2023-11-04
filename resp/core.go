package resp

import (
	"errors"
	"net/http"
)

type Err struct {
	Message *string
}

func (e Err) Error() string {
	return *e.Message
}

func (e Err) Is(target error) bool {
	return e == target
}

// ================================================================
//
// ================================================================
type Payload struct {
	Message string `json:"message"`
	Result  any    `json:"result,omitempty"`
}

type Resp struct {
	StatusCode int
	*Payload
	*Err
}

// ================================================================
func New(code int, result any) *Resp {
	return &Resp{
		StatusCode: code,
		Payload: &Payload{
			Message: http.StatusText(code),
			Result:  result,
		},
		Err: nil,
	}
}

// ================================================================
func NewError(code int, err error, result any) *Resp {
	var resp *Resp

	if err != nil {
		resp = newError(code, err.Error(), result)
	}

	return resp
}

// ================================================================
func newError(code int, msg string, result any) *Resp {
	resp := &Resp{
		StatusCode: code,
		Payload: &Payload{
			Message: msg,
			Result:  result,
		},
	}

	resp.Err = &Err{Message: &resp.Payload.Message}
	return resp
}

var (
	ErrBadRequest          = newError(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
	ErrUnauthorized        = newError(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
	ErrForbidden           = newError(http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
	ErrNotFound            = newError(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
	ErrConflict            = newError(http.StatusConflict, http.StatusText(http.StatusConflict), nil)
	ErrGone                = newError(http.StatusGone, http.StatusText(http.StatusGone), nil)
	ErrUnprocessableEntity = newError(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil)
	ErrServiceUnavailable  = newError(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable), nil)
	ErrInternalServerError = newError(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)

	Errs = errors.Join(
		ErrBadRequest,
		ErrUnauthorized,
		ErrForbidden,
		ErrNotFound,
		ErrConflict,
		ErrGone,
		ErrUnprocessableEntity,
		ErrServiceUnavailable,
		ErrInternalServerError,
	)
)

// ================================================================
//
// ================================================================
func (r Resp) O() (int, *Payload) {
	if r.StatusCode == http.StatusNoContent {
		return r.StatusCode, nil
	} else {
		return r.StatusCode, r.Payload
	}
}
