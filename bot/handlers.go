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

	if update.Message != nil {
		resp, err := ProcessMessage(update)
		if err != nil {
			log.Println(err)
			// log.Print(err)
			NoResponse(writer)
		}
		encoder.Encode(&resp)

	} else if update.CallbackQuery != nil {
		log.Printf("Callback Query recieved: %v\n", update.CallbackQuery)
	} else {
		NoResponse(writer)
	}
}

func NoResponse(writer http.ResponseWriter) {
	writer.Header().Set("Content-Type", "application/json")
	fmt.Fprint(writer, "{}")
}
