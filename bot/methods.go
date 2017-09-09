package blurblurbot

type BaseMethod interface {
	method() string
}

// Methods

type SendMessage struct {
	BaseMethod            `json:"-"`
	Method                string                `json:"method"`
	ChatId                int64                 `json:"chat_id"`
	Text                  string                `json:"text"`
	ParseMode             string                `json:"parse_mode,omitempty"`
	DisableWebPagePreview bool                  `json:"disable_web_page_preview,omitempty"`
	DisableNotification   bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageId      int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup           *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (sendMessage *SendMessage) method() string {
	sendMessage.Method = "sendMessage"
	return "sendMessage"
}

type SendPhoto struct {
	BaseMethod          `json:"-"`
	Method              string                `json:"method"`
	ChatId              int64                 `json:"chat_id"`
	Photo               string                `json:"photo,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	DisableNotification bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (sendPhoto *SendPhoto) method() string {
	sendPhoto.Method = "sendPhoto"
	return "sendPhoto"
}

type SendDocument struct {
	BaseMethod          `json:"-"`
	Method              string                `json:"method"`
	ChatId              int64                 `json:"chat_id"`
	Document            string                `json:"photo,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	DisableNotification bool                  `json:"disable_notification,omitempty"`
	ReplyToMessageId    int64                 `json:"reply_to_message_id,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

func (sendDocument *SendDocument) method() string {
	sendDocument.Method = "sendDocument"
	return "sendDocument"
}

type SendChatAction struct {
	BaseMethod `json:"-"`
	Method     string `json:"method"`
	ChatId     int64  `json:"chat_id"`
	Action     string `json:"action"`
}

func (sendChatAction *SendChatAction) method() string {
	sendChatAction.Method = "sendChatAction"
	return "sendChatAction"
}
