package tgbotapi

/**
InlineQueryResult
This object represents one result of an inline query. Telegram clients currently support results of the following 20 types:

InlineQueryResultCachedAudio
InlineQueryResultCachedDocument
InlineQueryResultCachedGif
InlineQueryResultCachedMpeg4Gif
InlineQueryResultCachedPhoto
InlineQueryResultCachedSticker
InlineQueryResultCachedVideo
InlineQueryResultCachedVoice
InlineQueryResultArticle
InlineQueryResultAudio
InlineQueryResultContact
InlineQueryResultGame
InlineQueryResultDocument
InlineQueryResultGif
InlineQueryResultLocation
InlineQueryResultMpeg4Gif
InlineQueryResultPhoto
InlineQueryResultVenue
InlineQueryResultVideo
InlineQueryResultVoice
*/

// --- InputTextMessageContent ---
type InputTextMessageContent struct {
	MessageText        string              `json:"message_text"`
	ParseMode          string              `json:"parse_mode,omitempty"`
	Entities           []MessageEntity     `json:"entities,omitempty"`
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
}

// --- InputLocationMessageContent ---
type InputLocationMessageContent struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	HorizontalAccuracy   float64 `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int     `json:"live_period,omitempty"`
	Heading              int     `json:"heading,omitempty"`
	ProximityAlertRadius int     `json:"proximity_alert_radius,omitempty"`
}

// --- InputVenueMessageContent ---
type InputVenueMessageContent struct {
	Latitude        float64 `json:"latitude"`
	Longitude       float64 `json:"longitude"`
	Title           string  `json:"title"`
	Address         string  `json:"address"`
	FoursquareID    string  `json:"foursquare_id,omitempty"`
	FoursquareType  string  `json:"foursquare_type,omitempty"`
	GooglePlaceID   string  `json:"google_place_id,omitempty"`
	GooglePlaceType string  `json:"google_place_type,omitempty"`
}

// --- InputContactMessageContent ---
type InputContactMessageContent struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// --- InputInvoiceMessageContent ---
type InputInvoiceMessageContent struct {
	Title                     string         `json:"title"`
	Description               string         `json:"description"`
	Payload                   string         `json:"payload"`
	ProviderToken             string         `json:"provider_token,omitempty"`
	Currency                  string         `json:"currency"`
	Prices                    []LabeledPrice `json:"prices"`
	MaxTipAmount              int            `json:"max_tip_amount,omitempty"`
	SuggestedTipAmounts       []int          `json:"suggested_tip_amounts,omitempty"`
	ProviderData              string         `json:"provider_data,omitempty"`
	PhotoURL                  string         `json:"photo_url,omitempty"`
	PhotoSize                 int            `json:"photo_size,omitempty"`
	PhotoWidth                int            `json:"photo_width,omitempty"`
	PhotoHeight               int            `json:"photo_height,omitempty"`
	NeedName                  bool           `json:"need_name,omitempty"`
	NeedPhoneNumber           bool           `json:"need_phone_number,omitempty"`
	NeedEmail                 bool           `json:"need_email,omitempty"`
	NeedShippingAddress       bool           `json:"need_shipping_address,omitempty"`
	SendPhoneNumberToProvider bool           `json:"send_phone_number_to_provider,omitempty"`
	SendEmailToProvider       bool           `json:"send_email_to_provider,omitempty"`
	IsFlexible                bool           `json:"is_flexible,omitempty"`
}

type InlineQueryResult interface {
	GetType() string
	GetID() string
}

type BaseInlineQueryResult struct {
	Type string `json:"type"`
	ID   string `json:"id"`
}

func (b BaseInlineQueryResult) GetType() string {
	return b.Type
}

func (b BaseInlineQueryResult) GetID() string {
	return b.ID
}

type InlineQueryResultArticle struct {
	BaseInlineQueryResult
	Title               string                `json:"title"`
	InputMessageContent interface{}           `json:"input_message_content"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	URL                 string                `json:"url,omitempty"`
	Description         string                `json:"description,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

// --- Photo ---
type InlineQueryResultPhoto struct {
	BaseInlineQueryResult
	PhotoURL            string                `json:"photo_url"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	PhotoWidth          int                   `json:"photo_width,omitempty"`
	PhotoHeight         int                   `json:"photo_height,omitempty"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- Gif ---
type InlineQueryResultGif struct {
	BaseInlineQueryResult
	GifURL              string                `json:"gif_url"`
	GifWidth            int                   `json:"gif_width,omitempty"`
	GifHeight           int                   `json:"gif_height,omitempty"`
	GifDuration         int                   `json:"gif_duration,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	ThumbnailMimeType   string                `json:"thumbnail_mime_type,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- Mpeg4Gif ---
type InlineQueryResultMpeg4Gif struct {
	BaseInlineQueryResult
	Mpeg4URL            string                `json:"mpeg4_url"`
	Mpeg4Width          int                   `json:"mpeg4_width,omitempty"`
	Mpeg4Height         int                   `json:"mpeg4_height,omitempty"`
	Mpeg4Duration       int                   `json:"mpeg4_duration,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	ThumbnailMimeType   string                `json:"thumbnail_mime_type,omitempty"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- Video ---
type InlineQueryResultVideo struct {
	BaseInlineQueryResult
	VideoURL            string                `json:"video_url"`
	MimeType            string                `json:"mime_type"`
	ThumbnailURL        string                `json:"thumbnail_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	VideoWidth          int                   `json:"video_width,omitempty"`
	VideoHeight         int                   `json:"video_height,omitempty"`
	VideoDuration       int                   `json:"video_duration,omitempty"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- Audio ---
type InlineQueryResultAudio struct {
	BaseInlineQueryResult
	AudioURL            string                `json:"audio_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	Performer           string                `json:"performer,omitempty"`
	AudioDuration       int                   `json:"audio_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- Voice ---
type InlineQueryResultVoice struct {
	BaseInlineQueryResult
	VoiceURL            string                `json:"voice_url"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	VoiceDuration       int                   `json:"voice_duration,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- Document ---
type InlineQueryResultDocument struct {
	BaseInlineQueryResult
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	DocumentURL         string                `json:"document_url"`
	MimeType            string                `json:"mime_type"`
	Description         string                `json:"description,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

// --- Location ---
type InlineQueryResultLocation struct {
	BaseInlineQueryResult
	Latitude             float64               `json:"latitude"`
	Longitude            float64               `json:"longitude"`
	Title                string                `json:"title"`
	HorizontalAccuracy   float64               `json:"horizontal_accuracy,omitempty"`
	LivePeriod           int                   `json:"live_period,omitempty"`
	Heading              int                   `json:"heading,omitempty"`
	ProximityAlertRadius int                   `json:"proximity_alert_radius,omitempty"`
	ReplyMarkup          *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent  interface{}           `json:"input_message_content,omitempty"`
	ThumbnailURL         string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth       int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight      int                   `json:"thumbnail_height,omitempty"`
}

// --- Venue ---
type InlineQueryResultVenue struct {
	BaseInlineQueryResult
	Latitude            float64               `json:"latitude"`
	Longitude           float64               `json:"longitude"`
	Title               string                `json:"title"`
	Address             string                `json:"address"`
	FoursquareID        string                `json:"foursquare_id,omitempty"`
	FoursquareType      string                `json:"foursquare_type,omitempty"`
	GooglePlaceID       string                `json:"google_place_id,omitempty"`
	GooglePlaceType     string                `json:"google_place_type,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

// --- Contact ---
type InlineQueryResultContact struct {
	BaseInlineQueryResult
	PhoneNumber         string                `json:"phone_number"`
	FirstName           string                `json:"first_name"`
	LastName            string                `json:"last_name,omitempty"`
	VCard               string                `json:"vcard,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
	ThumbnailURL        string                `json:"thumbnail_url,omitempty"`
	ThumbnailWidth      int                   `json:"thumbnail_width,omitempty"`
	ThumbnailHeight     int                   `json:"thumbnail_height,omitempty"`
}

// --- Game ---
type InlineQueryResultGame struct {
	BaseInlineQueryResult
	GameShortName string                `json:"game_short_name"`
	ReplyMarkup   *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
}

// --- CachedPhoto ---
type InlineQueryResultCachedPhoto struct {
	BaseInlineQueryResult
	PhotoFileID         string                `json:"photo_file_id"`
	Title               string                `json:"title,omitempty"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedGif ---
type InlineQueryResultCachedGif struct {
	BaseInlineQueryResult
	GifFileID           string                `json:"gif_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedMpeg4Gif ---
type InlineQueryResultCachedMpeg4Gif struct {
	BaseInlineQueryResult
	Mpeg4FileID         string                `json:"mpeg4_file_id"`
	Title               string                `json:"title,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedVideo ---
type InlineQueryResultCachedVideo struct {
	BaseInlineQueryResult
	VideoFileID         string                `json:"video_file_id"`
	Title               string                `json:"title"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ShowCaptionAbove    bool                  `json:"show_caption_above_media,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedAudio ---
type InlineQueryResultCachedAudio struct {
	BaseInlineQueryResult
	AudioFileID         string                `json:"audio_file_id"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedVoice ---
type InlineQueryResultCachedVoice struct {
	BaseInlineQueryResult
	VoiceFileID         string                `json:"voice_file_id"`
	Title               string                `json:"title"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedSticker ---
type InlineQueryResultCachedSticker struct {
	BaseInlineQueryResult
	StickerFileID       string                `json:"sticker_file_id"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

// --- CachedDocument ---
type InlineQueryResultCachedDocument struct {
	BaseInlineQueryResult
	Title               string                `json:"title"`
	DocumentFileID      string                `json:"document_file_id"`
	Description         string                `json:"description,omitempty"`
	Caption             string                `json:"caption,omitempty"`
	ParseMode           string                `json:"parse_mode,omitempty"`
	CaptionEntities     []MessageEntity       `json:"caption_entities,omitempty"`
	ReplyMarkup         *InlineKeyboardMarkup `json:"reply_markup,omitempty"`
	InputMessageContent interface{}           `json:"input_message_content,omitempty"`
}

/*
savePreparedInlineMessage
Stores a message that can be sent by a user of a Mini App. Returns a PreparedInlineMessage object.

Parameter	Type	Required	Description
user_id	Integer	Yes	Unique identifier of the target user that can use the prepared message
result	InlineQueryResult	Yes	A JSON-serialized object describing the message to be sent
allow_user_chats	Boolean	Optional	Pass True if the message can be sent to private chats with users
allow_bot_chats	Boolean	Optional	Pass True if the message can be sent to private chats with bots
allow_group_chats	Boolean	Optional	Pass True if the message can be sent to group and supergroup chats
allow_channel_chats	Boolean	Optional	Pass True if the message can be sent to channel chats
PreparedInlineMessage
Describes an inline message to be sent by a user of a Mini App.

Field	Type	Description
id	String	Unique identifier of the prepared message
expiration_date	Integer	Expiration date of the prepared message, in Unix time. Expired prepared messages can no longer be used
*/

type PreparedInlineMessage struct {
	ID             string `json:"id"`
	ExpirationDate int    `json:"expiration_date"`
}

type SavePreparedInlineMessageConfig struct {
	UserID            int64       `json:"user_id"`
	Result            interface{} `json:"result"`
	AllowUserChats    bool        `json:"allow_user_chats"`
	AllowBotChats     bool        `json:"allow_bot_chats"`
	AllowGroupChats   bool        `json:"allow_group_chats"`
	AllowChannelChats bool        `json:"allow_channel_chats"`
}

func (config SavePreparedInlineMessageConfig) method() string {
	return "savePreparedInlineMessage"
}

func (config SavePreparedInlineMessageConfig) params() (Params, error) {
	params := make(Params)

	params.AddNonZero64("user_id", config.UserID)
	params.AddInterface("result", config.Result)
	params.AddBool("allow_user_chats", config.AllowUserChats)
	params.AddBool("allow_bot_chats", config.AllowBotChats)
	params.AddBool("allow_group_chats", config.AllowGroupChats)
	params.AddBool("allow_channel_chats", config.AllowChannelChats)

	return params, nil
}
