# tg-downloader

`tg-downloader` is a Go-based utility designed to downloading all media from channels in Telegram. 

> [!NOTE]
> This project was inspired by the principles of [Copyism](https://en.wikipedia.org/wiki/Missionary_Church_of_Kopimism), a movement that advocates for the free sharing and replication of information. The author, motivated by these ideals, created this program to manage and download digital content efficiently.

## Overview

`tg-downloader` automates the process of downloading all media from specific Telegram channels. It works by interfacing with the Telegram API, using credentials and configuration provided via a `.env` file. The application does not require command-line arguments, making it simple to set up and use.

### Features:
- Downloads all media from specified Telegram channels.
- Utilizes a `.env` file for configuration, making setup straightforward.
- Supports multiple channels.
- Efficient, automated process for media removal on errors.

## Setup

To configure `tg-downloader`, create a `resources/.env` file in the root of the project directory, using the structure provided in the `resources/.env.example` file. This file will be automatically loaded by the program when it runs.

### Example `.env` file

```bash
# This file describes the parameters to run the program

# See https://core.telegram.org/api/obtaining_api_id
APP_ID=1
APP_HASH=""

# See https://www.google.com/search?q=How+to+find+tdata+directory+location
TDATA_PATH="/home/user/.local/share/TelegramDesktop/tdata"

# See https://www.google.com/search?q=How+to+find+out+your+user+id+on+Telegram
NEEDED_USER_ID=1111111111

# See https://www.google.com/search?q=How+to+find+out+the+id+of+a+Telegram+channel
CHANNELS_IDS=1,2,3

# Other, can be changed to any other name if desired
DEVICE="Honor 50"
```

### Parameters:

- **APP_ID**: Your Telegram API ID. Obtain it [here](https://core.telegram.org/api/obtaining_api_id).
- **APP_HASH**: Your Telegram API hash key.
- **TDATA_PATH**: Path to your Telegram Desktop `tdata` folder. This folder contains your user session data.
- **NEEDED_USER_ID**: Your Telegram user ID.
- **CHANNELS_IDS**: Comma-separated list of Telegram channel IDs from which media should be downloaded.
- **DEVICE**: Optional name for the device being used (can be customized).

## Usage

1. Clone the repository:
   ```bash
   git clone https://github.com/KoNekoD/tg-downloader.git
   ```
2. Navigate to the project directory:
   ```bash
   cd tg-downloader
   ```
3. Build the application:
   ```bash
   go build -o tg_downloader cmd/tg_downloader.go
   ```
4. Ensure your `.env` file is correctly set up with your Telegram credentials and channel information.
5. Run the program:
   ```bash
   ./tg_downloader
   ```

The program will automatically load your configuration from the `.env` file and download media from the specified Telegram channels.

## Inspiration

This project was strongly inspired by the tenets of Copyism, which promotes the free flow and replication of information. The author, motivated by these ideas, created `tg-downloader` to assist in the effective management of media in Telegram channels.
