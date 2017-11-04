package bot

import (
	"log"
	"os"
)

var tgBotToken string
var tgBotAPIEndpoint string
var tgBotAPIEndpointFile string

// var resizeApiUrl string

func init() {
	if val, ok := os.LookupEnv("TG_BOT_TOKEN"); ok {
		tgBotToken = val
	} else {
		tgBotToken = "SAMPLETOKEN"
		// return "12345:12345"
		log.Println("Variable TG_BOT_TOKEN not set.")
		// return errors.New("Bot Token environment variable not set.")
	}
	tgBotAPIEndpoint = "https://api.telegram.org/bot" + tgBotToken + "/"          // trailing slash needed.
	tgBotAPIEndpointFile = "https://api.telegram.org/file/bot" + tgBotToken + "/" // trailing slash needed.

}

func getBotToken() string {
	return tgBotToken
}

func getBotAPIEndpoint(forFile bool) string {
	if forFile {
		return tgBotAPIEndpointFile
	}
	return tgBotAPIEndpoint
}
