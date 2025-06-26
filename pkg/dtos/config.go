package dtos

type Config struct {
	AppID        int
	AppHash      string
	TDataPath    string
	Device       string
	NeededUserId int64
	ChannelsIds  []int64
}
