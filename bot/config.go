package bot

import (
	"log"
	"os"
)

var bot_token string
var bot_api_endpoint string
var bot_api_endpoint_file string
var resize_api_url string

// var resize_api_url string

func init() {
	if val, ok := os.LookupEnv("TG_BOT_TOKEN"); ok {
		bot_token = val
	} else {
		// return "12345:12345"
		log.Fatal("Variable TG_BOT_TOKEN not set.")
		// return errors.New("Bot Token environment variable not set.")
	}
	bot_api_endpoint = "https://api.telegram.org/bot" + bot_token + "/"           // trailing slash needed.
	bot_api_endpoint_file = "https://api.telegram.org/file/bot" + bot_token + "/" // trailing slash needed.
	// resize_api_url = "https://4e5c.cf/tgbot/" + bot_token
	if val, ok := os.LookupEnv("TG_RESIZE_API_URL"); ok {
		resize_api_url = val
	}
}

func GetBotToken() string {
	return bot_token
}

func GetBotApiEndpoint(forFile bool) string {
	if forFile {
		return bot_api_endpoint_file
	} else {
		return bot_api_endpoint
	}
}

func GetResizeApiUrl() string {
	return resize_api_url
}
