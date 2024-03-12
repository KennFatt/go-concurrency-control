package usecases

import (
	"context"
	"errors"
	"go-cc/ent"
	"go-cc/ent/seat"
	"time"

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

func optimisticUpdate(ctx context.Context, tx *ent.Tx, prev *ent.Seat, updateFunc func(*ent.SeatUpdate) *ent.SeatUpdate) (int, error) {
	nextVersion := time.Now().UnixNano()

	updater := tx.Seat.Update().Where(seat.ID(prev.ID), seat.Version(prev.Version)).SetVersion(uint64(nextVersion))

	saveUpdate := updateFunc(updater)
	updatedRows, err := saveUpdate.Save(ctx)
	if err != nil {
		return 0, err
	}

	if updatedRows != 1 {
		return 0, ErrDataUpdatedByAnotherProcess
	}

	return updatedRows, nil
}

func (us *UseCases) BookSeat(ctx context.Context, in InBookSeat) (*OutBookSeat, error) {
	tx, err := us.db.Tx(ctx)
	defer func() {
		if err != nil {
			tx.Rollback()
			return
		}

		tx.Commit()
	}()

	if err != nil {
		return nil, err
	}

	// Find the requested seat
	requestedSeat, err := tx.Seat.Query().Where(seat.ID(in.SeatID)).Only(ctx)
	if err != nil {
		return nil, err
	}

	if requestedSeat.IsBooked {
		return nil, ErrSeatIsBooked
	}

	// Book a seat
	// _, err = optimisticUpdate(ctx, tx, requestedSeat, func(su *ent.SeatUpdate) *ent.SeatUpdate {
	// 	return su.
	// 		SetIsBooked(true).
	// 		SetPassengerName(in.PassengerName)
	// })

	_, err = tx.Seat.
		Update().
		Where(seat.ID(in.SeatID)).
		SetIsBooked(true).
		SetPassengerName(in.PassengerName).
		Save(ctx)
	if err != nil {
		return nil, err
	}

	// Return the updated data from the transaction.
	updatedSeat, err := tx.Seat.
		Query().
		Where(seat.ID(in.SeatID)).
		Only(ctx)
	if err != nil {
		return nil, err
	}

	out := OutBookSeat{}
	err = mapstructure.Decode(updatedSeat, &out)
	return &out, err
}
