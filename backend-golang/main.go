package main

import (
	"context"
	"fmt"
	"github.com/guillermoares/agile-engine/backend-golang/server"
	"os"
	"os/signal"
)

func main() {
	port := "8000"
	svr, ready := server.StartServer(port)
	for <-ready {
		fmt.Printf("Server listening at port %v\n", port)
	}

	// graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	for range c {
		ctx, cancel := context.WithTimeout(context.Background(), 2000)
		defer cancel()

		fmt.Println("Successfully shut down http server!")
		if err := svr.Shutdown(ctx); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		<-ctx.Done()
		os.Exit(0)
	}
}
