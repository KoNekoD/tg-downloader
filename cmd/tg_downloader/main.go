package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, os.Interrupt)

	wg.Add(1)
	go run(ctx, &wg)

	sig := <-sigCh
	fmt.Printf("\nReceived signal: %v\n%s", sig, "Stop program...\n")

	cancel()
	wg.Wait()

	fmt.Printf("\nBye 👋\n")
}
