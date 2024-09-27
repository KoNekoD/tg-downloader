package main

import (
	"context"
	"main/pkg/clients"
	"main/pkg/env"
	"sync"
)

func run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	e := env.NewEnvironment()

	client := clients.NewClientOverride(ctx, e)

	if err := client.Run(context.Background(), client.RunFunc); err != nil {
		panic(err)
	}
}
