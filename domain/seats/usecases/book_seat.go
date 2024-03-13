package usecases

import (
	"context"
	"errors"
	"go-cc/ent/seat"

	"github.com/mitchellh/mapstructure"
)

type (
	InBookSeat struct {
		SeatID        int    `json:"seatId"`
		PassengerName string `json:"passengerName"`
	}

	OutBookSeat struct {
		ID            int     `json:"id"`
		IsBooked      bool    `json:"isBooked"`
		PassengerName *string `json:"passengerName"`
	}
)

var (
	ErrSeatIsBooked                = errors.New("the seat is already booked")
	ErrDataUpdatedByAnotherProcess = errors.New("the seat is updated by another process")
)

func (us *UseCases) BookSeat(ctx context.Context, in InBookSeat) (*OutBookSeat, error) {
	tx, err := us.db.Tx(ctx)
	if err != nil {
		return nil, err
	}

	// Find the requested seat
	requestedSeat, err := tx.Seat.Query().Where(seat.ID(in.SeatID)).Only(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if requestedSeat.IsBooked {
		tx.Rollback()
		return nil, ErrSeatIsBooked
	}

	// Book a seat
	n, err := tx.Seat.
		Update().
		Where(seat.ID(in.SeatID), seat.Version(requestedSeat.Version)).
		SetIsBooked(true).
		SetPassengerName(in.PassengerName).
		Save(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}

	if n != 1 {
		tx.Rollback()
		return nil, ErrDataUpdatedByAnotherProcess
	}

	// Return the updated data from the transaction.
	updatedSeat, err := tx.Seat.
		Query().
		Where(seat.ID(in.SeatID)).
		Only(ctx)
	if err != nil {
		tx.Rollback()
		return nil, err
	}
	tx.Commit()

	out := OutBookSeat{}
	err = mapstructure.Decode(updatedSeat, &out)
	return &out, err
}
