package clients

import (
	"context"
	"github.com/gotd/td/session"
	"main/pkg/env"
)

func getStorageMemory(ctx context.Context, e *env.Environment) *session.StorageMemory {
	var sMem *session.StorageMemory
	for _, pathSession := range getPathSessions(ctx, e) {
		if pathSession.acc.Authorization.UserID == uint64(e.NeededUserId) {
			sMem = pathSession.sMem
			break
		}
	}
	if sMem == nil {
		panic("sMem is nil")
	}

	return sMem
}
