package main

import (
	"context"
	"main/pkg/clients"
	"main/pkg/factories"
	"sync"
)

func run(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()

	e := factories.NewConfig()

	client := clients.NewClientOverride(ctx, e)

	if err := client.Run(context.Background(), client.RunFunc); err != nil {
		panic(err)
	}
}
