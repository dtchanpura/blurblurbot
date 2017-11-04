package bot

import (
	"fmt"
	// "html"
	// "errors"
	"log"
	"net/http"
	// "os"
)

// func GetBotToken() {
//
// }

// Run method for starting the server.
func Run() {
	// Following for debugging purposes only.
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, r.URL.Path)
	})

	// For Handling Bot Updates
	// log.Println(GetBotToken())
	http.HandleFunc("/"+getBotToken(), UpdateHandler)
	log.Println("Listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
	// return nil
}
