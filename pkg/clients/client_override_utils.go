package clients

import (
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/tg"
)

func NewDeviceConfig(device string) telegram.DeviceConfig {
	return telegram.DeviceConfig{
		DeviceModel:    device,
		SystemVersion:  "Linux 5.15.75-Debian-x86_64",
		AppVersion:     "Telegraph 9.0.0",
		SystemLangCode: "en",
		LangPack:       "en",
		LangCode:       "en",
		Proxy:          tg.InputClientProxy{},
		Params:         nil,
	}
}
