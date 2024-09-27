package clients

import (
	"context"
	"errors"
	"fmt"
	"github.com/gotd/td/telegram/query/messages"
	"os"
)

func (c *ClientOverride) handleHistoryItem(ctx context.Context, getHistoryIter *messages.Iterator, channelId int64) {
	if getHistoryIter.Err() != nil {
		fmt.Printf("get history error: %v\n", getHistoryIter.Err())
	}

	msg := getHistoryIter.Value()

	file, ok := msg.File()
	if !ok {
		return
	}

	dirPath := fmt.Sprintf("./data/%d", channelId)
	if _, err := os.Stat(dirPath); errors.Is(err, os.ErrNotExist) {
		if err := os.MkdirAll(dirPath, 0755); err != nil {
			fmt.Printf("create dir error: %v\n", err)
		}
	}

	fileName := fmt.Sprintf("%d_%s", msg.Msg.GetID(), file.Name)
	filePath := fmt.Sprintf("%s/%s", dirPath, fileName)
	if _, err := os.Stat(filePath); !errors.Is(err, os.ErrNotExist) {
		return
	}

	c.downloadFile(file, ctx, filePath, fileName)
}
