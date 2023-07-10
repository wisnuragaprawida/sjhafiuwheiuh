package response

import (
	"context"
	"errors"
	"net/http"

	"github.com/go-chi/render"
	"github.com/wisnuragaprawida/project/pkg/crashy"
	"github.com/wisnuragaprawida/project/pkg/log"
)

type Response struct {
	Status string      `json:"status"`
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
}

type ResponseCallback struct {
	Status string `json:"status"`
}

type ErroResponse struct {
	Status  string `json:"status"`
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func (resp Response) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (resp ErroResponse) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func Yay(rw http.ResponseWriter, r *http.Request, data interface{}, code int) {
	p := Response{
		Status: "success",
		Code:   code,
		Data:   data,
	}

	r = r.WithContext(context.WithValue(r.Context(), render.StatusCtxKey, "ok"))
	if err := render.Render(rw, r, p); err != nil {
		Nay(rw, r, crashy.Wrapc(err, crashy.ErrCodeFormatting), http.StatusInternalServerError)
	}
}

func Nay(rw http.ResponseWriter, r *http.Request, err error, code int) {
	p := ErroResponse{
		Status:  "error",
		Code:    code,
		Message: err.Error(),
	}

	rawError := errors.Unwrap(err)
	if rawError != nil {
		rawError = err
	}

	log.Errorf("error: %s", rawError)

	render.Status(r, code)
	if err := render.Render(rw, r, p); err != nil {
		http.Error(rw, "unexpected error while processing the request", http.StatusInternalServerError)
	}
}

func renderHttpError(rw http.ResponseWriter, r *http.Request, code int, err string) {
	p := ErroResponse{
		Status:  "error",
		Code:    code,
		Message: err,
	}

	render.Status(r, code)
	if err := render.Render(rw, r, p); err != nil {
		http.Error(rw, "unexpected error while processing the request", http.StatusInternalServerError)
	}
}

func ExpiredAccess(rw http.ResponseWriter, r *http.Request) {
	renderHttpError(rw, r, http.StatusUnauthorized, crashy.ErrCodeExpired)
}

func ForbidenAccess(rw http.ResponseWriter, r *http.Request) {
	renderHttpError(rw, r, http.StatusForbidden, crashy.ErrCodeForbidden)
}

func UnauthorizedAccess(rw http.ResponseWriter, r *http.Request) {
	renderHttpError(rw, r, http.StatusUnauthorized, crashy.ErrCodeUnauthorized)
}
