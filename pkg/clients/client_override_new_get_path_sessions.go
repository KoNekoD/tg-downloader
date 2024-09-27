package clients

import (
	"context"
	"github.com/gotd/td/session"
	"github.com/gotd/td/session/tdesktop"
)

func getPathSessions(ctx context.Context) []*PathSession {
	accounts, err := tdesktop.Read(path, nil)
	if err != nil {
		return nil
	}

	var pathSessions []*PathSession
	for _, account := range accounts {
		data, err := session.TDesktopSession(account)
		if err != nil {
			continue
		}

		var (
			sMem   = &session.StorageMemory{}
			loader = session.Loader{Storage: sMem}
		)

		if err = loader.Save(ctx, data); err != nil { // *session.Data to InMemorySession
			continue
		}

		pathSessions = append(pathSessions, &PathSession{acc: account, sMem: sMem})
	}

	return pathSessions
}
