package tgbotapi

/**
Stickers
The following methods and objects allow your bot to handle stickers and sticker sets.

Sticker
This object represents a sticker.

Field	Type	Description
file_id	String	Identifier for this file, which can be used to download or reuse the file
file_unique_id	String	Unique identifier for this file, which is supposed to be the same over time and for different bots. Can't be used to download or reuse the file.
type	String	Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
width	Integer	Sticker width
height	Integer	Sticker height
is_animated	Boolean	True, if the sticker is animated
is_video	Boolean	True, if the sticker is a video sticker
thumbnail	PhotoSize	Optional. Sticker thumbnail in the .WEBP or .JPG format
emoji	String	Optional. Emoji associated with the sticker
set_name	String	Optional. Name of the sticker set to which the sticker belongs
premium_animation	File	Optional. For premium regular stickers, premium animation for the sticker
mask_position	MaskPosition	Optional. For mask stickers, the position where the mask should be placed
custom_emoji_id	String	Optional. For custom emoji stickers, unique identifier of the custom emoji
needs_repainting	True	Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
file_size	Integer	Optional. File size in bytes
StickerSet
This object represents a sticker set.

Field	Type	Description
name	String	Sticker set name
title	String	Sticker set title
sticker_type	String	Type of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
stickers	Array of Sticker	List of all set stickers
thumbnail	PhotoSize	Optional. Sticker set thumbnail in the .WEBP, .TGS, or .WEBM format
MaskPosition
This object describes the position on faces where a mask should be placed by default.

Field	Type	Description
point	String	The part of the face relative to which the mask should be placed. One of “forehead”, “eyes”, “mouth”, or “chin”.
x_shift	Float	Shift by X-axis measured in widths of the mask scaled to the face size, from left to right. For example, choosing -1.0 will place mask just to the left of the default mask position.
y_shift	Float	Shift by Y-axis measured in heights of the mask scaled to the face size, from top to bottom. For example, 1.0 will place the mask just below the default mask position.
scale	Float	Mask scaling coefficient. For example, 2.0 means double size.
InputSticker
This object describes a sticker to be added to a sticker set.

Field	Type	Description
sticker	String	The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new file using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
format	String	Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a .WEBM video
emoji_list	Array of String	List of 1-20 emoji associated with the sticker
mask_position	MaskPosition	Optional. Position where the mask should be placed on faces. For “mask” stickers only.
keywords	Array of String	Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
sendSticker
Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.

Parameter	Type	Required	Description
business_connection_id	String	Optional	Unique identifier of the business connection on behalf of which the message will be sent
chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
message_thread_id	Integer	Optional	Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
sticker	InputFile or String	Yes	Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files ». Video and animated stickers can't be sent via an HTTP URL.
emoji	String	Optional	Emoji associated with the sticker; only for just uploaded stickers
disable_notification	Boolean	Optional	Sends the message silently. Users will receive a notification with no sound.
protect_content	Boolean	Optional	Protects the contents of the sent message from forwarding and saving
allow_paid_broadcast	Boolean	Optional	Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
message_effect_id	String	Optional	Unique identifier of the message effect to be added to the message; for private chats only
reply_parameters	ReplyParameters	Optional	Description of the message to reply to
reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user
getStickerSet
Use this method to get a sticker set. On success, a StickerSet object is returned.

Parameter	Type	Required	Description
name	String	Yes	Name of the sticker set
getCustomEmojiStickers
Use this method to get information about custom emoji stickers by their identifiers. Returns an Array of Sticker objects.

Parameter	Type	Required	Description
custom_emoji_ids	Array of String	Yes	A JSON-serialized list of custom emoji identifiers. At most 200 custom emoji identifiers can be specified.
uploadStickerFile
Use this method to upload a file with a sticker for later use in the createNewStickerSet, addStickerToSet, or replaceStickerInSet methods (the file can be used multiple times). Returns the uploaded File on success.

Parameter	Type	Required	Description
user_id	Integer	Yes	User identifier of sticker file owner
sticker	InputFile	Yes	A file with the sticker in .WEBP, .PNG, .TGS, or .WEBM format. See https://core.telegram.org/stickers for technical requirements. More information on Sending Files »
sticker_format	String	Yes	Format of the sticker, must be one of “static”, “animated”, “video”
createNewStickerSet
Use this method to create a new sticker set owned by a user. The bot will be able to edit the sticker set thus created. Returns True on success.

Parameter	Type	Required	Description
user_id	Integer	Yes	User identifier of created sticker set owner
name	String	Yes	Short name of sticker set, to be used in t.me/addstickers/ URLs (e.g., animals). Can contain only English letters, digits and underscores. Must begin with a letter, can't contain consecutive underscores and must end in "_by_<bot_username>". <bot_username> is case insensitive. 1-64 characters.
title	String	Yes	Sticker set title, 1-64 characters
stickers	Array of InputSticker	Yes	A JSON-serialized list of 1-50 initial stickers to be added to the sticker set
sticker_type	String	Optional	Type of stickers in the set, pass “regular”, “mask”, or “custom_emoji”. By default, a regular sticker set is created.
needs_repainting	Boolean	Optional	Pass True if stickers in the sticker set must be repainted to the color of text when used in messages, the accent color if used as emoji status, white on chat photos, or another appropriate color based on context; for custom emoji sticker sets only
addStickerToSet
Use this method to add a new sticker to a set created by the bot. Emoji sticker sets can have up to 200 stickers. Other sticker sets can have up to 120 stickers. Returns True on success.

Parameter	Type	Required	Description
user_id	Integer	Yes	User identifier of sticker set owner
name	String	Yes	Sticker set name
sticker	InputSticker	Yes	A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set isn't changed.
setStickerPositionInSet
Use this method to move a sticker in a set created by the bot to a specific position. Returns True on success.

Parameter	Type	Required	Description
sticker	String	Yes	File identifier of the sticker
position	Integer	Yes	New sticker position in the set, zero-based
deleteStickerFromSet
Use this method to delete a sticker from a set created by the bot. Returns True on success.

Parameter	Type	Required	Description
sticker	String	Yes	File identifier of the sticker
replaceStickerInSet
Use this method to replace an existing sticker in a sticker set with a new one. The method is equivalent to calling deleteStickerFromSet, then addStickerToSet, then setStickerPositionInSet. Returns True on success.

Parameter	Type	Required	Description
user_id	Integer	Yes	User identifier of the sticker set owner
name	String	Yes	Sticker set name
old_sticker	String	Yes	File identifier of the replaced sticker
sticker	InputSticker	Yes	A JSON-serialized object with information about the added sticker. If exactly the same sticker had already been added to the set, then the set remains unchanged.
setStickerEmojiList
Use this method to change the list of emoji assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.

Parameter	Type	Required	Description
sticker	String	Yes	File identifier of the sticker
emoji_list	Array of String	Yes	A JSON-serialized list of 1-20 emoji associated with the sticker
setStickerKeywords
Use this method to change search keywords assigned to a regular or custom emoji sticker. The sticker must belong to a sticker set created by the bot. Returns True on success.

Parameter	Type	Required	Description
sticker	String	Yes	File identifier of the sticker
keywords	Array of String	Optional	A JSON-serialized list of 0-20 search keywords for the sticker with total length of up to 64 characters
setStickerMaskPosition
Use this method to change the mask position of a mask sticker. The sticker must belong to a sticker set that was created by the bot. Returns True on success.

Parameter	Type	Required	Description
sticker	String	Yes	File identifier of the sticker
mask_position	MaskPosition	Optional	A JSON-serialized object with the position where the mask should be placed on faces. Omit the parameter to remove the mask position.
setStickerSetTitle
Use this method to set the title of a created sticker set. Returns True on success.

Parameter	Type	Required	Description
name	String	Yes	Sticker set name
title	String	Yes	Sticker set title, 1-64 characters
setStickerSetThumbnail
Use this method to set the thumbnail of a regular or mask sticker set. The format of the thumbnail file must match the format of the stickers in the set. Returns True on success.

Parameter	Type	Required	Description
name	String	Yes	Sticker set name
user_id	Integer	Yes	User identifier of the sticker set owner
thumbnail	InputFile or String	Optional	A .WEBP or .PNG image with the thumbnail, must be up to 128 kilobytes in size and have a width and height of exactly 100px, or a .TGS animation with a thumbnail up to 32 kilobytes in size (see https://core.telegram.org/stickers#animation-requirements for animated sticker technical requirements), or a .WEBM video with the thumbnail up to 32 kilobytes in size; see https://core.telegram.org/stickers#video-requirements for video sticker technical requirements. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or upload a new one using multipart/form-data. More information on Sending Files ». Animated and video sticker set thumbnails can't be uploaded via HTTP URL. If omitted, then the thumbnail is dropped and the first sticker is used as the thumbnail.
format	String	Yes	Format of the thumbnail, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, or “video” for a .WEBM video
setCustomEmojiStickerSetThumbnail
Use this method to set the thumbnail of a custom emoji sticker set. Returns True on success.

Parameter	Type	Required	Description
name	String	Yes	Sticker set name
custom_emoji_id	String	Optional	Custom emoji identifier of a sticker from the sticker set; pass an empty string to drop the thumbnail and use the first sticker as the thumbnail.
deleteStickerSet
Use this method to delete a sticker set that was created by the bot. Returns True on success.

Parameter	Type	Required	Description
name	String	Yes	Sticker set name

InputSticker
This object describes a sticker to be added to a sticker set.

Field	Type	Description
sticker	String	The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new file using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
format	String	Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a .WEBM video
emoji_list	Array of String	List of 1-20 emoji associated with the sticker
mask_position	MaskPosition	Optional. Position where the mask should be placed on faces. For “mask” stickers only.
keywords	Array of String	Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
sendSticker
Use this method to send static .WEBP, animated .TGS, or video .WEBM stickers. On success, the sent Message is returned.

Parameter	Type	Required	Description
business_connection_id	String	Optional	Unique identifier of the business connection on behalf of which the message will be sent
chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
message_thread_id	Integer	Optional	Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
sticker	InputFile or String	Yes	Sticker to send. Pass a file_id as String to send a file that exists on the Telegram servers (recommended), pass an HTTP URL as a String for Telegram to get a .WEBP sticker from the Internet, or upload a new .WEBP, .TGS, or .WEBM sticker using multipart/form-data. More information on Sending Files ». Video and animated stickers can't be sent via an HTTP URL.
emoji	String	Optional	Emoji associated with the sticker; only for just uploaded stickers
disable_notification	Boolean	Optional	Sends the message silently. Users will receive a notification with no sound.
protect_content	Boolean	Optional	Protects the contents of the sent message from forwarding and saving
allow_paid_broadcast	Boolean	Optional	Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
message_effect_id	String	Optional	Unique identifier of the message effect to be added to the message; for private chats only
reply_parameters	ReplyParameters	Optional	Description of the message to reply to
reply_markup	InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply	Optional	Additional interface options. A JSON-serialized object for an inline keyboard, custom reply keyboard, instructions to remove a reply keyboard or to force a reply from the user

*/

// StickerSetConfig contains information about a GetStickerSet request.
type GetStickerSetConfig struct {
	Name string
}

func (config GetStickerSetConfig) method() string {
	return "getStickerSet"
}

func (config GetStickerSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("name", config.Name)
	return params, nil
}

// GetCustomEmojiStickersConfig contains information about a GetCustomEmojiStickers request.
type GetCustomEmojiStickersConfig struct {
	CustomEmojiIDs []string
}

func (config GetCustomEmojiStickersConfig) method() string {
	return "getCustomEmojiStickers"
}

func (config GetCustomEmojiStickersConfig) params() (Params, error) {
	params := Params{}
	params.AddInterface("custom_emoji_ids", config.CustomEmojiIDs)
	return params, nil
}

// UploadStickerFileConfig contains information about a UploadStickerFile request.
type UploadStickerFileConfig struct {
	UserID        int64
	Sticker       RequestFileData
	StickerFormat string
}

func (config UploadStickerFileConfig) method() string {
	return "uploadStickerFile"
}

func (config UploadStickerFileConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("sticker_format", config.StickerFormat)
	return params, nil
}

func (config UploadStickerFileConfig) files() []RequestFile {
	return []RequestFile{{
		Name: "sticker",
		Data: config.Sticker,
	}}
}

// InputSticker represents a sticker to be added to a sticker set.
type InputSticker struct {
	Sticker      RequestFileData
	Format       string
	EmojiList    []string
	MaskPosition *MaskPosition
	Keywords     []string
}

// CreateNewStickerSetConfig contains information about a CreateNewStickerSet request.
type CreateNewStickerSetConfig struct {
	UserID          int64
	Name            string
	Title           string
	Stickers        []InputSticker
	StickerType     string
	NeedsRepainting bool
}

func (config CreateNewStickerSetConfig) method() string {
	return "createNewStickerSet"
}

func (config CreateNewStickerSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("name", config.Name)
	params.AddNonEmpty("title", config.Title)
	params.AddInterface("stickers", config.Stickers)
	params.AddNonEmpty("sticker_type", config.StickerType)
	params.AddBool("needs_repainting", config.NeedsRepainting)
	return params, nil
}

// AddStickerToSetConfig contains information about an AddStickerToSet request.
type AddStickerToSetConfig struct {
	UserID  int64
	Name    string
	Sticker InputSticker
}

func (config AddStickerToSetConfig) method() string {
	return "addStickerToSet"
}

func (config AddStickerToSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("name", config.Name)
	params.AddInterface("sticker", config.Sticker)
	return params, nil
}

// SetStickerPositionInSetConfig contains information about a SetStickerPositionInSet request.
type SetStickerPositionInSetConfig struct {
	Sticker  string
	Position int
}

func (config SetStickerPositionInSetConfig) method() string {
	return "setStickerPositionInSet"
}

func (config SetStickerPositionInSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("sticker", config.Sticker)
	params.AddNonZero("position", config.Position)
	return params, nil
}

// DeleteStickerFromSetConfig contains information about a DeleteStickerFromSet request.
type DeleteStickerFromSetConfig struct {
	Sticker string
}

func (config DeleteStickerFromSetConfig) method() string {
	return "deleteStickerFromSet"
}

func (config DeleteStickerFromSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("sticker", config.Sticker)
	return params, nil
}

// ReplaceStickerInSetConfig contains information about a ReplaceStickerInSet request.
type ReplaceStickerInSetConfig struct {
	UserID     int64
	Name       string
	OldSticker string
	Sticker    InputSticker
}

func (config ReplaceStickerInSetConfig) method() string {
	return "replaceStickerInSet"
}

func (config ReplaceStickerInSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("name", config.Name)
	params.AddNonEmpty("old_sticker", config.OldSticker)
	params.AddInterface("sticker", config.Sticker)
	return params, nil
}

// SetStickerEmojiListConfig contains information about a SetStickerEmojiList request.
type SetStickerEmojiListConfig struct {
	Sticker   string
	EmojiList []string
}

func (config SetStickerEmojiListConfig) method() string {
	return "setStickerEmojiList"
}

func (config SetStickerEmojiListConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("sticker", config.Sticker)
	params.AddInterface("emoji_list", config.EmojiList)
	return params, nil
}

// SetStickerKeywordsConfig contains information about a SetStickerKeywords request.
type SetStickerKeywordsConfig struct {
	Sticker  string
	Keywords []string
}

func (config SetStickerKeywordsConfig) method() string {
	return "setStickerKeywords"
}

func (config SetStickerKeywordsConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("sticker", config.Sticker)
	params.AddInterface("keywords", config.Keywords)
	return params, nil
}

// SetStickerMaskPositionConfig contains information about a SetStickerMaskPosition request.
type SetStickerMaskPositionConfig struct {
	Sticker      string
	MaskPosition *MaskPosition
}

func (config SetStickerMaskPositionConfig) method() string {
	return "setStickerMaskPosition"
}

func (config SetStickerMaskPositionConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("sticker", config.Sticker)
	params.AddInterface("mask_position", config.MaskPosition)
	return params, nil
}

// SetStickerSetTitleConfig contains information about a SetStickerSetTitle request.
type SetStickerSetTitleConfig struct {
	Name  string
	Title string
}

func (config SetStickerSetTitleConfig) method() string {
	return "setStickerSetTitle"
}

func (config SetStickerSetTitleConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("name", config.Name)
	params.AddNonEmpty("title", config.Title)
	return params, nil
}

// SetStickerSetThumbnailConfig contains information about a SetStickerSetThumbnail request.
type SetStickerSetThumbnailConfig struct {
	Name      string
	UserID    int64
	Thumbnail RequestFileData
	Format    string
}

func (config SetStickerSetThumbnailConfig) method() string {
	return "setStickerSetThumbnail"
}

func (config SetStickerSetThumbnailConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("name", config.Name)
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("format", config.Format)
	return params, nil
}

func (config SetStickerSetThumbnailConfig) files() []RequestFile {
	return []RequestFile{{
		Name: "thumbnail",
		Data: config.Thumbnail,
	}}
}

// SetCustomEmojiStickerSetThumbnailConfig contains information about a SetCustomEmojiStickerSetThumbnail request.
type SetCustomEmojiStickerSetThumbnailConfig struct {
	Name          string
	CustomEmojiID string
}

func (config SetCustomEmojiStickerSetThumbnailConfig) method() string {
	return "setCustomEmojiStickerSetThumbnail"
}

func (config SetCustomEmojiStickerSetThumbnailConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("name", config.Name)
	params.AddNonEmpty("custom_emoji_id", config.CustomEmojiID)
	return params, nil
}

// DeleteStickerSetConfig contains information about a DeleteStickerSet request.
type DeleteStickerSetConfig struct {
	Name string
}

func (config DeleteStickerSetConfig) method() string {
	return "deleteStickerSet"
}

func (config DeleteStickerSetConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("name", config.Name)
	return params, nil
}

type SendStickerConfig struct {
	BusinessConnectionID string
	ChatID               int64
	MessageThreadID      int64
	Sticker              RequestFileData
	Emoji                string
	DisableNotification  bool
	ProtectContent       bool
	AllowPaidBroadcast   bool
	MessageEffectID      string
	ReplyParameters      *ReplyParameters
	ReplyMarkup          *InlineKeyboardMarkup
}

func (config SendStickerConfig) method() string {
	return "sendSticker"
}

func (config SendStickerConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonZero64("chat_id", config.ChatID)
	params.AddNonZero64("message_thread_id", config.MessageThreadID)
	params.AddInterface("sticker", config.Sticker)
	params.AddNonEmpty("emoji", config.Emoji)
	params.AddBool("disable_notification", config.DisableNotification)
	params.AddBool("protect_content", config.ProtectContent)
	params.AddBool("allow_paid_broadcast", config.AllowPaidBroadcast)
	params.AddNonEmpty("message_effect_id", config.MessageEffectID)
	params.AddInterface("reply_parameters", config.ReplyParameters)
	params.AddInterface("reply_markup", config.ReplyMarkup)
	return params, nil
}

/*
InputSticker
This object describes a sticker to be added to a sticker set.

Field	Type	Description
sticker	String	The added sticker. Pass a file_id as a String to send a file that already exists on the Telegram servers, pass an HTTP URL as a String for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new file using multipart/form-data under <file_attach_name> name. Animated and video stickers can't be uploaded via HTTP URL. More information on Sending Files »
format	String	Format of the added sticker, must be one of “static” for a .WEBP or .PNG image, “animated” for a .TGS animation, “video” for a .WEBM video
emoji_list	Array of String	List of 1-20 emoji associated with the sticker
mask_position	MaskPosition	Optional. Position where the mask should be placed on faces. For “mask” stickers only.
keywords	Array of String	Optional. List of 0-20 search keywords for the sticker with total length of up to 64 characters. For “regular” and “custom_emoji” stickers only.
*/
type InputStickerConfig struct {
	Sticker      string
	Format       string
	EmojiList    []string
	MaskPosition *MaskPosition
	Keywords     []string
}

func (config InputStickerConfig) method() string {
	return "inputSticker"
}

func (config InputStickerConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("sticker", config.Sticker)
	params.AddNonEmpty("format", config.Format)
	params.AddInterface("emoji_list", config.EmojiList)
	params.AddInterface("mask_position", config.MaskPosition)
	params.AddInterface("keywords", config.Keywords)
	return params, nil
}

type SetChatStickerSetConfig struct {
	ChatID          int64
	ChannelUsername string
	StickerSetName  string
}

func (config SetChatStickerSetConfig) method() string {
	return "setChatStickerSet"
}

func (config SetChatStickerSetConfig) params() (Params, error) {
	params := Params{}
	params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
	params.AddNonEmpty("sticker_set_name", config.StickerSetName)
	return params, nil
}

type DeleteChatStickerSetConfig struct {
	ChatID          int64
	ChannelUsername string
}

func (config DeleteChatStickerSetConfig) method() string {
	return "deleteChatStickerSet"
}

func (config DeleteChatStickerSetConfig) params() (Params, error) {
	params := Params{}
	params.AddFirstValid("chat_id", config.ChatID, config.ChannelUsername)
	return params, nil
}
