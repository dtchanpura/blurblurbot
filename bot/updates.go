package bot

import (
	// "fmt"
	"log"
	// "io"
	"fmt"
)

// ProcessMessage method
func ProcessMessage(update Update) (BaseMethod, error) {
	var r BaseMethod
	// var bgR BaseMethod

	log.Printf("update.Message.Chat.ID: %v\n", update.Message.Chat.ID)
	if update.Message.Photo != nil { // || update.Message.Document != nil {
		log.Println("update.Message.Photo != nil")
		photoFileID := getMaxResolutionPhoto(*update.Message.Photo).FileID
		photoCaption := update.Message.Caption
		scale, blurFactor, returnSameSize, _ := captionScaleBlur(photoCaption)
		s := SendChatAction{ChatID: update.Message.Chat.ID, Action: "upload_photo"}
		// TODO: ReplyMarkup to be added.
		s.method()
		r = &s
		go SendTgPhotoFormData(update.Message.Chat.ID, photoFileID, scale, blurFactor, returnSameSize)
	} else if update.Message.Document != nil {
		log.Println("update.Message.Document != nil")
		photoFileID := update.Message.Document.FileID
		photoCaption := update.Message.Caption
		scale, blurFactor, returnSameSize, _ := captionScaleBlur(photoCaption)
		s := SendChatAction{ChatID: update.Message.Chat.ID, Action: "upload_photo"}
		// TODO: ReplyMarkup to be added.
		s.method()
		r = &s
		go SendTgPhotoFormData(update.Message.Chat.ID, photoFileID, scale, blurFactor, returnSameSize)
	} else if update.Message.Text != "" {
		log.Println("update.Message.Text != nil")
		var s SendMessage
		if update.Message.Text == "/version" {
			s = SendMessage{Text: fmt.Sprintf("Version: %s\nDate: %s", Version, BuildDate), ChatID: update.Message.Chat.ID}
		} else if update.Message.Text == "/help" {
			s = SendMessage{Text: HelpText, ChatID: update.Message.Chat.ID}
		} else {
			s = SendMessage{Text: "Can't talk. Send a photo.", ChatID: update.Message.Chat.ID}
		}
		s.method()
		r = &s
	} else {
		log.Println("Unexpected message type")
		s := SendMessage{Text: "I wasn't expecting that.", ChatID: update.Message.Chat.ID}
		s.method()
		r = &s
	}
	return r, nil
}
