package domain

import (
	"context"
	"go-cc/domain/seats"
	"go-cc/ent"
	"net/http"
)

func InitRoutes(mux *http.ServeMux, db *ent.Client) {
	ctx := context.Background()

	seats.RegisterHandlers(ctx, mux, db)
}
