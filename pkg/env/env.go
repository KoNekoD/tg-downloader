package env

import (
	"github.com/joho/godotenv"
	_ "main/resources"
	"os"
)

type Environment struct {
	AppID        int
	AppHash      string
	TdataPath    string
	Device       string
	NeededUserId int64
	ChannelsIds  []int64
}

func NewEnvironment() *Environment {
	if err := godotenv.Load("./resources/.env"); err != nil {
		panic(err)
	}

	return &Environment{
		AppID:        asInt(os.Getenv("APP_ID")),
		AppHash:      os.Getenv("APP_HASH"),
		TdataPath:    os.Getenv("TDATA_PATH"),
		Device:       os.Getenv("DEVICE"),
		NeededUserId: asInt64(os.Getenv("NEEDED_USER_ID")),
		ChannelsIds:  asSliceOfInt64(os.Getenv("CHANNELS_IDS")),
	}
}
