package resp

import (
	"encoding/json"
	"errors"
	"net/http"
)

type Err struct {
	text *string
}

func (e Err) Error() string {
	if e.text != nil {
		return *e.text
	}
	return ""
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

func NewPayload(result any) *Payload {
	return &Payload{
		Result: result,
	}
}

type Resp struct {
	StatusCode int
	*Payload
	*Err
}

// ================================================================
func New(code int, result any) *Resp {
	r := &Resp{
		StatusCode: code,
		Payload: &Payload{
			Message: http.StatusText(code),
			Result:  result,
		},
	}

	if code >= 400 {
		r.Err = &Err{text: &r.Message}
	}

	return r
}

// Return an *Resp with err passing in. return nil if err is nil.
func NewError(code int, err error, result any) *Resp {
	var resp *Resp

	if err != nil {
		resp = NewErrorWithMessage(code, err.Error(), result)
	}

	return resp
}

// ================================================================
func NewErrorWithMessage(code int, msg string, result any) *Resp {
	resp := &Resp{
		StatusCode: code,
		Payload: &Payload{
			Message: msg,
			Result:  result,
		},
	}

	resp.Err = &Err{text: &resp.Payload.Message}
	return resp
}

var (
	ErrBadRequest          = NewErrorWithMessage(http.StatusBadRequest, http.StatusText(http.StatusBadRequest), nil)
	ErrUnauthorized        = NewErrorWithMessage(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized), nil)
	ErrForbidden           = NewErrorWithMessage(http.StatusForbidden, http.StatusText(http.StatusForbidden), nil)
	ErrNotFound            = NewErrorWithMessage(http.StatusNotFound, http.StatusText(http.StatusNotFound), nil)
	ErrConflict            = NewErrorWithMessage(http.StatusConflict, http.StatusText(http.StatusConflict), nil)
	ErrGone                = NewErrorWithMessage(http.StatusGone, http.StatusText(http.StatusGone), nil)
	ErrUnprocessableEntity = NewErrorWithMessage(http.StatusUnprocessableEntity, http.StatusText(http.StatusUnprocessableEntity), nil)
	ErrServiceUnavailable  = NewErrorWithMessage(http.StatusServiceUnavailable, http.StatusText(http.StatusServiceUnavailable), nil)
	ErrInternalServerError = NewErrorWithMessage(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError), nil)

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

func Assert(err error) *Resp {
	if resp, ok := err.(*Resp); ok {
		return resp
	} else {
		return nil
	}
}

func FetchHexcApiResult(resp *http.Response, payload *Payload) *Resp {
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	if err := decoder.Decode(payload); err != nil {
		return NewError(http.StatusInternalServerError, err, nil)
	}

	if resp.StatusCode >= 500 {
		return NewErrorWithMessage(http.StatusServiceUnavailable, payload.Message, nil)
	}

	return nil
}
