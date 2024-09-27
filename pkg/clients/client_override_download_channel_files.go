package clients

import (
	"context"
	"github.com/gotd/td/telegram/query"
	"github.com/gotd/td/tg"
)

func (c *ClientOverride) downloadChannelFiles(ctx context.Context, req *tg.InputPeerChannel) {
	api := c.API()

	q := query.NewQuery(api)

	getHistoryQb := q.Messages().GetHistory(req).BatchSize(100)

	getHistoryIter := getHistoryQb.Iter()
	for getHistoryIter.Next(ctx) {
		if c.needStop() {
			return
		}

		c.handleHistoryItem(ctx, getHistoryIter, req.ChannelID)
	}
}
