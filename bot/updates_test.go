// updates_test.go

package bot

import (
	"log"
	"os"
	"testing"
)

func TestProcessMessage(t *testing.T) {
	os.Setenv("TG_BOT_TOKEN", "12345:12345")
	// update := Update()
	updateID := 1000
	messageID := 1001
	chatID := 1002
	userID := 1002
	user := User{Id: int64(userID), FirstName: "Test", LastName: "Person"}
	chat := Chat{Id: int64(chatID), Type: "private"}
	message := Message{MessageId: int64(messageID), FromUser: &user, Chat: &chat, Text: "Hello"}
	update := Update{UpdateId: int64(updateID), Message: &message}
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
	case s.ChatID != int64(chatID):
		t.Fatalf("unexpected chatId: %q", s.ChatID)

	}
	os.Exit(0)
}
