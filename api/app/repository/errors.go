package repository

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"
)

var (
	ErrInternal            = errors.New("internal error")
	ErrUnexpectedBehavior  = errors.New("unexpected behavior")
	ErrUnauthorized        = errors.New("unauthorized")
	ErrForbidden           = errors.New("permission denied")
	ErrReportIdNotFound    = errors.New("report id not found")
	ErrSickLeaveIdNotFound = errors.New("sick leave id not found")
	ErrCredentials         = errors.New("error credentials")
	ErrLoginAlreadyInUse   = errors.New("login already in use")
)

func NewInternalError(err error) error {
	if err == nil {
		return nil
	}
	ErrInternal := errors.New(err.Error())
	return ErrInternal
}

type queryErrorDetail struct {
	Name   string `json:"name,omitempty"`
	Type   string `json:"type,omitempty"`
	Reason string `json:"reason,omitempty"`
}

type queryError struct {
	Code    int                `json:"code,omitempty"`
	Details []queryErrorDetail `json:"details,omitempty"`
	Message string             `json:"message,omitempty"`
}

// AnalyzeError выделяет сообщение ошибки из бд и возвращает соответствующий её error.
func AnalyzeError(errorJson []byte) error {
	if len(errorJson) == 0 {
		return fmt.Errorf("%w: empty query error", ErrUnexpectedBehavior)
	}

	data := queryError{}
	err := json.Unmarshal(errorJson, &data)

	if err != nil {
		return fmt.Errorf("%w: %v, json:%v", ErrUnexpectedBehavior, err, string(errorJson))
	}

	switch data.Code {
	case 1:
		return ErrUnauthorized
	case 2:
		return ErrForbidden
	case 3:
		for _, obj := range data.Details {

			if obj.Name == "_login" && obj.Reason == "exists" {
				return ErrLoginAlreadyInUse
			}

		}
	}
	return fmt.Errorf("%w: errorCodeUnknown json:%v", ErrUnexpectedBehavior, string(errorJson))
}

func AnalyzeRowsError(err error) error {
	msg := err.Error()
	// Выделяем сообщение
	msg = strings.TrimPrefix(msg, "ERROR: ")
	messages := strings.Split(msg, " (SQLSTATE")

	if len(messages) > 1 {
		msg = messages[0]

		if strings.Contains(msg, "invalid input") {
			values := strings.SplitN(msg, ": ", 2)

			if len(values) > 1 {
				return errors.New(strings.TrimSpace(values[1]))
			}
		}
	}
	return fmt.Errorf("%w: %s", ErrUnexpectedBehavior, msg)
}
