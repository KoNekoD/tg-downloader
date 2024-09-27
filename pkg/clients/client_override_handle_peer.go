package clients

import (
	"context"
	"fmt"
	"github.com/gotd/td/tg"
)

func (c *ClientOverride) handlePeer(ctx context.Context, peerClass tg.InputPeerClass) {
	var (
		channel *tg.InputPeerChannel
		ok      bool
	)

	if channel, ok = peerClass.(*tg.InputPeerChannel); !ok {
		return
	}

	for _, id := range c.channelsIds {
		if c.needStop() {
			return
		}

		if channel.ChannelID == id {
			fmt.Printf("channel found: %d\n", id)
			c.downloadChannelFiles(ctx, channel)
		}
	}
}
