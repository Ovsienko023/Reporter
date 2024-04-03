package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/Ovsienko023/reporter/app/domain/constants"
	"github.com/Ovsienko023/reporter/app/transport/http/httperror"
	"github.com/labstack/echo/v4"
	"net/http"
)

func JsonResponse(w http.ResponseWriter, code int, resp any) {
	if resp == nil {
		w.WriteHeader(code)
		return
	}

	response, err := json.Marshal(resp)
	if err != nil {
		errorContainer := httperror.ErrorResponse{}
		errorContainer.Done(w, http.StatusInternalServerError, "internal error")
		return
	}

	w.WriteHeader(code)
	_, _ = w.Write(response)
}

func FileResponse(w http.ResponseWriter, file []byte, filename string) error {
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment;filename="+filename)

	_, err := w.Write(file)
	if err != nil {
		return err
	}

	return nil
}

func HtmlPostMsgResponse(w http.ResponseWriter, origin string, resp *AuthResponse) {
	w.Header().Add(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)

	var encodeResp []byte

	encodeResp, err := json.Marshal(resp)
	if err != nil {
		errorContainer := httperror.ErrorResponse{
			Error: httperror.ErrorResponseError{
				Code:        http.StatusInternalServerError,
				Description: "Internal Server Error",
				Details:     nil,
			},
		}
		encodeResp, _ = errorContainer.Marshaling()
	}

	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(constants.HtmlTemplatePostMsg, encodeResp, origin)))
}

func HtmlPostMsgErrResponse(w http.ResponseWriter, code int, description string, origin string) {
	w.Header().Add(echo.HeaderContentType, echo.MIMETextHTMLCharsetUTF8)

	var encodeResp []byte

	errorContainer := httperror.ErrorResponse{
		Error: httperror.ErrorResponseError{
			Code:        code,
			Description: description,
			Details:     nil,
		},
	}

	encodeResp, _ = errorContainer.Marshaling()
	w.WriteHeader(http.StatusOK)
	_, _ = w.Write([]byte(fmt.Sprintf(constants.HtmlTemplatePostMsg, encodeResp, origin)))
}
