package providers

import (
	"context"
	"github.com/gotd/td/session"
	"github.com/gotd/td/session/tdesktop"
	"main/pkg/dtos"
)

func ProvidePathSessions(ctx context.Context, e *dtos.Config) []*dtos.PathSession {
	accounts, err := tdesktop.Read(e.TDataPath, nil)
	if err != nil {
		return nil
	}

	var pathSessions []*dtos.PathSession
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

		pathSessions = append(pathSessions, &dtos.PathSession{Acc: account, SMem: sMem})
	}

	return pathSessions
}
