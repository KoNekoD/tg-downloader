package clients

import (
	"context"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/downloader"
	"main/pkg/dtos"
)

type ClientOverride struct {
	*telegram.Client
	rootCtx     context.Context
	channelsIds []int64
	dl          *downloader.Downloader
	e           *dtos.Config
}

func (c *ClientOverride) needStop() bool {
	return c.rootCtx.Err() != nil
}
