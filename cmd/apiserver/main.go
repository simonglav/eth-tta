package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"

	"github.com/simonglav/total-transactions-amount-eth/config"
	"github.com/simonglav/total-transactions-amount-eth/internal/handlers"
)

func main() {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/api/block/{block_number:[0-9]+}/total", handlers.ETHBlockTotal).Methods("GET")

	server := http.Server{
		Addr:    config.ServerAddr,
		Handler: rtr,
	}

	// Create context that listens for the interrupt signal from the OS
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt)
	defer stop()

	go server.ListenAndServe()

	// Listen for the interrupt signal
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown
	stop()
	log.Println("Shutting down gracefully, press Ctrl+C again to force")

	// Perform application shutdown with a maximum timeout of 5 seconds
	timeoutCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := server.Shutdown(timeoutCtx); err != nil {
		log.Println(err)
	}
}
