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
		// photoApiUrl := createApiUrl(photoFileId, 0, 0)
		// s := SendPhoto{ChatId: update.Message.Chat.Id, Photo: photoApiUrl}
		s := SendChatAction{ChatId: update.Message.Chat.Id, Action: "upload_photo"}
		// ReplyMarkup to add....
		s.method()
		r = &s
		// bgS := SendPhoto{ChatId: update.Message.Chat.Id, Photo: photoApiUrl}
		// bgR = &bgS
		// go SendTgApiRequest(bgR) // send photo in background
		go SendTgPhotoFormData(update.Message.Chat.Id, photoFileId, scale, blurFactor, returnSameSize)
	} else if update.Message.Document != nil { // || update.Message.Document != nil {
		log.Println("update.Message.Document != nil")
		photoFileId := update.Message.Document.FileId
		photoApiUrl := createApiUrl(photoFileId, 0, 0)
		s := SendPhoto{ChatId: update.Message.Chat.Id, Photo: photoApiUrl}
		// ReplyMarkup to add....
		s.method()
		r = &s
	} else if update.Message.Text != "" {
		log.Println("update.Message.Text != nil")
		s := SendMessage{Text: "Can't talk. Send a photo.", ChatId: update.Message.Chat.Id}
		s.method()
		r = &s
	} else {
		log.Println("Unexpected message type")
		s := SendMessage{Text: "Some error occurred. We are checking it.", ChatId: update.Message.Chat.Id}
		s.method()
		r = &s
	}
	// r := BaseMethod{Method: "sendMessage"} // , SendMessage: s}
	// fmt.Println(r.method())
	// r.Text = "text"
	// r.Method = "sendMessage"
	return r, nil
}
