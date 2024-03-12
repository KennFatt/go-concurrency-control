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
	PassangerName string `json:"passangerName"`
}

type requestInfo struct {
	payload BookFlightSeatPayload
	err     error

	start   int64
	end     int64
	latency time.Duration

	response        []byte
	isResponseError int
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
		payload:         payload,
		err:             err,
		isResponseError: resp.StatusCode,
		response:        responseBody,

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
			PassangerName: "Foo",
		},
		{
			SeatID:        targetedFlightSeat,
			PassangerName: "Bar",
		},
		{
			SeatID:        targetedFlightSeat,
			PassangerName: "Fizz",
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

		if reqInfo.isResponseError != http.StatusOK {
			slog.Error("request sent", logArgs...)
		} else {
			slog.Info("request sent", logArgs...)
		}
	}
}
