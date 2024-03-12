package usecases

import (
	"context"

	"github.com/mitchellh/mapstructure"
)

type (
	OutGetSeats struct {
		ID            int     `json:"id"`
		IsBooked      bool    `json:"isBooked"`
		PassengerName *string `json:"passengerName"`
	}
)

func (us *UseCases) GetSeats(ctx context.Context) ([]*OutGetSeats, error) {
	res, err := us.db.Seat.Query().All(ctx)
	if err != nil {
		return nil, err
	}

	output := make([]*OutGetSeats, 0, len(res))
	err = mapstructure.Decode(res, &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}
