package main

import (
	"context"
	"flag"
	"github.com/hecomp/catchall/internal/config"
	"github.com/hecomp/catchall/internal/db"
	"github.com/hecomp/catchall/pkg/app"
	"github.com/hecomp/catchall/pkg/repository"
	"github.com/hecomp/catchall/pkg/services"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	var httpAddr = flag.String("http.addr", "0.0.0.0:8080", "Address for HTTP (JSON) server")

	flag.Parse()

	l := log.New(os.Stdout, "catchall ", log.LstdFlags)

	dbConfig := config.NewConfigurations()
	conn, err := db.NewConnect(l, dbConfig)
	if err != nil {
		l.Printf("Error connecting to db: %s\n", err)
		os.Exit(1)
	}
	repo := repository.NewCatchAllRepository(conn)
	serv := services.NewCatchAllService(repo)
	// create the handlers
	catchAll := app.NewCatchAllHandler(l, serv)
	r := app.Routes(catchAll)

	// create a new server
	s := http.Server{
		Addr:         *httpAddr,         // configure the bind address
		Handler:      r,                 // set the default handler
		ErrorLog:     l,                 // set the logger for the server
		ReadTimeout:  5 * time.Second,   // max time to read request from the client
		WriteTimeout: 10 * time.Second,  // max time to write response to the client
		IdleTimeout:  120 * time.Second, // max time for connections using TCP Keep-Alive
	}

	// start the server
	go func() {
		l.Println("Starting server on port 8080")

		err := s.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	// trap sigterm or interupt and gracefully shutdown the server
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	s.Shutdown(ctx)
}
