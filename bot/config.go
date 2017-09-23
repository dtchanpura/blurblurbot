package bot

import (
	"log"
	"os"
)

var tgBotToken string
var tgBotApiEndpoint string
var tgBotApiEndpointFile string
var resizeApiUrl string

// var resizeApiUrl string

func init() {
	if val, ok := os.LookupEnv("TG_BOT_TOKEN"); ok {
		tgBotToken = val
	} else {
		// return "12345:12345"
		log.Fatal("Variable TG_BOT_TOKEN not set.")
		// return errors.New("Bot Token environment variable not set.")
	}
	tgBotApiEndpoint = "https://api.telegram.org/bot" + tgBotToken + "/"          // trailing slash needed.
	tgBotApiEndpointFile = "https://api.telegram.org/file/bot" + tgBotToken + "/" // trailing slash needed.
	// resizeApiUrl = "https://4e5c.cf/tgbot/" + tgBotToken
	if val, ok := os.LookupEnv("TG_RESIZE_API_URL"); ok {
		resizeApiUrl = val
	}
}

func GetBotToken() string {
	return tgBotToken
}

func GetBotApiEndpoint(forFile bool) string {
	if forFile {
		return tgBotApiEndpointFile
	} else {
		return tgBotApiEndpoint
	}
}

func GetResizeApiUrl() string {
	return resizeApiUrl
}
