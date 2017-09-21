// updates_test.go

package bot

import (
	"log"
	"testing"
)

func TestProcessMessage(t *testing.T) {
	// update := Update()
	updateId := 1000
	messageId := 1001
	chatId := 1002
	userId := 1002
	user := User{Id: int64(userId), FirstName: "Test", LastName: "Person"}
	chat := Chat{Id: int64(chatId), Type: "private"}
	message := Message{MessageId: int64(messageId), FromUser: &user, Chat: &chat, Text: "Hello"}
	update := Update{UpdateId: int64(updateId), Message: &message}
	// update.UpdateId = updateId

	// update.Message = Message(123, User(123, false, "Test", "Person", "none"))
	// update.Message.Text = "Test Message"

	result, _ := ProcessMessage(update)
	// var s SendMessage
	log.Println(result)
	s := result.(*SendMessage)
	// s.method()
	switch {
	case s.Method != "sendMessage":
		t.Fatalf("unexpected method: %q", s.Method)
	case s.Text != "Can't talk. Send a photo.":
		t.Fatalf("unexpected text: %q", s.Text)
	case s.ChatId != int64(chatId):
		t.Fatalf("unexpected chatId: %q", s.ChatId)

	}
}
