package bot

import (
	"log"
)

type Logger struct {
	Error error
	Data  string
}

func LogData(logObj Logger) {
	if logObj.Data == "" {
		log.Println(logObj.Data)
	}
	if logObj.Error != nil {
		log.Println(logObj.Error)
	}
}
