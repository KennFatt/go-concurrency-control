package handler

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

func Bind[T any](requestBody io.ReadCloser, out *T) error {
	decoder := json.NewDecoder(requestBody)
	return decoder.Decode(out)
}

func HandleResponse(w http.ResponseWriter, r *http.Request, err error, data any) {
	w.Header().Set("Content-Type", "application/json")
	jsonEncoder := json.NewEncoder(w)

	if err == nil {
		res := NewResponseSuccessData(data)

		w.WriteHeader(http.StatusOK)
		_ = jsonEncoder.Encode(res)
		slog.Info("request finished",
			"method", r.Method,
			"requestURI", r.RequestURI,
			"requestHeader", r.Header,
		)
		return
	}

	// NOTE: we can map the error (e.g. internal error, framework error, etc)
	// for the simplicity, we can just go this way.
	res := NewResponseError(err)

	w.WriteHeader(http.StatusInternalServerError)
	_ = jsonEncoder.Encode(res)

	slog.Error("request error",
		"error", err,
		"method", r.Method,
		"requestURI", r.RequestURI,
		"requestHeader", r.Header,
	)
}
