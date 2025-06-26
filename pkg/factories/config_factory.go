package factories

import (
	"github.com/KoNekoD/dotenv/pkg/dotenv"
	"main/pkg/dtos"
	"os"
)

func NewConfig() *dtos.Config {
	if err := dotenv.LoadEnv(); err != nil {
		panic(err)
	}

	return &dtos.Config{
		AppID:        asInt(os.Getenv("APP_ID")),
		AppHash:      os.Getenv("APP_HASH"),
		TDataPath:    os.Getenv("TDATA_PATH"),
		Device:       os.Getenv("DEVICE"),
		NeededUserId: asInt64(os.Getenv("NEEDED_USER_ID")),
		ChannelsIds:  asSliceOfInt64(os.Getenv("CHANNELS_IDS")),
	}
}
