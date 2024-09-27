package clients

import (
	"context"
	"github.com/gotd/td/session"
)

func getStorageMemory(ctx context.Context) *session.StorageMemory {
	var sMem *session.StorageMemory
	for _, pathSession := range getPathSessions(ctx) {
		if pathSession.acc.Authorization.UserID == neededUserId {
			sMem = pathSession.sMem
			break
		}
	}
	if sMem == nil {
		panic("sMem is nil")
	}

	return sMem
}
