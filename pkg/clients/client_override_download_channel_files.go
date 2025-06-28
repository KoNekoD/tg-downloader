package clients

import (
	"context"
	"fmt"
	"github.com/gotd/td/telegram/query"
	"github.com/gotd/td/telegram/query/messages"
	"github.com/gotd/td/tg"
	"math/rand"
	"time"
)

func (c *ClientOverride) downloadChannelFiles(ctx context.Context, req *tg.InputPeerChannel) {
	api := c.API()

	q := query.NewQuery(api)

	const BatchSize = 100

	getHistoryQb := q.Messages().GetHistory(req).BatchSize(BatchSize)

	getHistoryIter := getHistoryQb.Iter()

	first := true
	firstMessageId := 0

	i := BatchSize
	for c.nextHistoryValue(getHistoryIter, ctx) {
		if c.needStop() {
			return
		}

		if first {
			first = false
			firstMessageId = getHistoryIter.Value().Msg.GetID()
		}

		c.handleHistoryItem(ctx, getHistoryIter, req.ChannelID)

		i--
		if i == 0 {
			lastMessageId := getHistoryIter.Value().Msg.GetID()

			fmt.Printf("handled batch between %d - %d\n", firstMessageId, lastMessageId)

			first = true
			i = BatchSize

			c.smartWait(ctx)
		}
	}

	if getHistoryIter.Err() != nil {
		fmt.Printf("get next history error: %v\n", getHistoryIter.Err())
	}
}

func (c *ClientOverride) smartWait(ctx context.Context) {
	fmt.Println("smart wait")
	defer fmt.Println("smart wait done")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	select {
	case <-ctx.Done():
		return
	case <-time.After(100 * time.Second):
		fmt.Println("done wait 100 sec")
		return
	case <-time.After(time.Duration(10+r.Intn(100)) * time.Second):
		fmt.Println("done wait 10-110 sec")
		return
	}
}

func (c *ClientOverride) nextHistoryValue(iter *messages.Iterator, ctx context.Context) bool {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	return iter.Next(ctx)
}
