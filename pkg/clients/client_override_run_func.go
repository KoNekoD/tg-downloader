package clients

import (
	"context"
	"github.com/gotd/td/telegram/query"
)

func (c *ClientOverride) RunFunc(ctx context.Context) error {
	dialogsIter := query.NewQuery(c.API()).GetDialogs().Iter()

	for dialogsIter.Next(ctx) {
		if c.needStop() {
			return nil
		}

		if dialogsIter.Err() != nil {
			return dialogsIter.Err()
		}

		c.handlePeer(ctx, dialogsIter.Value().Peer)
	}

	return nil
}
