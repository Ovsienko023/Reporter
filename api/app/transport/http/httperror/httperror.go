package httperror

import (
	"encoding/json"
	"net/http"
)

const (
	InvalidRequest = "invalid request"
)

type ErrorResponseDetails struct {
	Reason      string `json:"reason,omitempty"`
	Description string `json:"description,omitempty"`
	Position    string `json:"position,omitempty"`
}

type ErrorResponseError struct {
	Code        int                    `json:"code,omitempty"`
	Description string                 `json:"description,omitempty"`
	Details     []ErrorResponseDetails `json:"details"`
}

type ErrorResponse struct {
	Error ErrorResponseError `json:"error,omitempty"`
}

func (r *ErrorResponse) Add(reason, description, position string) {
	details := ErrorResponseDetails{
		Reason:      reason,
		Description: description,
		Position:    position,
	}
	r.Error.Details = append(r.Error.Details, details)
}

func (r *ErrorResponse) Done(w http.ResponseWriter, code int, description string) {
	r.Error.Code = code
	r.Error.Description = description
	if len(r.Error.Details) == 0 {
		r.Error.Details = []ErrorResponseDetails{}
	}

	response, _ := json.Marshal(r)

	w.WriteHeader(code)
	w.Header().Add("Content-Type", "application/json")
	_, _ = w.Write([]byte(response))

}

func (r *ErrorResponse) Marshaling() ([]byte, error) {
	if len(r.Error.Details) == 0 {
		r.Error.Details = []ErrorResponseDetails{}
	}

	response, err := json.Marshal(r)
	if err != nil {
		return nil, err
	}

	return response, nil
}
