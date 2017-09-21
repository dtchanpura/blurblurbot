package bot

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func BotUpdateHandler(writer http.ResponseWriter, request *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(request.Body)
	encoder := json.NewEncoder(writer)
	var update Update
	err := decoder.Decode(&update)
	if err != nil {
		log.Println(err)
		NoResponse(writer)
	}

	// log.Printf("Message: %v\n", update.Message != nil)
	// log.Printf("CallbackQuery: %v\n", update.CallbackQuery != nil)

	if update.Message != nil {
		resp, err := ProcessMessage(update)
		if err != nil {
			log.Println(err)
			// log.Print(err)
			NoResponse(writer)
		}
		encoder.Encode(&resp)

	} else {
		NoResponse(writer)
	}
}

func NoResponse(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(writer, "{}")
}
