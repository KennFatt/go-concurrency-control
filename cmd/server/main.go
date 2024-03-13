package main

import (
	"fmt"
	"go-cc/conf"
	"go-cc/data"
	"go-cc/domain"
	"log"
	"net/http"
)

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
