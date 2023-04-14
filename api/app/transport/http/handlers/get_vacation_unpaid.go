package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/go-chi/chi/v5"
	"net/http"
)

func GetVacationUnpaid(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
	errorContainer := httperror.ErrorResponse{}

	message := &domain.GetVacationUnpaidRequest{
		Token:            r.Header.Get("Authorization"),
		VacationUnpaidId: chi.URLParam(r, "vacation_unpaid_id"),
	}

	result, err := c.GetVacationUnpaid(r.Context(), message)

	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		case errors.Is(err, core.ErrVacationIdNotFound):
			errorContainer.Done(w, http.StatusNotFound, err.Error())
			return
		case errors.Is(err, core.ErrPermissionDenied):
			errorContainer.Done(w, http.StatusForbidden, err.Error())
			return

		}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	JsonResponse(w, http.StatusOK, result)
}
