package tgbotapi

/*
sendChecklist
Use this method to send a checklist on behalf of a connected business account. On success, the sent Message is returned.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection on behalf of which the message will be sent
chat_id	Integer	Yes	Unique identifier for the target chat
checklist	InputChecklist	Yes	A JSON-serialized object for the checklist to send
disable_notification	Boolean	Optional	Sends the message silently. Users will receive a notification with no sound.
protect_content	Boolean	Optional	Protects the contents of the sent message from forwarding and saving
message_effect_id	String	Optional	Unique identifier of the message effect to be added to the message
reply_parameters	ReplyParameters	Optional	A JSON-serialized object for description of the message to reply to
reply_markup	InlineKeyboardMarkup	Optional	A JSON-serialized object for an inline keyboard
*/

type SendChecklistConfig struct {
	BaseChat
	Checklist       Checklist       `json:"checklist"`
	MessageEffectID string          `json:"message_effect_id"`
	ReplyParameters ReplyParameters `json:"reply_parameters"`
}

func (config SendChecklistConfig) params() (Params, error) {
	params, err := config.BaseChat.params()
	if err != nil {
		return nil, err
	}
	params.AddInterface("checklist", config.Checklist)
	params.AddNonEmpty("message_effect_id", config.MessageEffectID)
	params.AddInterface("reply_parameters", config.ReplyParameters)

	err = params.CheckArgs("business_connection_id", "chat_id", "checklist")
	if err != nil {
		return nil, err
	}
	return params, nil
}

func (config SendChecklistConfig) method() string {
	return "sendChecklist"
}

type EditMessageChecklistConfig struct {
	BaseChat
	MessageID int64     `json:"message_id"`
	Checklist Checklist `json:"checklist"`
}

func (config EditMessageChecklistConfig) params() (Params, error) {
	params, err := config.BaseChat.params()
	if err != nil {
		return nil, err
	}
	params.AddNonZero64("message_id", config.MessageID)
	params.AddInterface("checklist", config.Checklist)
	err = params.CheckArgs("business_connection_id", "chat_id", "message_id", "checklist")
	if err != nil {
		return nil, err
	}
	return params, nil
}

func (config EditMessageChecklistConfig) method() string {
	return "editMessageChecklist"
}
