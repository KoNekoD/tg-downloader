package dtos

import (
	"github.com/gotd/td/session"
	"github.com/gotd/td/session/tdesktop"
)

type PathSession struct {
	Acc  tdesktop.Account
	SMem *session.StorageMemory
}
