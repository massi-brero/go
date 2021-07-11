package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync/atomic"
	"time"
)

type key int

const (
	requestIDKey = 0
)

var (
	listenTo string
	healthStatus int32
)

func main() {
	flag.StringVar(&listenTo, "listen", ":8080", "address server listens to")
	flag.Parse()

	logger := log.New(os.Stdout, "http: ", log.LstdFlags)
	logger.Println("server started...")

	router := http.NewServeMux()
	router.Handle("/", index())
	router.Handle("/health", health())

	nextRequestId := func() string {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}

	server := &http.Server{
		Addr: listenTo,
		Handler: handler(nextRequestId)(logging(logger)(router)),
		ErrorLog: logger,
		ReadTimeout: 6 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	done := make(chan bool)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	// server shuts down, when (blocking) quit channel gets interrupt value
	go func()  {
		<-quit
		logger.Println("server shuts down...")
		atomic.StoreInt32(&healthStatus, 0)

		ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
		defer cancel()

		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Fatalf("could not gracefully shutdowwn the server: %v\n", err)
		}
		close(done)
	}()

	logger.Println("server is ready to handle requests at ", listenTo)
	atomic.StoreInt32(&healthStatus, 1)
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		logger.Fatalf("could not listen on %s: %v\n", listenTo, err)
	}

	<-done
	logger.Println("server stopped")
}

func index() http.Handler {
	//not implemented yet
}

func health() http.Handler {
	//not implemented yet
}

func logging(logger *log.Logger) func(http.Handler) http.Handler {
	return func(http.Handler) http.Handler {
		//not implemented yet
	}
}

func handler(nextRequestId func() string) func(http.Handler) http.Handler {
	return func(http.Handler) http.Handler {
		//not implemented yet
	}
}
