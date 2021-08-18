package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

var logger = log.New(os.Stdout, "planning-app", log.LstdFlags)
var port = 9090

func main() {

	sm := mux.NewRouter()

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", port), // configure the bind address
		Handler:      sm,                       // set the default handler
		ErrorLog:     logger,                   // set the logger for the server
		ReadTimeout:  5 * time.Second,          // max time to read request from the client
		WriteTimeout: 10 * time.Second,         // max time to write response to the client
		IdleTimeout:  120 * time.Second,        // max time for connections using TCP Keep-Alive
	}

	go func() {
		logger.Println("Starting server on port 9090")

		err := server.ListenAndServe()
		if err != nil {
			logger.Printf("Error starting server: %server\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	logger.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	_ = server.Shutdown(ctx)
}
