package bot

// Update struct for decoding Update Object from Telegram API
type Update struct {
	UpdateID           int64               `json:"update_id"`
	Message            *Message            `json:"message,omitempty"`
	EditedMessage      *Message            `json:"edited_message,omitempty"`
	ChannelPost        *Message            `json:"channel_post,omitempty"`
	EditChannelPost    *Message            `json:"edit_channel_post,omitempty"`
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	CallbackQuery      *CallbackQuery      `json:"callback_query,omitempty"`
	// ShippingQuery ShippingQuery `json:"shipping_query"`
	// PreCheckoutQuery PreCheckoutQuery `json:"pre_checkout_query"`
}

// Message struct for decoding Message Object from Telegram API
type Message struct {
	MessageID             int64            `json:"message_id"`
	FromUser              *User            `json:"from"`
	Date                  int64            `json:"date"`
	Chat                  *Chat            `json:"chat"`
	ForwardFrom           *User            `json:"forward_from"`
	ForwardFromChat       *Chat            `json:"forward_from_chat"`
	ForwardFromMessageID  int64            `json:"forward_from_message_id"`
	ForwardSignature      string           `json:"forward_signature"`
	ForwardDate           int64            `json:"forward_date"`
	ReplyToMessage        *Message         `json:"reply_to_message"`
	EditDate              int64            `json:"edit_date"`
	AuthorSignature       string           `json:"author_signature"`
	Text                  string           `json:"text"`
	Entities              *[]MessageEntity `json:"entities"`
	Audio                 *Audio           `json:"audio"`
	Document              *Document        `json:"document"`
	Game                  *Game            `json:"game"`
	Photo                 *[]PhotoSize     `json:"photo"`
	Sticker               *Sticker         `json:"sticker"`
	Video                 *Video           `json:"video"`
	Voice                 *Voice           `json:"voice"`
	VideoNote             *VideoNote       `json:"video_note"`
	NewChatMembers        *[]User          `json:"new_chat_members"`
	Caption               string           `json:"caption"`
	Contact               *Contact         `json:"contact"`
	Location              *Location        `json:"location"`
	Venue                 *Venue           `json:"venue"`
	NewChatMember         *User            `json:"new_chat_member"`
	LeftChatMember        *User            `json:"left_chat_member"`
	NewChatTitle          bool             `json:"new_chat_title"`
	NewChatPhoto          bool             `json:"new_chat_photo"`
	DeleteChatPhoto       bool             `json:"delete_chat_photo"`
	GroupChatCreated      bool             `json:"group_chat_created"`
	SupergroupChatCreated bool             `json:"supergroup_chat_created"`
	ChannelChatCreated    bool             `json:"channel_chat_created"`
	MigrateToChatID       int64            `json:"migrate_to_chat_id"`
	MigrateFromChatID     int64            `json:"migrate_from_chat_id"`
	// PinnedMessage         PinnedMessage         `json:"pinned_message"`
	// Invoice Invoice `json:"invoice"`
	// SuccessfulPayment SuccessfulPayment `json:"successful_payment"`
}

// User struct for decoding User Object from Telegram API
type User struct {
	ID           int64  `json:"id"`
	IsBot        bool   `json:"is_bot"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Username     string `json:"username"`
	LanguageCode string `json:"language_code"`
}

// Chat struct for decoding Chat Object from Telegram API
type Chat struct {
	ID                          int64      `json:"id"`
	Type                        string     `json:"type"` // Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Title                       string     `json:"title"`
	Username                    string     `json:"username"`
	FirstName                   string     `json:"first_name"`
	LastName                    string     `json:"last_name"`
	AllMembersAreAdministrators bool       `json:"all_members_are_administrators"`
	Photo                       *ChatPhoto `json:"photo"`
	Description                 string     `json:"description"`
	InviteLink                  string     `json:"invite_link"`
	PinnedMessage               *Message   `json:"pinned_message"`
}

// MessageEntity struct for decoding MessageEntity Object from Telegram API
type MessageEntity struct {
	Type   string `json:"type"` //Type of the entity. Can be mention (@username), hashtag, bot_command, url, email, bold (bold text), italic (italic text), code (monowidth string), pre (monowidth block), text_link (for clickable text URLs), text_mention (for users without usernames)
	Offset int    `json:"offset"`
	Length int    `json:"length"`
	URL    string `json:"url"`
	User   User   `json:"user"`
}

// PhotoSize struct for decoding PhotoSize Object from Telegram API
type PhotoSize struct {
	FileID   string `json:"file_id"`
	Width    int    `json:"width"`
	Height   int    `json:"height"`
	FileSize int    `json:"file_size"`
}

// Audio struct for decoding Audio Object from Telegram API
type Audio struct {
	FileID    string `json:"file_id"`
	Duration  int    `json:"duration"`
	Performer string `json:"performer"`
	Title     string `json:"title"`
	MimeType  string `json:"mime_type"`
	FileSize  int    `json:"file_size"`
}

// Document struct for decoding Document Object from Telegram API
type Document struct {
	FileID   string      `json:"file_id"`
	Thumb    []PhotoSize `json:"thumb"`
	FileName string      `json:"file_name"`
	MimeType string      `json:"mime_type"`
	FileSize int         `json:"file_size"`
}

// Video struct for decoding Video Object from Telegram API
type Video struct {
	FileID   string      `json:"file_id"`
	Width    int         `json:"width"`
	Height   int         `json:"height"`
	Duration int         `json:"duration"`
	Thumb    []PhotoSize `json:"thumb"`
	MimeType string      `json:"mime_type"`
	FileSize int         `json:"file_size"`
}

// Voice struct for decoding Voice Object from Telegram API
type Voice struct {
	FileID   string `json:"file_id"`
	Duration int    `json:"duration"`
	MimeType string `json:"mime_type"`
	FileSize int    `json:"file_size"`
}

// VideoNote struct for decoding VideoNote Object from Telegram API
type VideoNote struct {
	FileID   string      `json:"file_id"`
	Length   int         `json:"length"`
	Duration int         `json:"duration"`
	Thumb    []PhotoSize `json:"thumb"`
	FileSize int         `json:"file_size"`
}

// Contact struct for decoding Contact Object from Telegram API
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	UserID      int64  `json:"user_id"`
}

// Location struct for decoding Location Object from Telegram API
type Location struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Venue struct for decoding Venue Object from Telegram API
type Venue struct {
	Location     Location `json:"location"`
	Title        string   `json:"title"`
	Address      string   `json:"address"`
	FoursquareID string   `json:"foursquare_id"`
}

// UserProfilePhotos struct for decoding UserProfilePhotos Object from Telegram API
type UserProfilePhotos struct {
	TotalCount int           `json:"total_count"`
	Photos     [][]PhotoSize `json:"photos"`
}

// ChatPhoto struct for decoding ChatPhoto Object from Telegram API
type ChatPhoto struct {
	SmallFileID string `json:"small_file_id"`
	BigFileID   string `json:"big_file_id"`
}

// File struct for decoding File Object from Telegram API
type File struct {
	FileID   string `json:"file_id"`
	FileSize int    `json:"file_size"`
	FilePath string `json:"file_path"`
}

// Game struct for decoding Game Object from Telegram API
type Game struct {
	Title        string          `json:"title"`
	Description  string          `json:"description"`
	Photo        []PhotoSize     `json:"photo"`
	Text         string          `json:"text"`
	TextEntities []MessageEntity `json:"text_entities"`
	Animation    Animation       `json:"animation"`
}

// Animation struct for decoding Animation Object from Telegram API
type Animation struct {
	FileID   string      `json:"file_id"`
	Thumb    []PhotoSize `json:"thumb"`
	FileName string      `json:"file_name"`
	MimeType string      `json:"mime_type"`
	FileSize int         `json:"file_size"`
}

// Sticker struct for decoding Sticker Object from Telegram API
type Sticker struct {
	FileID  string      `json:"file_id"`
	Width   int         `json:"width"`
	Height  int         `json:"height"`
	Thumb   []PhotoSize `json:"thumb"`
	Emoji   string      `json:"emoji"`
	SetName string      `json:"set_name"`
	// MaskPosition MaskPosition `json:"mask_position"`
	FileSize int `json:"file_size"`
}

// InlineQuery struct for decoding InlineQuery Object from Telegram API
type InlineQuery struct {
	InlineQueryID int64
}

// ChosenInlineResult struct for decoding ChosenInlineResult Object from Telegram API
type ChosenInlineResult struct {
	ChosenInlineResultID int64
}

// CallbackQuery struct for decoding CallbackQuery Object from Telegram API
type CallbackQuery struct {
	CallbackQueryID int64
}

// InlineKeyboardMarkup struct for decoding InlineKeyboardMarkup Object from Telegram API
type InlineKeyboardMarkup struct {
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard,omitempty"`
}

// InlineKeyboardButton struct for decoding InlineKeyboardButton Object from Telegram API
type InlineKeyboardButton struct {
	Text                         string `json:"text,omitempty"`
	URL                          string `json:"url,omitempty"`
	CallbackData                 string `json:"callback_data,omitempty"`
	SwitchInlineQuery            string `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string `json:"switch_inline_query_current_chat,omitempty"`
	// `json:"callback_game"`
	// `json:"pay"`
}

// FileResult struct for decoding FileResult Object from Telegram API
type FileResult struct {
	Ok     bool  `json:"ok"`
	Result *File `json:"result"`
}

// ByResolution for sorting photos by its Resolution
type ByResolution []PhotoSize

func (a ByResolution) Len() int           { return len(a) }
func (a ByResolution) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByResolution) Less(i, j int) bool { return a[i].Width*a[i].Height < a[j].Width*a[j].Height }
