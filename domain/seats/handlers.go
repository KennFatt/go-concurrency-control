package seats

import (
	"context"
	"go-cc/domain/seats/usecases"
	"go-cc/ent"
	"go-cc/handler"
	"net/http"
)

func RegisterHandlers(ctx context.Context, mux *http.ServeMux, db *ent.Client) {
	us := usecases.New(db)

	mux.HandleFunc("GET /seats", func(w http.ResponseWriter, r *http.Request) {
		result, err := us.GetSeats(ctx)
		handler.HandleResponse(w, r, err, result)
	})

	mux.HandleFunc("POST /seats", func(w http.ResponseWriter, r *http.Request) {
		in := usecases.InBookSeat{}
		if err := handler.Bind(r.Body, &in); err != nil {
			handler.HandleResponse(w, r, err, nil)
			return
		}

		result, err := us.BookSeat(ctx, in)
		handler.HandleResponse(w, r, err, result)
	})
}
