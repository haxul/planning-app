package main

import (
	"context"
	"fmt"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/haxul/planning-app/backend/common"
	"github.com/haxul/planning-app/backend/controller"
	"github.com/haxul/planning-app/backend/persistance/postgres"
	"github.com/jackc/pgx/v4"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	sm := mux.NewRouter()

	// GET
	getRouter := sm.Methods(http.MethodGet).Subrouter()
	getRouter.HandleFunc("/card", controller.GetCardsCtrlInstance().GetAllCards)

	// POST
	postRouter := sm.Methods(http.MethodPost).Subrouter()
	postRouter.HandleFunc("/card", controller.GetCardsCtrlInstance().CreateCard)
	postRouter.HandleFunc("/card/{id}/move", controller.GetCardsCtrlInstance().MoveCard)
	postRouter.HandleFunc("/card/{id}/reject", controller.GetCardsCtrlInstance().RejectCard)

	// CORS
	cors := handlers.CORS(
		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS", "DELETE"}),
		handlers.AllowedOrigins([]string{"*"}))

	// create server
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", common.Port), // configure the bind address
		Handler:      cors(sm),                        // set the default handler
		ErrorLog:     common.Logger,                   // set the logger for the server
		ReadTimeout:  5 * time.Second,                 // max time to read request from the client
		WriteTimeout: 10 * time.Second,                // max time to write response to the client
		IdleTimeout:  120 * time.Second,               // max time for connections using TCP Keep-Alive
	}

	//postgres
	connection, errPs := pgx.Connect(context.Background(), "user=haxul password=test host=localhost port=5432 dbname=planning_db")
	if errPs != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s", errPs)
		os.Exit(1)
	}
	postgres.PostgreConn = connection
	defer postgres.PostgreConn.Close(context.Background())

	go func() {
		common.Logger.Println(fmt.Sprintf("Starting server on port %d", common.Port))

		err := server.ListenAndServe()
		if err != nil {
			common.Logger.Printf("Error starting server: %server\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	sig := <-c
	common.Logger.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	cancel()
	_ = server.Shutdown(ctx)
}
