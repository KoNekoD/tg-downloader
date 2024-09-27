package main

import (
	"context"
	"main/pkg/clients"
	"sync"
)

func run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	client := clients.NewClientOverride(ctx)

	if err := client.Run(context.Background(), client.RunFunc); err != nil {
		panic(err)
	}
}
