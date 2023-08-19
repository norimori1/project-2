package main

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {
	ctx, _ := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
	)

	srv := http.Server{
		Addr: "localhost:8080",
	}

	fmt.Println("Starting at :8080")

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			if !errors.Is(err, http.ErrServerClosed) {
				panic(err)
			}
		}
	}()

	<-ctx.Done()

	serverCtxTimeout, _ := context.WithTimeout(context.Background(), 5*time.Second)
	if err := srv.Shutdown(serverCtxTimeout); err != nil {
		panic(err)
	}

	fmt.Println("Done.")
}
