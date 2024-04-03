package handlers

import (
	"errors"
	"github.com/Ovsienko023/reporter/app/core"
	"github.com/Ovsienko023/reporter/app/domain"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/labstack/echo/v4"
	"net/http"
)

// --------------------------------
//      		AUTH
// --------------------------------

func Auth(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)

	stateId := r.URL.Query().Get("state")

	authParams, err := c.Auth(r.Context(), domain.AuthRequest{
		AuthorizationCode: r.URL.Query().Get("code"),
		StateId:           stateId,
	})

	if err != nil {
		state, ok := c.Cache.AuthState[stateId]
		if !ok {
			HtmlPostMsgErrResponse(w, http.StatusInternalServerError, "todo", "*")
			return
		}

		switch {
		case errors.Is(err, core.ErrUnauthorized):
			HtmlPostMsgErrResponse(w, http.StatusUnauthorized, err.Error(), state.ClientOrigin)
			return
		case errors.Is(err, core.ErrInvalidArguments):
			HtmlPostMsgErrResponse(w, http.StatusBadRequest, err.Error(), state.ClientOrigin)
			return
		default:
			HtmlPostMsgErrResponse(w, http.StatusInternalServerError, core.ErrInternal.Error(), state.ClientOrigin)
			return
		}
	}

	// todo передалать с использованием html/template
	HtmlPostMsgResponse(w, authParams.ClientOrigin, &AuthResponse{Token: authParams.Token})
}

type AuthResponse struct {
	Token string `json:"token,omitempty"`
}

// --------------------------------
//         GET PROVIDER URI
// --------------------------------

type GetProviderUriRequest struct {
	Host string `json:"host,omitempty"`
}

func GetProviderUri(c *core.Core, w http.ResponseWriter, r *http.Request) {
	w.Header().Add(echo.HeaderContentType, echo.MIMEApplicationJSON)
	errorContainer := httperror.ErrorResponse{}

	res, err := c.GetProviderUri(r.Context(),
		&domain.GetProviderUriRequest{ClientOrigin: r.Header.Get("Origin")},
	)

	if err != nil {
		switch {
		case errors.Is(err, core.ErrUnauthorized):
			errorContainer.Done(w, http.StatusUnauthorized, err.Error())
			return
		case errors.Is(err, core.ErrInvalidArguments):
			errorContainer.Done(w, http.StatusBadRequest, err.Error())
			return
		case errors.Is(err, core.ErrHostNotFound):
			errorContainer.Done(w, http.StatusNotFound, err.Error())
			return
		case errors.Is(err, core.ErrProviderServerNotAvailable):
			errorContainer.Done(w, http.StatusBadGateway, err.Error())
			return
		default:
			errorContainer.Done(w, http.StatusInternalServerError, core.ErrInternal.Error())
			return
		}
	}

	result := GetProviderUriResponse{
		Uri: res.Uri,
	}

	JsonResponse(w, http.StatusOK, result)
}

type GetProviderUriResponse struct {
	Uri string `json:"uri,omitempty"`
}
