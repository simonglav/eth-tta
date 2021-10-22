package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/total-transactions-amount-eth/internal/handlers"
)

func serve(ctx context.Context) (err error) {

	rtr := mux.NewRouter()
	http.Handle("/", rtr)
	rtr.HandleFunc("/api/block/{block_number:[0-9]+}/total", handlers.ETHBlockTotal).Methods("GET")

	srv := &http.Server{
		Addr:    ":8080",
		Handler: rtr,
	}

	go func() {
		if err = srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen:%+s\n", err)
		}
	}()

	log.Println("Server started")

	<-ctx.Done()

	log.Println("Server stopped")

	ctxShutDown, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		cancel()
	}()

	if err = srv.Shutdown(ctxShutDown); err != nil {
		log.Fatalf("Server Shutdown Failed:%+s", err)
	}

	log.Println("Server exited properly")

	if err == http.ErrServerClosed {
		err = nil
	}
	return
}

func main() {

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		oscall := <-c
		log.Printf("System call:%+v\n", oscall)
		cancel()
	}()

	if err := serve(ctx); err != nil {
		log.Printf("Failed to serve:+%v\n", err)
	}
}
