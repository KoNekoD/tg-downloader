package providers

import (
	"context"
	"fmt"
	"github.com/gotd/td/session"
	"main/pkg/dtos"
)

func ProvideStorageMemory(ctx context.Context, e *dtos.Config) *session.StorageMemory {
	var storageMemory *session.StorageMemory

	pathSessions := ProvidePathSessions(ctx, e)

	for _, pathSession := range pathSessions {
		if pathSession.Acc.Authorization.UserID == uint64(e.NeededUserId) {
			fmt.Printf("Found user: %d\n", pathSession.Acc.Authorization.UserID)

			storageMemory = pathSession.SMem
			break
		}

		fmt.Printf("Not this user: %d\n", pathSession.Acc.Authorization.UserID)
	}

	if storageMemory == nil {
		panic("Not found selected user in TData")
	}

	return storageMemory
}
