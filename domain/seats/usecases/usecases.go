package usecases

import "go-cc/ent"

type UseCases struct {
	db *ent.Client
}

func New(db *ent.Client) *UseCases {
	return &UseCases{db}
}
