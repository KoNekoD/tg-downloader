package clients

import (
	"context"
	"github.com/gotd/contrib/middleware/ratelimit"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/downloader"
	"golang.org/x/time/rate"
	"main/pkg/env"
	"time"
)

func NewClientOverride(ctx context.Context, e *env.Environment) *ClientOverride {
	middlewares := []telegram.Middleware{ratelimit.New(rate.Every(100*time.Millisecond), 5)}

	opts := telegram.Options{SessionStorage: getStorageMemory(ctx, e), Middlewares: middlewares}
	opts.Device = NewDeviceConfig(e.Device)

	client := telegram.NewClient(e.AppID, e.AppHash, opts)

	return &ClientOverride{
		Client:      client,
		rootCtx:     ctx,
		channelsIds: e.ChannelsIds,
		dl:          downloader.NewDownloader(),
		e:           e,
	}
}
