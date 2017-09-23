package bot

import (
	// "fmt"
	"log"
	// "io"
)

func ProcessMessage(update Update) (BaseMethod, error) {
	var r BaseMethod
	// var bgR BaseMethod

	log.Printf("update.Message.Chat.Id: %v\n", update.Message.Chat.Id)
	if update.Message.Photo != nil { // || update.Message.Document != nil {
		log.Println("update.Message.Photo != nil")
		photoFileId := getMaxResolutionPhoto(*update.Message.Photo).FileId
		photoCaption := update.Message.Caption
		scale, blurFactor, returnSameSize, _ := captionScaleBlur(photoCaption)
		s := SendChatAction{ChatId: update.Message.Chat.Id, Action: "upload_photo"}
		// TODO: ReplyMarkup to be added.
		s.method()
		r = &s
		go SendTgPhotoFormData(update.Message.Chat.Id, photoFileId, scale, blurFactor, returnSameSize)
	} else if update.Message.Document != nil {
		log.Println("update.Message.Document != nil")
		photoFileId := update.Message.Document.FileId
		photoCaption := update.Message.Caption
		scale, blurFactor, returnSameSize, _ := captionScaleBlur(photoCaption)
		s := SendChatAction{ChatId: update.Message.Chat.Id, Action: "upload_photo"}
		// TODO: ReplyMarkup to be added.
		s.method()
		r = &s
		go SendTgPhotoFormData(update.Message.Chat.Id, photoFileId, scale, blurFactor, returnSameSize)
	} else if update.Message.Text != "" {
		log.Println("update.Message.Text != nil")
		s := SendMessage{Text: "Can't talk. Send a photo.", ChatId: update.Message.Chat.Id}
		s.method()
		r = &s
	} else {
		log.Println("Unexpected message type")
		s := SendMessage{Text: "I wasn't expecting that.", ChatId: update.Message.Chat.Id}
		s.method()
		r = &s
	}
	return r, nil
}
