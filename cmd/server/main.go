package main

import (
	"fmt"
	"go-cc/conf"
	"go-cc/data"
	"go-cc/domain"
	"log"
	"net/http"
)

// TODO:
// - [x] Read a DB config
// - [x] Initiate a Ent client
//   - [x] Define the seat schema
//
// - [x] Initiate a Mux
//   - [x] Get list of the available seats
//   - [x] Book a specific seat
//   - [x] Reproduce the problem
//   - [x] Solve the problem with OCC
func main() {
	config, err := conf.ReadConfig()
	failureOnError("fail to read the config", err)

	db, err := data.NewDB(config.Database)
	failureOnError("fail to instantiate Ent client", err)

	mux := http.NewServeMux()
	domain.InitRoutes(mux, db)

	serverAddr := fmt.Sprintf(":%d", config.Application.Port)
	log.Printf("the server is running on http://localhost%s\n", serverAddr)
	log.Fatal(
		http.ListenAndServe(serverAddr, mux),
	)
}

func failureOnError(message string, err error) {
	if err == nil {
		return
	}

	log.Fatal(message, err)
}
