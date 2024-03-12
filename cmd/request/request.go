package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"time"
)

const (
	FlightServiceURL = "http://localhost:1337"
)

type BookFlightSeatPayload struct {
	SeatID        int    `json:"seatId"`
	PassengerName string `json:"passengerName"`
}

type requestInfo struct {
	payload BookFlightSeatPayload
	err     error

	start   int64
	end     int64
	latency time.Duration

	response     []byte
	responseCode int
}

func bookFlightSeat(payload BookFlightSeatPayload, ch chan<- requestInfo) {
	endpoint := fmt.Sprintf("%s/seats", FlightServiceURL)

	data, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	resp, err := http.Post(endpoint, "application/json", bytes.NewBuffer(data))
	defer resp.Body.Close()
	responseBody, err := io.ReadAll(resp.Body)

	ch <- requestInfo{
		payload:      payload,
		err:          err,
		responseCode: resp.StatusCode,
		response:     responseBody,

		latency: time.Since(start).Round(time.Millisecond),
		start:   start.UnixMicro(),
		end:     time.Now().UnixMicro(),
	}
}

func main() {
	targetedFlightSeat := 18
	passangers := []BookFlightSeatPayload{
		{
			SeatID:        targetedFlightSeat,
			PassengerName: "Foo",
		},
		{
			SeatID:        targetedFlightSeat,
			PassengerName: "Bar",
		},
		{
			SeatID:        targetedFlightSeat,
			PassengerName: "Fizz",
		},
	}

	reqInfoCh := make(chan requestInfo)

	for _, passanger := range passangers {
		go bookFlightSeat(passanger, reqInfoCh)
	}

	// Wait and listen for finished requests
	for range passangers {
		reqInfo := <-reqInfoCh

		logArgs := []any{
			"start", reqInfo.start,
			"end", reqInfo.end,
			"lat", reqInfo.latency,

			"passanger", reqInfo.payload,

			"err", reqInfo.err,
			"response", string(reqInfo.response),
		}

		if reqInfo.responseCode != http.StatusOK {
			slog.Error("request sent", logArgs...)
		} else {
			slog.Info("request sent", logArgs...)
		}
	}
}
