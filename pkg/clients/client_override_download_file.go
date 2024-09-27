package clients

import (
	"context"
	"fmt"
	"github.com/gotd/td/telegram/query/messages"
	"os"
)

func (c *ClientOverride) downloadFile(file messages.File, ctx context.Context, filePath string, fileName string) {
	api := c.API()

	dlQb := c.dl.Download(api, file.Location)

	fmt.Printf("downloading: %s\n", fileName)

	_, err := dlQb.ToPath(ctx, filePath)
	if err != nil {
		fmt.Printf("download error: %v\n", err)

		if err = os.Remove(filePath); err != nil {
			fmt.Printf("remove error: %v\n", err)
		}
	}
}
