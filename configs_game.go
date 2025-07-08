package tgbotapi

/*
sendGame
Use this method to send a game. On success, the sent Message is returned.

Parameter	Type	Required	Description
business_connection_id	String	Optional	Unique identifier of the business connection on behalf of which the message will be sent
chat_id	Integer	Yes	Unique identifier for the target chat
message_thread_id	Integer	Optional	Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
game_short_name	String	Yes	Short name of the game, serves as the unique identifier for the game. Set up your games via @BotFather.
disable_notification	Boolean	Optional	Sends the message silently. Users will receive a notification with no sound.
protect_content	Boolean	Optional	Protects the contents of the sent message from forwarding and saving
allow_paid_broadcast	Boolean	Optional	Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
message_effect_id	String	Optional	Unique identifier of the message effect to be added to the message; for private chats only
reply_parameters	ReplyParameters	Optional	Description of the message to reply to
reply_markup	InlineKeyboardMarkup	Optional	A JSON-serialized object for an inline keyboard. If empty, one 'Play game_title' button will be shown. If not empty, the first button must launch the game.
*/

type SendGameConfig struct {
	BaseChat
	GameShortName       string               `json:"game_short_name"`
	DisableNotification bool                 `json:"disable_notification"`
	ProtectContent      bool                 `json:"protect_content"`
	AllowPaidBroadcast  bool                 `json:"allow_paid_broadcast"`
	MessageEffectID     string               `json:"message_effect_id"`
	ReplyParameters     ReplyParameters      `json:"reply_parameters"`
	ReplyMarkup         InlineKeyboardMarkup `json:"reply_markup"`
}

func (config SendGameConfig) method() string {
	return "sendGame"
}

func (config SendGameConfig) params() (Params, error) {
	params, _ := config.BaseChat.params()
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("game_short_name", config.GameShortName)
	params.AddBool("disable_notification", config.DisableNotification)
	params.AddBool("protect_content", config.ProtectContent)
	params.AddBool("allow_paid_broadcast", config.AllowPaidBroadcast)
	params.AddNonEmpty("message_effect_id", config.MessageEffectID)
	params.AddInterface("reply_parameters", config.ReplyParameters)
	params.AddInterface("reply_markup", config.ReplyMarkup)
	return params, nil
}

// SetGameScoreConfig allows you to update the game score in a chat.
type SetGameScoreConfig struct {
	UserID             int64
	Score              int
	Force              bool
	DisableEditMessage bool
	ChatID             int64
	ChannelUsername    string
	MessageID          int
	InlineMessageID    string
}

func (config SetGameScoreConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonZero64("user_id", config.UserID)
	params.AddNonZero("scrore", config.Score)
	params.AddBool("disable_edit_message", config.DisableEditMessage)

	if config.InlineMessageID != "" {
		params["inline_message_id"] = config.InlineMessageID
	} else {
		params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
		params.AddNonZero("message_id", config.MessageID)
	}

	return params, nil
}

func (config SetGameScoreConfig) method() string {
	return "setGameScore"
}

// GetGameHighScoresConfig allows you to fetch the high scores for a game.
type GetGameHighScoresConfig struct {
	UserID          int64
	ChatID          int64
	ChannelUsername string
	MessageID       int
	InlineMessageID string
}

func (config GetGameHighScoresConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonZero64("user_id", config.UserID)

	if config.InlineMessageID != "" {
		params["inline_message_id"] = config.InlineMessageID
	} else {
		params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
		params.AddNonZero("message_id", config.MessageID)
	}

	return params, nil
}

func (config GetGameHighScoresConfig) method() string {
	return "getGameHighScores"
}
