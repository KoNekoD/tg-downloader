package clients

import (
	"github.com/gotd/td/session"
	"github.com/gotd/td/session/tdesktop"
)

type PathSession struct {
	acc  tdesktop.Account
	sMem *session.StorageMemory
}
