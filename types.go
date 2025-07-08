package tgbotapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"
)

// APIResponse is a response from the Telegram API with the result
// stored raw.
type APIResponse struct {
	Ok          bool                `json:"ok"`
	Result      json.RawMessage     `json:"result,omitempty"`
	ErrorCode   int                 `json:"error_code,omitempty"`
	Description string              `json:"description,omitempty"`
	Parameters  *ResponseParameters `json:"parameters,omitempty"`
}

// Error is an error containing extra information returned by the Telegram API.
type Error struct {
	Code    int
	Message string
	ResponseParameters
}

// Error message string.
func (e Error) Error() string {
	return e.Message
}

// Update is an update response, from GetUpdates.
type Update struct {
	// UpdateID is the update's unique identifier.
	// Update identifiers start from a certain positive number and increase
	// sequentially.
	// This ID becomes especially handy if you're using Webhooks,
	// since it allows you to ignore repeated updates or to restore
	// the correct update sequence, should they get out of order.
	// If there are no new updates for at least a week, then identifier
	// of the next update will be chosen randomly instead of sequentially.
	UpdateID int `json:"update_id"`
	// Message new incoming message of any kind — text, photo, sticker, etc.
	//
	// optional
	Message *Message `json:"message,omitempty"`
	// EditedMessage new version of a message that is known to the bot and was
	// edited
	//
	// optional
	EditedMessage *Message `json:"edited_message,omitempty"`
	// ChannelPost new version of a message that is known to the bot and was
	// edited
	//
	// optional
	ChannelPost *Message `json:"channel_post,omitempty"`
	// EditedChannelPost new incoming channel post of any kind — text, photo,
	// sticker, etc.
	//
	// optional
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"`
	// The bot was connected to or disconnected from a business account,
	// or a user edited an existing connection with the bot
	//
	// optional
	BusinessConnection *BusinessConnection `json:"business_connection,omitempty"`
	// New message from a connected business account
	//
	// optional
	BusinessMessage *Message `json:"business_message,omitempty"`
	// New version of a message from a connected business account
	//
	// optional
	EditedBusinessMessage *Message `json:"edited_business_message,omitempty"`
	// Messages were deleted from a connected business account
	//
	// optional
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"`
	// A reaction to a message was changed by a user. The bot must be an administrator in the chat
	// and must explicitly specify "message_reaction" in the list of allowed_updates to receive these updates.
	// The update isn't received for reactions set by bots.
	//
	// optional
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`
	// Reactions to a message with anonymous reactions were changed. The bot must be an administrator
	// in the chat and must explicitly specify "message_reaction_count" in the list of
	// allowed_updates to receive these updates. The updates are grouped and can be sent with delay up to a few minutes.
	//
	// optional
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`
	// InlineQuery new incoming inline query
	//
	// optional
	InlineQuery *InlineQuery `json:"inline_query,omitempty"`
	// ChosenInlineResult is the result of an inline query
	// that was chosen by a user and sent to their chat partner.
	// Please see our documentation on the feedback collecting
	// for details on how to enable these updates for your bot.
	//
	// optional
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"`
	// CallbackQuery new incoming callback query
	//
	// optional
	CallbackQuery *CallbackQuery `json:"callback_query,omitempty"`
	// ShippingQuery new incoming shipping query. Only for invoices with
	// flexible price
	//
	// optional
	ShippingQuery *ShippingQuery `json:"shipping_query,omitempty"`
	// PreCheckoutQuery new incoming pre-checkout query. Contains full
	// information about checkout
	//
	// optional
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"`

	//purchased_paid_media
	PaidMediaPurchased *PaidMediaPurchased `json:"paid_media_purchased,omitempty"`

	// Pool new poll state. Bots receive only updates about stopped polls and
	// polls, which are sent by the bot
	//
	// optional
	Poll *Poll `json:"poll,omitempty"`
	// PollAnswer user changed their answer in a non-anonymous poll. Bots
	// receive new votes only in polls that were sent by the bot itself.
	//
	// optional
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"`
	// MyChatMember is the bot's chat member status was updated in a chat. For
	// private chats, this update is received only when the bot is blocked or
	// unblocked by the user.
	//
	// optional
	MyChatMember *ChatMemberUpdated `json:"my_chat_member,omitempty"`
	// ChatMember is a chat member's status was updated in a chat. The bot must
	// be an administrator in the chat and must explicitly specify "chat_member"
	// in the list of allowed_updates to receive these updates.
	//
	// optional
	ChatMember *ChatMemberUpdated `json:"chat_member,omitempty"`
	// ChatJoinRequest is a request to join the chat has been sent. The bot must
	// have the can_invite_users administrator right in the chat to receive
	// these updates.
	//
	// optional
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"`
	// A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	//
	// optional
	ChatBoost *ChatBoostUpdated `json:"chat_boost,omitempty"`
	// A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.
	//
	// optional
	RemovedChatBoost *ChatBoostRemoved `json:"removed_chat_boost,omitempty"`

	//purchased_paid_media
	//Optional. A user purchased paid media with a non-empty payload sent by the bot in a non-channel chat
	PurchasedPaidMedia *PaidMediaPurchased `json:"purchased_paid_media,omitempty"`
}

// SentFrom returns the user who sent an update. Can be nil, if Telegram did not provide information
// about the user in the update object.
func (u *Update) SentFrom() *User {
	switch {
	case u.Message != nil:
		return u.Message.From
	case u.EditedMessage != nil:
		return u.EditedMessage.From
	case u.ChannelPost != nil:
		return u.ChannelPost.From
	case u.EditedChannelPost != nil:
		return u.EditedChannelPost.From
	case u.BusinessConnection != nil:
		return u.BusinessConnection.User
	case u.BusinessMessage != nil:
		return u.BusinessMessage.From
	case u.EditedBusinessMessage != nil:
		return u.EditedBusinessMessage.From
	case u.MessageReaction != nil:
		return u.MessageReaction.User
	case u.InlineQuery != nil:
		return u.InlineQuery.From
	case u.ChosenInlineResult != nil:
		return u.ChosenInlineResult.From
	case u.CallbackQuery != nil:
		return u.CallbackQuery.From
	case u.ShippingQuery != nil:
		return u.ShippingQuery.From
	case u.PreCheckoutQuery != nil:
		return u.PreCheckoutQuery.From
	case u.PollAnswer != nil:
		return u.PollAnswer.User
	case u.ChatMember != nil:
		return u.ChatMember.From
	case u.ChatJoinRequest != nil:
		return u.ChatJoinRequest.From
	default:
		return nil
	}
}

// CallbackData returns the callback query data, if it exists.
func (u *Update) CallbackData() string {
	if u.CallbackQuery != nil {
		return u.CallbackQuery.Data
	}
	return ""
}

// FromChat returns the chat where an update occurred.
func (u *Update) FromChat() *Chat {
	switch {
	case u.Message != nil:
		return u.Message.Chat
	case u.EditedMessage != nil:
		return u.EditedMessage.Chat
	case u.ChannelPost != nil:
		return u.ChannelPost.Chat
	case u.EditedChannelPost != nil:
		return u.EditedChannelPost.Chat
	case u.BusinessMessage != nil:
		return u.BusinessMessage.Chat
	case u.EditedBusinessMessage != nil:
		return u.EditedBusinessMessage.Chat
	case u.DeletedBusinessMessages != nil:
		return u.DeletedBusinessMessages.Chat
	case u.MessageReaction != nil:
		return u.MessageReaction.Chat
	case u.MessageReactionCount != nil:
		return u.MessageReactionCount.Chat
	case u.CallbackQuery != nil:
		return u.CallbackQuery.Message.Chat
	case u.MyChatMember != nil:
		return u.MyChatMember.Chat
	case u.ChatMember != nil:
		return u.ChatMember.Chat
	case u.ChatJoinRequest != nil:
		return u.ChatJoinRequest.Chat
	case u.ChatBoost != nil:
		return u.ChatBoost.Chat
	case u.RemovedChatBoost != nil:
		return u.RemovedChatBoost.Chat
	default:
		return nil
	}
}

// UpdatesChannel is the channel for getting updates.
type UpdatesChannel <-chan Update

// Clear discards all unprocessed incoming updates.
func (ch UpdatesChannel) Clear() {
	for len(ch) != 0 {
		<-ch
	}
}

// User represents a Telegram user or bot.
type User struct {
	// ID is a unique identifier for this user or bot
	ID int64 `json:"id"`
	// IsBot true, if this user is a bot
	//
	// optional
	IsBot bool `json:"is_bot,omitempty"`
	// FirstName user's or bot's first name
	FirstName string `json:"first_name"`
	// LastName user's or bot's last name
	//
	// optional
	LastName string `json:"last_name,omitempty"`
	// UserName user's or bot's username
	//
	// optional
	UserName string `json:"username,omitempty"`
	// LanguageCode IETF language tag of the user's language
	// more info: https://en.wikipedia.org/wiki/IETF_language_tag
	//
	// optional
	LanguageCode string `json:"language_code,omitempty"`
	// IsPremium true, if user has Telegram Premium
	//
	// optional
	IsPremium bool `json:"is_premium,omitempty"`
	// True, if this user added the bot to the attachment menu
	//
	// optional
	AddedToAttachmentMenu bool `json:"added_to_attachment_menu,omitempty"`
	// CanJoinGroups is true, if the bot can be invited to groups.
	// Returned only in getMe.
	//
	// optional
	CanJoinGroups bool `json:"can_join_groups,omitempty"`
	// CanReadAllGroupMessages is true, if privacy mode is disabled for the bot.
	// Returned only in getMe.
	//
	// optional
	CanReadAllGroupMessages bool `json:"can_read_all_group_messages,omitempty"`
	// SupportsInlineQueries is true, if the bot supports inline queries.
	// Returned only in getMe.
	//
	// optional
	SupportsInlineQueries bool `json:"supports_inline_queries,omitempty"`
	// True, if the bot can be connected to a Telegram Business account to receive its messages. Returned only in getMe.
	//
	// optional
	CanConnectToBusiness bool `json:"can_connect_to_business,omitempty"`
	// rue, if the bot has a main Web App. Returned only in getMe.
	//
	// optional
	HasMainWebApp bool `json:"has_main_web_app,omitempty"`
}

// String displays a simple text version of a user.
//
// It is normally a user's username, but falls back to a first/last
// name as available.
func (u *User) String() string {
	if u == nil {
		return ""
	}
	if u.UserName != "" {
		return u.UserName
	}

	name := u.FirstName
	if u.LastName != "" {
		name += " " + u.LastName
	}

	return name
}

// Chat represents a chat.
type Chat struct {
	// ID is a unique identifier for this chat
	ID int64 `json:"id"`
	// Type of chat, can be either “private”, “group”, “supergroup” or “channel”
	Type string `json:"type"`
	// Title for supergroups, channels and group chats
	//
	// optional
	Title string `json:"title,omitempty"`
	// UserName for private chats, supergroups and channels if available
	//
	// optional
	UserName string `json:"username,omitempty"`
	// FirstName of the other party in a private chat
	//
	// optional
	FirstName string `json:"first_name,omitempty"`
	// LastName of the other party in a private chat
	//
	// optional
	LastName string `json:"last_name,omitempty"`
	// True, if the supergroup chat is a forum (has topics enabled)
	//
	// optional
	IsForum bool `json:"is_forum"`
	//// Photo is a chat photo
	//Photo *ChatPhoto `json:"photo"`
	//// Bio is the bio of the other party in a private chat. Returned only in
	//// getChat
	////
	//// optional
	//Bio string `json:"bio,omitempty"`
	//// HasPrivateForwards is true if privacy settings of the other party in the
	//// private chat allows to use tg://user?id=<user_id> links only in chats
	//// with the user. Returned only in getChat.
	////
	//// optional
	//HasPrivateForwards bool `json:"has_private_forwards,omitempty"`
	//// Description for groups, supergroups and channel chats
	////
	//// optional
	//Description string `json:"description,omitempty"`
	//// InviteLink is a chat invite link, for groups, supergroups and channel chats.
	//// Each administrator in a chat generates their own invite links,
	//// so the bot must first generate the link using exportChatInviteLink
	////
	//// optional
	//InviteLink string `json:"invite_link,omitempty"`
	//// PinnedMessage is the pinned message, for groups, supergroups and channels
	////
	//// optional
	//PinnedMessage *Message `json:"pinned_message,omitempty"`
	//// Permissions are default chat member permissions, for groups and
	//// supergroups. Returned only in getChat.
	////
	//// optional
	//Permissions *ChatPermissions `json:"permissions,omitempty"`
	//// SlowModeDelay is for supergroups, the minimum allowed delay between
	//// consecutive messages sent by each unprivileged user. Returned only in
	//// getChat.
	////
	//// optional
	//SlowModeDelay int `json:"slow_mode_delay,omitempty"`
	//// MessageAutoDeleteTime is the time after which all messages sent to the
	//// chat will be automatically deleted; in seconds. Returned only in getChat.
	////
	//// optional
	//MessageAutoDeleteTime int `json:"message_auto_delete_time,omitempty"`
	//// HasProtectedContent is true if messages from the chat can't be forwarded
	//// to other chats. Returned only in getChat.
	////
	//// optional
	//HasProtectedContent bool `json:"has_protected_content,omitempty"`
	//// StickerSetName is for supergroups, name of group sticker set.Returned
	//// only in getChat.
	////
	//// optional
	//StickerSetName string `json:"sticker_set_name,omitempty"`
	//// CanSetStickerSet is true, if the bot can change the group sticker set.
	//// Returned only in getChat.
	////
	//// optional
	//CanSetStickerSet bool `json:"can_set_sticker_set,omitempty"`
	//// LinkedChatID is a unique identifier for the linked chat, i.e. the
	//// discussion group identifier for a channel and vice versa; for supergroups
	//// and channel chats.
	////
	//// optional
	//LinkedChatID int64 `json:"linked_chat_id,omitempty"`
	//// Location is for supergroups, the location to which the supergroup is
	//// connected. Returned only in getChat.
	////
	//// optional
	//Location *ChatLocation `json:"location,omitempty"`
}

// IsPrivate returns if the Chat is a private conversation.
func (c Chat) IsPrivate() bool {
	return c.Type == "private"
}

// IsGroup returns if the Chat is a group.
func (c Chat) IsGroup() bool {
	return c.Type == "group"
}

// IsSuperGroup returns if the Chat is a supergroup.
func (c Chat) IsSuperGroup() bool {
	return c.Type == "supergroup"
}

// IsChannel returns if the Chat is a channel.
func (c Chat) IsChannel() bool {
	return c.Type == "channel"
}

// ChatConfig returns a ChatConfig struct for chat related methods.
func (c Chat) ChatConfig() ChatConfig {
	return ChatConfig{ChatID: c.ID}
}

// Message represents a message.
type Message struct {
	// MessageID is a unique message identifier inside this chat
	MessageID int `json:"message_id"`
	// Unique identifier of a message thread to which the message belongs; for supergroups only
	MessageThreadID int `json:"message_thread_id"`
	// From is a sender, empty for messages sent to channels;
	//
	// optional
	From *User `json:"from,omitempty"`
	// SenderChat is the sender of the message, sent on behalf of a chat. The
	// channel itself for channel messages. The supergroup itself for messages
	// from anonymous group administrators. The linked channel for messages
	// automatically forwarded to the discussion group
	//
	// optional
	SenderChat *Chat `json:"sender_chat,omitempty"`
	// if the sender of the message boosted the chat, the number of boosts added by the user
	//
	// optional
	SenderBoostCount int `json:"sender_boost_count"`
	// Date of the message was sent in Unix time
	Date int `json:"date"`
	// Unique identifier of the business connection from which the message was received.
	// If non-empty, the message belongs to a chat of the corresponding business account that is
	// independent from any potential bot chat which might share the same identifier.
	//
	// optional
	BusinessConnectionId string `json:"business_connection_id"`
	// Chat is the conversation the message belongs to
	Chat *Chat `json:"chat"`
	// Information about the original message for forwarded messages
	//
	// optional
	ForwardOrigin *MessageOrigin `json:"forward_origin,omitempty"`
	// True, if the message is sent to a forum topic
	//
	// optional
	IsTopicMessage bool `json:"is_topic_message,omitempty"`
	// True, if the message is a channel post that was automatically forwarded to the connected discussion group
	//
	// optional
	IsAutomaticForward bool `json:"is_automatic_forward,omitempty"`

	// Optional. Information about the message that is being replied to, which may come from another chat or forum topic
	// optional
	ExternalReplyInfo *ExternalReplyInfo `json:"external_reply,omitempty"`

	// ReplyToMessage for replies, the original message.
	// Note that the Message object in this field will not contain further ReplyToMessage fields
	// even if it itself is a reply;
	//
	// optional
	ReplyToMessage *Message `json:"reply_to_message,omitempty"`
	// Information about the message that is being replied to, which may come from another chat or forum topic

	// optional
	Quote *TextQuote `json:"quote,omitempty"`
	// For replies to a story, the original story
	//
	// optional
	ReplyToStory *Story `json:"reply_to_story,omitempty"`
	// ViaBot through which the message was sent;
	//
	// optional
	ViaBot *User `json:"via_bot,omitempty"`
	// EditDate of the message was last edited in Unix time;
	//
	// optional
	EditDate int `json:"edit_date,omitempty"`
	// HasProtectedContent is true if the message can't be forwarded.
	//
	// optional
	HasProtectedContent bool `json:"has_protected_content,omitempty"`
	// True, if the message was sent by an implicit action, for example, as an away or a greeting business message,
	// or as a scheduled message
	//
	// optional
	IsFromOffline bool `json:"is_from_offline,omitempty"`
	// MediaGroupID is the unique identifier of a media message group this message belongs to;
	//
	// optional
	MediaGroupID string `json:"media_group_id,omitempty"`
	// AuthorSignature is the signature of the post author for messages in channels;
	//
	// optional
	AuthorSignature string `json:"author_signature,omitempty"`
	// Text is for text messages, the actual UTF-8 text of the message, 0-4096 characters;
	//
	// optional
	Text string `json:"text,omitempty"`
	// Entities are for text messages, special entities like usernames,
	// URLs, bot commands, etc. that appear in the text;
	//
	// optional
	Entities []MessageEntity `json:"entities,omitempty"`
	// Options used for link preview generation for the message, if it is a text message and link preview options were changed
	//
	// optional
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options,omitempty"`
	// Unique identifier of the message effect added to the message
	//
	// optional
	EffectId string `json:"effect_id,omitempty"`
	// Animation message is an animation, information about the animation.
	// For backward compatibility, when this field is set, the document field will also be set;
	//
	// optional
	Animation *Animation `json:"animation,omitempty"`
	// Audio message is an audio file, information about the file;
	//
	// optional
	Audio *Audio `json:"audio,omitempty"`
	// Document message is a general file, information about the file;
	//
	// optional
	Document *Document `json:"document,omitempty"`
	// Message contains paid media; information about the paid media
	//
	// optional
	PaidMedia *PaidMediaInfo `json:"paid_media,omitempty"`
	// Photo message is a photo, available sizes of the photo;
	//
	// optional
	Photo []PhotoSize `json:"photo,omitempty"`
	// Sticker message is a sticker, information about the sticker;
	//
	// optional
	Sticker *Sticker `json:"sticker,omitempty"`
	// Message is a forwarded story
	//
	// optional
	Story *Story `json:"story,omitempty"`
	// Video message is a video, information about the video;
	//
	// optional
	Video *Video `json:"video,omitempty"`
	// VideoNote message is a video note, information about the video message;
	//
	// optional
	VideoNote *VideoNote `json:"video_note,omitempty"`
	// Voice message is a voice message, information about the file;
	//
	// optional
	Voice *Voice `json:"voice,omitempty"`
	// Caption for the animation, audio, document, photo, video or voice, 0-1024 characters;
	//
	// optional
	Caption string `json:"caption,omitempty"`
	// For messages with a caption, special entities like usernames, URLs, bot commands, etc. that appear in the caption
	//
	// optional
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
	// True, if the caption must be shown above the message media
	//
	// optional
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// optional
	//Optional. True, if the message media is covered by a spoiler animation
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Optional. Message is a checklist
	// optional
	Checklist *Checklist `json:"checklist,omitempty"`

	// Contact message is a shared contact, information about the contact;
	//
	// optional
	Contact *Contact `json:"contact,omitempty"`
	// Dice is a dice with random value;
	//
	// optional
	Dice *Dice `json:"dice,omitempty"`
	// Game message is a game, information about the game;
	//
	// optional
	Game *Game `json:"game,omitempty"`
	// Poll is a native poll, information about the poll;
	//
	// optional
	Poll *Poll `json:"poll,omitempty"`
	// Venue message is a venue, information about the venue.
	// For backward compatibility, when this field is set, the location field
	// will also be set;
	//
	// optional
	Venue *Venue `json:"venue,omitempty"`
	// Location message is a shared location, information about the location;
	//
	// optional
	Location *Location `json:"location,omitempty"`
	// NewChatMembers that were added to the group or supergroup
	// and information about them (the bot itself may be one of these members);
	//
	// optional
	NewChatMembers []User `json:"new_chat_members,omitempty"`
	// LeftChatMember is a member was removed from the group,
	// information about them (this member may be the bot itself);
	//
	// optional
	LeftChatMember *User `json:"left_chat_member,omitempty"`
	// NewChatTitle is a chat title was changed to this value;
	//
	// optional
	NewChatTitle string `json:"new_chat_title,omitempty"`
	// NewChatPhoto is a chat photo was change to this value;
	//
	// optional
	NewChatPhoto []PhotoSize `json:"new_chat_photo,omitempty"`
	// DeleteChatPhoto is a service message: the chat photo was deleted;
	//
	// optional
	DeleteChatPhoto bool `json:"delete_chat_photo,omitempty"`
	// GroupChatCreated is a service message: the group has been created;
	//
	// optional
	GroupChatCreated bool `json:"group_chat_created,omitempty"`
	// SuperGroupChatCreated is a service message: the supergroup has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a supergroup when it is created.
	// It can only be found in ReplyToMessage if someone replies to a very first message
	// in a directly created supergroup;
	//
	// optional
	SuperGroupChatCreated bool `json:"supergroup_chat_created,omitempty"`
	// ChannelChatCreated is a service message: the channel has been created.
	// This field can't be received in a message coming through updates,
	// because bot can't be a member of a channel when it is created.
	// It can only be found in ReplyToMessage
	// if someone replies to a very first message in a channel;
	//
	// optional
	ChannelChatCreated bool `json:"channel_chat_created,omitempty"`
	// MessageAutoDeleteTimerChanged is a service message: auto-delete timer
	// settings changed in the chat.
	//
	// optional
	MessageAutoDeleteTimerChanged *MessageAutoDeleteTimerChanged `json:"message_auto_delete_timer_changed,omitempty"`
	// MigrateToChatID is the group has been migrated to a supergroup with the specified identifier.
	// This number may be greater than 32 bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64-bit integer
	// or double-precision float type are safe for storing this identifier;
	//
	// optional
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	// MigrateFromChatID is the supergroup has been migrated from a group with the specified identifier.
	// This number may be greater than 32 bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it is smaller than 52 bits, so a signed 64-bit integer
	// or double-precision float type are safe for storing this identifier;
	//
	// optional
	MigrateFromChatID int64 `json:"migrate_from_chat_id,omitempty"`
	// PinnedMessage is a specified message was pinned.
	// Note that the Message object in this field will not contain further ReplyToMessage
	// fields even if it is itself a reply;
	//
	// optional
	PinnedMessage *Message `json:"pinned_message,omitempty"`
	// Invoice message is an invoice for a payment;
	//
	// optional
	Invoice *Invoice `json:"invoice,omitempty"`
	// SuccessfulPayment message is a service message about a successful payment,
	// information about the payment;
	//
	// optional
	SuccessfulPayment *SuccessfulPayment `json:"successful_payment,omitempty"`
	// Message is a service message about a refunded payment, information about the payment.
	//
	// optional
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`
	// Service message: users were shared with the bot
	//
	// optional
	UsersShared *UsersShared `json:"users_shared,omitempty"`
	// Service message: a chat was shared with the bot
	//
	// optional
	ChatShared *ChatShared `json:"chat_shared,omitempty"`
	// The domain name of the website on which the user has logged in.
	//
	// optional
	ConnectedWebsite string `json:"connected_website,omitempty"`
	// Service message: the user allowed the bot to write messages after adding it to the attachment or side menu,
	// launching a Web App from a link, or accepting an explicit request from a Web App sent by the method
	//
	// optional
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed"`
	// PassportData is a Telegram Passport data;
	//
	// optional
	PassportData *PassportData `json:"passport_data,omitempty"`
	// ProximityAlertTriggered is a service message. A user in the chat
	// triggered another user's proximity alert while sharing Live Location
	//
	// optional
	ProximityAlertTriggered *ProximityAlertTriggered `json:"proximity_alert_triggered,omitempty"`
	// Service message: user boosted the chat
	//
	// optional
	BoostAdded *ChatBoostAdded `json:"boost_added,omitempty"`
	// Service message: chat background set
	//
	// optional
	ChatBackgroundSet *ChatBackground `json:"chat_background_set,omitempty"`

	// Optional. Service message: some tasks in a checklist were marked as done or not done
	ChecklistTasksDone *ChecklistTasksDone `json:"checklist_tasks_done,omitempty"`

	// Optional. Service message: tasks were added to a checklist
	ChecklistTasksAdded *ChecklistTasksAdded `json:"checklist_tasks_added,omitempty"`

	//Optional. Service message: the price for paid messages in the corresponding direct messages chat of a channel has changed
	DirectMessagePriceChanged *DirectMessagePriceChanged `json:"direct_message_price_changed,omitempty"`

	// Service message: forum topic created
	//
	// optional
	ForumTopicCreated *ForumTopicCreated `json:"forum_topic_created,omitempty"`
	// Service message: forum post edited
	//
	// optional
	ForumTopicEdited *ForumTopicEdited `json:"forum_topic_edited,omitempty"`
	// Service message: forum post closed
	//
	// optional
	ForumTopicClosed *ForumTopicClosed `json:"forum_topic_closed,omitempty"`
	// Service message: forum post reopened
	//
	// optional
	ForumTopicReopened *ForumTopicReopened `json:"forum_topic_reopened,omitempty"`
	// Service message: the 'General' forum topic hidden
	//
	// optional
	GeneralForumTopicHidden *GeneralForumTopicHidden `json:"general_forum_topic_hidden,omitempty"`
	// Service message: the 'General' forum topic unhidden
	//
	// optional
	GeneralForumTopicUnhidden *GeneralForumTopicUnhidden `json:"general_forum_topic_unhidden,omitempty"`
	// Service message: a scheduled giveaway was created
	//
	// optional
	GiveawayCreated *GiveawayCreated `json:"giveaway_created,omitempty"`
	// The message is a scheduled giveaway message
	//
	// optional
	Giveaway *Giveaway `json:"giveaway,omitempty"`
	// A giveaway with public winners was completed
	//
	// optional
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners,omitempty"`
	// Service message: a giveaway without public winners was completed
	//
	// optional
	GiveawayCompleted *GiveawayCompleted `json:"giveaway_completed,omitempty"`
	// VideoChatScheduled is a service message: video chat scheduled.
	//
	// optional
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`
	// VideoChatStarted is a service message: video chat started.
	//
	// optional
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`
	// VideoChatEnded is a service message: video chat ended.
	//
	// optional
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`
	// VideoChatParticipantsInvited is a service message: new participants
	// invited to a video chat.
	//
	// optional
	VideoChatParticipantsInvited *VideoChatParticipantsInvited `json:"video_chat_participants_invited,omitempty"`
	// WebAppData is a service message: data sent by a Web App.
	//
	// optional
	WebAppData *WebAppData `json:"web_app_data,omitempty"`
	// ReplyMarkup is the Inline keyboard attached to the message.
	// login_url buttons are represented as ordinary url buttons.
	//
	// optional
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup,omitempty"`

	// PaidMessagePriceChanged is a service message: the price for paid messages has changed in the chat.
	//
	// optional
	PaidMessagePriceChanged *PaidMessagePriceChanged `json:"paid_message_price_changed,omitempty"`

	//sender_business_bot
	SenderBusinessBot *User `json:"sender_business_bot,omitempty"`

	//paid_star_count	Integer	Optional. The number of Telegram Stars that were paid by the sender of the message to send it
	PaidStarCount int `json:"paid_star_count,omitempty"`

	//gift	GiftInfo	Optional. Service message: a regular gift was sent or received

	Gift *GiftInfo `json:"gift,omitempty"`

	//unique_gift	UniqueGiftInfo	Optional. Service message: a unique gift was sent or received
	UniqueGift *UniqueGiftInfo `json:"unique_gift,omitempty"`

	// MaybeInaccessibleMessage
	// Message InaccessibleMessage
	Type string `json:"type,omitempty"`
}

// Time converts the message timestamp into a Time.
func (m *Message) Time() time.Time {
	return time.Unix(int64(m.Date), 0)
}

// IsCommand returns true if message starts with a "bot_command" entity.
func (m *Message) IsCommand() bool {
	if m.Entities == nil || len(m.Entities) == 0 {
		return false
	}

	entity := m.Entities[0]
	return entity.Offset == 0 && entity.IsCommand()
}

// Command checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is removed. Use
// CommandWithAt() if you do not want that.
func (m *Message) Command() string {
	command := m.CommandWithAt()

	if i := strings.Index(command, "@"); i != -1 {
		command = command[:i]
	}

	return command
}

// CommandWithAt checks if the message was a command and if it was, returns the
// command. If the Message was not a command, it returns an empty string.
//
// If the command contains the at name syntax, it is not removed. Use Command()
// if you want that.
func (m *Message) CommandWithAt() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := m.Entities[0]
	return m.Text[1:entity.Length]
}

// CommandArguments checks if the message was a command and if it was,
// returns all text after the command name. If the Message was not a
// command, it returns an empty string.
//
// Note: The first character after the command name is omitted:
// - "/foo bar baz" yields "bar baz", not " bar baz"
// - "/foo-bar baz" yields "bar baz", too
// Even though the latter is not a command conforming to the spec, the API
// marks "/foo" as command entity.
func (m *Message) CommandArguments() string {
	if !m.IsCommand() {
		return ""
	}

	// IsCommand() checks that the message begins with a bot_command entity
	entity := m.Entities[0]

	if len(m.Text) == entity.Length {
		return "" // The command makes up the whole message
	}

	return m.Text[entity.Length+1:]
}

// MessageID represents a unique message identifier.
type MessageID struct {
	MessageID int `json:"message_id"`
}

// MessageEntity represents one special entity in a text message.
type MessageEntity struct {
	// Type of the entity.
	// Can be:
	//  “mention” (@username),
	//  “hashtag” (#hashtag),
	//  “cashtag” ($USD),
	//  “bot_command” (/start@jobs_bot),
	//  “url” (https://telegram.org),
	//  “email” (do-not-reply@telegram.org),
	//  “phone_number” (+1-212-555-0123),
	//  “bold” (bold text),
	//  “italic” (italic text),
	//  “underline” (underlined text),
	//  “strikethrough” (strikethrough text),
	//  "spoiler" (spoiler message),
	//  “code” (monowidth string),
	//  “pre” (monowidth block),
	//  “text_link” (for clickable text URLs),
	//  “text_mention” (for users without usernames)
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity
	Offset int `json:"offset"`
	// Length
	Length int `json:"length"`
	// URL for “text_link” only, url that will be opened after user taps on the text
	//
	// optional
	URL string `json:"url,omitempty"`
	// User for “text_mention” only, the mentioned user
	//
	// optional
	User *User `json:"user,omitempty"`
	// Language for “pre” only, the programming language of the entity text
	//
	// optional
	Language string `json:"language,omitempty"`

	// For “custom_emoji” only, unique identifier of the custom emoji. Use getCustomEmojiStickers to get full information about the sticker
	// optional
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// ParseURL attempts to parse a URL contained within a MessageEntity.
func (e MessageEntity) ParseURL() (*url.URL, error) {
	if e.URL == "" {
		return nil, errors.New(ErrBadURL)
	}

	return url.Parse(e.URL)
}

// IsMention returns true if the type of the message entity is "mention" (@username).
func (e MessageEntity) IsMention() bool {
	return e.Type == "mention"
}

// IsTextMention returns true if the type of the message entity is "text_mention"
// (At this time, the user field exists, and occurs when tagging a member without a username)
func (e MessageEntity) IsTextMention() bool {
	return e.Type == "text_mention"
}

// IsHashtag returns true if the type of the message entity is "hashtag".
func (e MessageEntity) IsHashtag() bool {
	return e.Type == "hashtag"
}

// IsCommand returns true if the type of the message entity is "bot_command".
func (e MessageEntity) IsCommand() bool {
	return e.Type == "bot_command"
}

// IsURL returns true if the type of the message entity is "url".
func (e MessageEntity) IsURL() bool {
	return e.Type == "url"
}

// IsEmail returns true if the type of the message entity is "email".
func (e MessageEntity) IsEmail() bool {
	return e.Type == "email"
}

// IsBold returns true if the type of the message entity is "bold" (bold text).
func (e MessageEntity) IsBold() bool {
	return e.Type == "bold"
}

// IsItalic returns true if the type of the message entity is "italic" (italic text).
func (e MessageEntity) IsItalic() bool {
	return e.Type == "italic"
}

// IsCode returns true if the type of the message entity is "code" (monowidth string).
func (e MessageEntity) IsCode() bool {
	return e.Type == "code"
}

// IsPre returns true if the type of the message entity is "pre" (monowidth block).
func (e MessageEntity) IsPre() bool {
	return e.Type == "pre"
}

// IsTextLink returns true if the type of the message entity is "text_link" (clickable text URL).
func (e MessageEntity) IsTextLink() bool {
	return e.Type == "text_link"
}

// PhotoSize represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// FileID identifier for this file, which can be used to download or reuse
	// the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width photo width
	Width int `json:"width"`
	// Height photo height
	Height int `json:"height"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
}

// Animation represents an animation file.
type Animation struct {
	// FileID is the identifier for this file, which can be used to download or reuse
	// the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width video width as defined by sender
	Width int `json:"width"`
	// Height video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumbnail animation thumbnail as defined by sender
	//
	// optional
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// FileName original animation filename as defined by sender
	//
	// optional
	FileName string `json:"file_name,omitempty"`
	// MimeType of the file as defined by sender
	//
	// optional
	MimeType string `json:"mime_type,omitempty"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
}

// Audio represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID is an identifier for this file, which can be used to download or
	// reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// Performer of the audio as defined by sender or by audio tags
	//
	// optional
	Performer string `json:"performer,omitempty"`
	// Title of the audio as defined by sender or by audio tags
	//
	// optional
	Title string `json:"title,omitempty"`
	// FileName is the original filename as defined by sender
	//
	// optional
	FileName string `json:"file_name,omitempty"`
	// MimeType of the file as defined by sender
	//
	// optional
	MimeType string `json:"mime_type,omitempty"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
	// Thumbnail is the album cover to which the music file belongs
	//
	// optional
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
}

// Document represents a general file.
type Document struct {
	// FileID is an identifier for this file, which can be used to download or
	// reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Thumbnail document thumbnail as defined by sender
	//
	// optional
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// FileName original filename as defined by sender
	//
	// optional
	FileName string `json:"file_name,omitempty"`
	// MimeType  of the file as defined by sender
	//
	// optional
	MimeType string `json:"mime_type,omitempty"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
}

// Video represents a video file.
type Video struct {
	// FileID identifier for this file, which can be used to download or reuse
	// the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width video width as defined by sender
	Width int `json:"width"`
	// Height video height as defined by sender
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumbnail video thumbnail
	//
	// optional
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`

	//Optional. Available sizes of the cover of the video in the message
	Cover []*PhotoSize `json:"cover,omitempty"`

	// FileName is the original filename as defined by sender
	// optional
	FileName string `json:"file_name,omitempty"`
	// MimeType of a file as defined by sender
	//
	// optional
	MimeType string `json:"mime_type,omitempty"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`

	//StartTimestamp timestamp in seconds from which the video will play in the message
	//
	// optional
	StartTimestamp int `json:"start_timestamp,omitempty"`
}

// VideoNote object represents a video message.
type VideoNote struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Length video width and height (diameter of the video message) as defined by sender
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender
	Duration int `json:"duration"`
	// Thumbnail video thumbnail
	//
	// optional
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
}

// Voice represents a voice note.
type Voice struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender
	Duration int `json:"duration"`
	// MimeType of the file as defined by sender
	//
	// optional
	MimeType string `json:"mime_type,omitempty"`
	// FileSize file size
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
}

// Contact represents a phone contact.
//
// Note that LastName and UserID may be empty.
type Contact struct {
	// PhoneNumber contact's phone number
	PhoneNumber string `json:"phone_number"`
	// FirstName contact's first name
	FirstName string `json:"first_name"`
	// LastName contact's last name
	//
	// optional
	LastName string `json:"last_name,omitempty"`
	// UserID contact's user identifier in Telegram
	//
	// optional
	UserID int64 `json:"user_id,omitempty"`
	// VCard is additional data about the contact in the form of a vCard.
	//
	// optional
	VCard string `json:"vcard,omitempty"`
}

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	// Emoji on which the dice throw animation is based
	Emoji string `json:"emoji"`
	// Value of the dice
	Value int `json:"value"`
}

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	// Text is the option text, 1-100 characters
	Text string `json:"text"`
	// TextEntities is the Optional. Special entities that appear in the option text. Currently, only custom emoji entities are allowed in poll option texts
	TextEntities []MessageEntity `json:"text_entities"`
	// VoterCount is the number of users that voted for this option
	VoterCount int `json:"voter_count"`
}

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	// PollID is the unique poll identifier
	PollID string `json:"poll_id"`
	//voter_chat	Chat	Optional. The chat that changed the answer to the poll, if the voter is anonymous
	VoterChat *Chat `json:"voter_chat,omitempty"`

	// User who changed the answer to the poll
	User *User `json:"user,omitempty"`
	// OptionIDs is the 0-based identifiers of poll options chosen by the user.
	// May be empty if user retracted vote.
	OptionIDs []int `json:"option_ids"`
}

// Poll contains information about a poll.
type Poll struct {
	// ID is the unique poll identifier
	ID string `json:"id"`
	// Question is the poll question, 1-255 characters
	Question string `json:"question"`
	// QuestionEntities Optional, Special entities that appear in the question. Currently, only custom emoji entities are allowed in poll questions
	QuestionEntities string `json:"question_entities"`
	// Options is the list of poll options
	Options []PollOption `json:"options"`
	// TotalVoterCount is the total numbers of users who voted in the poll
	TotalVoterCount int `json:"total_voter_count"`
	// IsClosed is if the poll is closed
	IsClosed bool `json:"is_closed"`
	// IsAnonymous is if the poll is anonymous
	IsAnonymous bool `json:"is_anonymous"`
	// Type is the poll type, currently can be "regular" or "quiz"
	Type string `json:"type"`
	// AllowsMultipleAnswers is true, if the poll allows multiple answers
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// CorrectOptionID is the 0-based identifier of the correct answer option.
	// Available only for polls in quiz mode, which are closed, or was sent (not
	// forwarded) by the bot or to the private chat with the bot.
	//
	// optional
	CorrectOptionID int `json:"correct_option_id,omitempty"`
	// Explanation is text that is shown when a user chooses an incorrect answer
	// or taps on the lamp icon in a quiz-style poll, 0-200 characters
	//
	// optional
	Explanation string `json:"explanation,omitempty"`
	// ExplanationEntities are special entities like usernames, URLs, bot
	// commands, etc. that appear in the explanation
	//
	// optional
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	// OpenPeriod is the amount of time in seconds the poll will be active
	// after creation
	//
	// optional
	OpenPeriod int `json:"open_period,omitempty"`
	// CloseDate is the point in time (unix timestamp) when the poll will be
	// automatically closed
	//
	// optional
	CloseDate int `json:"close_date,omitempty"`
}

// Location represents a point on the map.
type Location struct {
	// Longitude as defined by sender
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
	// HorizontalAccuracy is the radius of uncertainty for the location,
	// measured in meters; 0-1500
	//
	// optional
	HorizontalAccuracy float64 `json:"horizontal_accuracy,omitempty"`
	// LivePeriod is time relative to the message sending date, during which the
	// location can be updated, in seconds. For active live locations only.
	//
	// optional
	LivePeriod int `json:"live_period,omitempty"`
	// Heading is the direction in which user is moving, in degrees; 1-360. For
	// active live locations only.
	//
	// optional
	Heading int `json:"heading,omitempty"`
	// ProximityAlertRadius is the maximum distance for proximity alerts about
	// approaching another chat member, in meters. For sent live locations only.
	//
	// optional
	ProximityAlertRadius int `json:"proximity_alert_radius,omitempty"`
}

// Venue represents a venue.
type Venue struct {
	// Location is the venue location
	Location Location `json:"location"`
	// Title is the name of the venue
	Title string `json:"title"`
	// Address of the venue
	Address string `json:"address"`
	// FoursquareID is the foursquare identifier of the venue
	//
	// optional
	FoursquareID string `json:"foursquare_id,omitempty"`
	// FoursquareType is the foursquare type of the venue
	//
	// optional
	FoursquareType string `json:"foursquare_type,omitempty"`
	// GooglePlaceID is the Google Places identifier of the venue
	//
	// optional
	GooglePlaceID string `json:"google_place_id,omitempty"`
	// GooglePlaceType is the Google Places type of the venue
	//
	// optional
	GooglePlaceType string `json:"google_place_type,omitempty"`
}

// WebAppData Contains data sent from a Web App to the bot.
type WebAppData struct {
	// Data is the data. Be aware that a bad client can send arbitrary data in this field.
	Data string `json:"data"`
	// ButtonText is the text of the web_app keyboard button, from which the Web App
	// was opened. Be aware that a bad client can send arbitrary data in this field.
	ButtonText string `json:"button_text"`
}

// ProximityAlertTriggered represents a service message sent when a user in the
// chat triggers a proximity alert sent by another user.
type ProximityAlertTriggered struct {
	// Traveler is the user that triggered the alert
	Traveler User `json:"traveler"`
	// Watcher is the user that set the alert
	Watcher User `json:"watcher"`
	// Distance is the distance between the users
	Distance int `json:"distance"`
}

// MessageAutoDeleteTimerChanged represents a service message about a change in
// auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}

// VideoChatScheduled represents a service message about a voice chat scheduled
// in the chat.
type VideoChatScheduled struct {
	// Point in time (Unix timestamp) when the voice chat is supposed to be
	// started by a chat administrator
	StartDate int `json:"start_date"`
}

// Time converts the scheduled start date into a Time.
func (m *VideoChatScheduled) Time() time.Time {
	return time.Unix(int64(m.StartDate), 0)
}

// VideoChatStarted represents a service message about a voice chat started in
// the chat.
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a voice chat ended in the
// chat.
type VideoChatEnded struct {
	// Voice chat duration; in seconds.
	Duration int `json:"duration"`
}

// VideoChatParticipantsInvited represents a service message about new members
// invited to a voice chat.
type VideoChatParticipantsInvited struct {
	// New members that were invited to the voice chat.
	//
	// optional
	Users []User `json:"users,omitempty"`
}

// UserProfilePhotos contains a set of user profile photos.
type UserProfilePhotos struct {
	// TotalCount total number of profile pictures the target user has
	TotalCount int `json:"total_count"`
	// Photos requested profile pictures (in up to 4 sizes each)
	Photos [][]PhotoSize `json:"photos"`
}

// File contains information about a file to download from Telegram.
type File struct {
	// FileID identifier for this file, which can be used to download or reuse
	// the file
	FileID string `json:"file_id"`
	// FileUniqueID is the unique identifier for this file, which is supposed to
	// be the same over time and for different bots. Can't be used to download
	// or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// FileSize file size, if known
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
	// FilePath file path
	//
	// optional
	FilePath string `json:"file_path,omitempty"`
}

// Link returns a full path to the download URL for a File.
//
// It requires the Bot token to create the link.
func (f *File) Link(token string) string {
	return fmt.Sprintf(FileEndpoint, token, f.FilePath)
}

// WebAppInfo contains information about a Web App.
type WebAppInfo struct {
	// URL is the HTTPS URL of a Web App to be opened with additional data as
	// specified in Initializing Web Apps.
	URL string `json:"url"`
}

// ReplyKeyboardMarkup represents a custom keyboard with reply options.
type ReplyKeyboardMarkup struct {
	// Keyboard is an array of button rows, each represented by an Array of KeyboardButton objects
	Keyboard [][]KeyboardButton `json:"keyboard"`
	// ResizeKeyboard requests clients to resize the keyboard vertically for optimal fit
	// (e.g., make the keyboard smaller if there are just two rows of buttons).
	// Defaults to false, in which case the custom keyboard
	// is always of the same height as the app's standard keyboard.
	//
	// optional
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`
	// OneTimeKeyboard requests clients to hide the keyboard as soon as it's been used.
	// The keyboard will still be available, but clients will automatically display
	// the usual letter-keyboard in the chat – the user can press a special button
	// in the input field to see the custom keyboard again.
	// Defaults to false.
	//
	// optional
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`
	// InputFieldPlaceholder is the placeholder to be shown in the input field when
	// the keyboard is active; 1-64 characters.
	//
	// optional
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// Selective use this parameter if you want to show the keyboard to specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has Message.ReplyToMessage not nil), sender of the original message.
	//
	// Example: A user requests to change the bot's language,
	// bot replies to the request with a keyboard to select the new language.
	// Other users in the group don't see the keyboard.
	//
	// optional
	Selective bool `json:"selective,omitempty"`
}

// KeyboardButton represents one button of the reply keyboard. For simple text
// buttons String can be used instead of this object to specify text of the
// button. Optional fields request_contact, request_location, and request_poll
// are mutually exclusive.
type KeyboardButton struct {
	// Text of the button. If none of the optional fields are used,
	// it will be sent as a message when the button is pressed.
	Text string `json:"text"`
	// RequestContact if True, the user's phone number will be sent
	// as a contact when the button is pressed.
	// Available in private chats only.
	//
	// optional
	RequestContact bool `json:"request_contact,omitempty"`
	// RequestLocation if True, the user's current location will be sent when
	// the button is pressed.
	// Available in private chats only.
	//
	// optional
	RequestLocation bool `json:"request_location,omitempty"`
	// RequestPoll if specified, the user will be asked to create a poll and send it
	// to the bot when the button is pressed. Available in private chats only
	//
	// optional
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
	// WebApp if specified, the described Web App will be launched when the button
	// is pressed. The Web App will be able to send a “web_app_data” service
	// message. Available in private chats only.
	//
	// optional
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// KeyboardButtonPollType represents type of poll, which is allowed to
// be created and sent when the corresponding button is pressed.
type KeyboardButtonPollType struct {
	// Type is if quiz is passed, the user will be allowed to create only polls
	// in the quiz mode. If regular is passed, only regular polls will be
	// allowed. Otherwise, the user will be allowed to create a poll of any type.
	Type string `json:"type"`
}

// ReplyKeyboardRemove Upon receiving a message with this object, Telegram
// clients will remove the current custom keyboard and display the default
// letter-keyboard. By default, custom keyboards are displayed until a new
// keyboard is sent by a bot. An exception is made for one-time keyboards
// that are hidden immediately after the user presses a button.
type ReplyKeyboardRemove struct {
	// RemoveKeyboard requests clients to remove the custom keyboard
	// (user will not be able to summon this keyboard;
	// if you want to hide the keyboard from sight but keep it accessible,
	// use one_time_keyboard in ReplyKeyboardMarkup).
	RemoveKeyboard bool `json:"remove_keyboard"`
	// Selective use this parameter if you want to remove the keyboard for specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has Message.ReplyToMessage not nil), sender of the original message.
	//
	// Example: A user votes in a poll, bot returns confirmation message
	// in reply to the vote and removes the keyboard for that user,
	// while still showing the keyboard with poll options to users who haven't voted yet.
	//
	// optional
	Selective bool `json:"selective,omitempty"`
}

// InlineKeyboardMarkup represents an inline keyboard that appears right next to
// the message it belongs to.
type InlineKeyboardMarkup struct {
	// InlineKeyboard array of button rows, each represented by an Array of
	// InlineKeyboardButton objects
	InlineKeyboard [][]InlineKeyboardButton `json:"inline_keyboard"`
}

// InlineKeyboardButton represents one button of an inline keyboard. You must
// use exactly one of the optional fields.
//
// Note that some values are references as even an empty string
// will change behavior.
//
// CallbackGame, if set, MUST be first button in first row.
type InlineKeyboardButton struct {
	// Text label text on the button
	Text string `json:"text"`
	// URL HTTP or tg:// url to be opened when button is pressed.
	//
	// optional
	URL *string `json:"url,omitempty"`
	// LoginURL is an HTTP URL used to automatically authorize the user. Can be
	// used as a replacement for the Telegram Login Widget
	//
	// optional
	LoginURL *LoginURL `json:"login_url,omitempty"`
	// CallbackData data to be sent in a callback query to the bot when button is pressed, 1-64 bytes.
	//
	// optional
	CallbackData *string `json:"callback_data,omitempty"`
	// WebApp is the Description of the Web App that will be launched when the user presses the button.
	// The Web App will be able to send an arbitrary message on behalf of the user using the method
	// answerWebAppQuery. Available only in private chats between a user and the bot.
	//
	// optional
	WebApp *WebAppInfo `json:"web_app,omitempty"`
	// SwitchInlineQuery if set, pressing the button will prompt the user to select one of their chats,
	// open that chat and insert the bot's username and the specified inline query in the input field.
	// Can be empty, in which case just the bot's username will be inserted.
	//
	// This offers an easy way for users to start using your bot
	// in inline mode when they are currently in a private chat with it.
	// Especially useful when combined with switch_pm… actions – in this case
	// the user will be automatically returned to the chat they switched from,
	// skipping the chat selection screen.
	//
	// optional
	SwitchInlineQuery *string `json:"switch_inline_query,omitempty"`
	// SwitchInlineQueryCurrentChat if set, pressing the button will insert the bot's username
	// and the specified inline query in the current chat's input field.
	// Can be empty, in which case only the bot's username will be inserted.
	//
	// This offers a quick way for the user to open your bot in inline mode
	// in the same chat – good for selecting something from multiple options.
	//
	// optional
	SwitchInlineQueryCurrentChat *string `json:"switch_inline_query_current_chat,omitempty"`
	// CallbackGame description of the game that will be launched when the user presses the button.
	//
	// optional
	CallbackGame *CallbackGame `json:"callback_game,omitempty"`
	// Pay specify True, to send a Pay button.
	//
	// NOTE: This type of button must always be the first button in the first row.
	//
	// optional
	Pay bool `json:"pay,omitempty"`

	//Optional. If set, pressing the button will prompt the user to select one of their chats of the specified type, open that chat and
	// insert the bot's username and the specified inline query in the input field. Not supported for messages sent on behalf of a Telegram Business account.
	// optional
	SwitchInlineQueryChosenChat *SwitchInlineQueryChosenChat `json:"switch_inline_query_chosen_chat,omitempty"`

	//copy_text
	// optional
	CopyText *CopyTextButton `json:"copy_text,omitempty"`
}

// This object represents an inline button that switches the current user to inline mode in a chosen chat, with an optional default inline query.
type SwitchInlineQueryChosenChat struct {

	/**
	query	String	Optional. The default inline query to be inserted in the input field. If left empty, only the bot's username will be inserted
	allow_user_chats	Boolean	Optional. True, if private chats with users can be chosen
	allow_bot_chats	Boolean	Optional. True, if private chats with bots can be chosen
	allow_group_chats	Boolean	Optional. True, if group and supergroup chats can be chosen
	allow_channel_chats	Boolean	Optional. True, if channel chats can be chosen
	**/

	Query             string `json:"query,omitempty"`
	AllowUserChats    bool   `json:"allow_user_chats,omitempty"`
	AllowBotChats     bool   `json:"allow_bot_chats,omitempty"`
	AllowGroupChats   bool   `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool   `json:"allow_channel_chats,omitempty"`
}

type CopyTextButton struct {
	// The text to be copied to the clipboard; 1-256 characters
	Text string `json:"text"`
}

// LoginURL represents a parameter of the inline keyboard button used to
// automatically authorize a user. Serves as a great replacement for the
// Telegram Login Widget when the user is coming from Telegram. All the user
// needs to do is tap/click a button and confirm that they want to log in.
type LoginURL struct {
	// URL is an HTTP URL to be opened with user authorization data added to the
	// query string when the button is pressed. If the user refuses to provide
	// authorization data, the original URL without information about the user
	// will be opened. The data added is the same as described in Receiving
	// authorization data.
	//
	// NOTE: You must always check the hash of the received data to verify the
	// authentication and the integrity of the data as described in Checking
	// authorization.
	URL string `json:"url"`
	// ForwardText is the new text of the button in forwarded messages
	//
	// optional
	ForwardText string `json:"forward_text,omitempty"`
	// BotUsername is the username of a bot, which will be used for user
	// authorization. See Setting up a bot for more details. If not specified,
	// the current bot's username will be assumed. The url's domain must be the
	// same as the domain linked with the bot. See Linking your domain to the
	// bot for more details.
	//
	// optional
	BotUsername string `json:"bot_username,omitempty"`
	// RequestWriteAccess if true requests permission for your bot to send
	// messages to the user
	//
	// optional
	RequestWriteAccess bool `json:"request_write_access,omitempty"`
}

// CallbackQuery represents an incoming callback query from a callback button in
// an inline keyboard. If the button that originated the query was attached to a
// message sent by the bot, the field message will be present. If the button was
// attached to a message sent via the bot (in inline mode), the field
// inline_message_id will be present. Exactly one of the fields data or
// game_short_name will be present.
type CallbackQuery struct {
	// ID unique identifier for this query
	ID string `json:"id"`
	// From sender
	From *User `json:"from"`
	// Message with the callback button that originated the query.
	// Note that message content and message date will not be available if the
	// message is too old.
	//
	// optional
	Message *Message `json:"message,omitempty"`
	// InlineMessageID identifier of the message sent via the bot in inline
	// mode, that originated the query.
	//
	// optional
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// ChatInstance global identifier, uniquely corresponding to the chat to
	// which the message with the callback button was sent. Useful for high
	// scores in games.
	ChatInstance string `json:"chat_instance"`
	// Data associated with the callback button. Be aware that
	// a bad client can send arbitrary data in this field.
	//
	// optional
	Data string `json:"data,omitempty"`
	// GameShortName short name of a Game to be returned, serves as the unique identifier for the game.
	//
	// optional
	GameShortName string `json:"game_short_name,omitempty"`
}

// ForceReply when receiving a message with this object, Telegram clients will
// display a reply interface to the user (act as if the user has selected the
// bot's message and tapped 'Reply'). This can be extremely useful if you  want
// to create user-friendly step-by-step interfaces without having to sacrifice
// privacy mode.
type ForceReply struct {
	// ForceReply shows reply interface to the user,
	// as if they manually selected the bot's message and tapped 'Reply'.
	ForceReply bool `json:"force_reply"`
	// InputFieldPlaceholder is the placeholder to be shown in the input field when
	// the reply is active; 1-64 characters.
	//
	// optional
	InputFieldPlaceholder string `json:"input_field_placeholder,omitempty"`
	// Selective use this parameter if you want to force reply from specific users only.
	// Targets:
	//  1) users that are @mentioned in the text of the Message object;
	//  2) if the bot's message is a reply (has Message.ReplyToMessage not nil), sender of the original message.
	//
	// optional
	Selective bool `json:"selective,omitempty"`
}

// ChatPhoto represents a chat photo.
type ChatPhoto struct {
	// SmallFileID is a file identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download and
	// only for as long as the photo is not changed.
	SmallFileID string `json:"small_file_id"`
	// SmallFileUniqueID is a unique file identifier of small (160x160) chat
	// photo, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueID string `json:"small_file_unique_id"`
	// BigFileID is a file identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download and
	// only for as long as the photo is not changed.
	BigFileID string `json:"big_file_id"`
	// BigFileUniqueID is a file identifier of big (640x640) chat photo, which
	// is supposed to be the same over time and for different bots. Can't be
	// used to download or reuse the file.
	BigFileUniqueID string `json:"big_file_unique_id"`
}

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	// InviteLink is the invite link. If the link was created by another chat
	// administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`
	// Creator of the link.
	Creator User `json:"creator"`
	// CreatesJoinRequest is true if users joining the chat via the link need to
	// be approved by chat administrators.
	//
	// optional
	CreatesJoinRequest bool `json:"creates_join_request,omitempty"`
	// IsPrimary is true, if the link is primary.
	IsPrimary bool `json:"is_primary"`
	// IsRevoked is true, if the link is revoked.
	IsRevoked bool `json:"is_revoked"`
	// Name is the name of the invite link.
	//
	// optional
	Name string `json:"name,omitempty"`
	// ExpireDate is the point in time (Unix timestamp) when the link will
	// expire or has been expired.
	//
	// optional
	ExpireDate int `json:"expire_date,omitempty"`
	// MemberLimit is the maximum number of users that can be members of the
	// chat simultaneously after joining the chat via this invite link; 1-99999.
	//
	// optional
	MemberLimit int `json:"member_limit,omitempty"`
	// PendingJoinRequestCount is the number of pending join requests created
	// using this link.
	//
	// optional
	PendingJoinRequestCount int `json:"pending_join_request_count,omitempty"`

	//subscription_period	Integer	Optional. The number of seconds the subscription will be active for before the next payment
	SubscriptionPeriod int `json:"subscription_period,omitempty"`

	//subscription_price	Integer	Optional. The amount of Telegram Stars a user must pay initially and after each subsequent subscription period to be a member of the chat using the link
	SubscriptionPrice int `json:"subscription_price,omitempty"`
}

type ChatAdministratorRights struct {
	IsAnonymous         bool `json:"is_anonymous"`
	CanManageChat       bool `json:"can_manage_chat"`
	CanDeleteMessages   bool `json:"can_delete_messages"`
	CanManageVideoChats bool `json:"can_manage_video_chats"`
	CanRestrictMembers  bool `json:"can_restrict_members"`
	CanPromoteMembers   bool `json:"can_promote_members"`
	CanChangeInfo       bool `json:"can_change_info"`
	CanInviteUsers      bool `json:"can_invite_users"`
	CanPostMessages     bool `json:"can_post_messages"`
	CanEditMessages     bool `json:"can_edit_messages"`
	CanPinMessages      bool `json:"can_pin_messages"`
}

// ChatMember contains information about one member of a chat.
type ChatMember struct {
	// User information about the user
	User *User `json:"user"`
	// Status the member's status in the chat.
	// Can be
	//  “creator”,
	//  “administrator”,
	//  “member”,
	//  “restricted”,
	//  “left” or
	//  “kicked”
	Status string `json:"status"`
	// CustomTitle owner and administrators only. Custom title for this user
	//
	// optional
	CustomTitle string `json:"custom_title,omitempty"`
	// IsAnonymous owner and administrators only. True, if the user's presence
	// in the chat is hidden
	//
	// optional
	IsAnonymous bool `json:"is_anonymous,omitempty"`
	// UntilDate restricted and kicked only.
	// Date when restrictions will be lifted for this user;
	// unix time.
	//
	// optional
	UntilDate int64 `json:"until_date,omitempty"`
	// CanBeEdited administrators only.
	// True, if the bot is allowed to edit administrator privileges of that user.
	//
	// optional
	CanBeEdited bool `json:"can_be_edited,omitempty"`
	// CanManageChat administrators only.
	// True, if the administrator can access the chat event log, chat
	// statistics, message statistics in channels, see channel members, see
	// anonymous administrators in supergroups and ignore slow mode. Implied by
	// any other administrator privilege.
	//
	// optional
	CanManageChat bool `json:"can_manage_chat,omitempty"`
	// CanPostMessages administrators only.
	// True, if the administrator can post in the channel;
	// channels only.
	//
	// optional
	CanPostMessages bool `json:"can_post_messages,omitempty"`
	// CanEditMessages administrators only.
	// True, if the administrator can edit messages of other users and can pin messages;
	// channels only.
	//
	// optional
	CanEditMessages bool `json:"can_edit_messages,omitempty"`
	// CanDeleteMessages administrators only.
	// True, if the administrator can delete messages of other users.
	//
	// optional
	CanDeleteMessages bool `json:"can_delete_messages,omitempty"`
	// CanManageVideoChats administrators only.
	// True, if the administrator can manage video chats.
	//
	// optional
	CanManageVideoChats bool `json:"can_manage_video_chats,omitempty"`
	// CanRestrictMembers administrators only.
	// True, if the administrator can restrict, ban or unban chat members.
	//
	// optional
	CanRestrictMembers bool `json:"can_restrict_members,omitempty"`
	// CanPromoteMembers administrators only.
	// True, if the administrator can add new administrators
	// with a subset of their own privileges or demote administrators that he has promoted,
	// directly or indirectly (promoted by administrators that were appointed by the user).
	//
	// optional
	CanPromoteMembers bool `json:"can_promote_members,omitempty"`
	// CanChangeInfo administrators and restricted only.
	// True, if the user is allowed to change the chat title, photo and other settings.
	//
	// optional
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// CanInviteUsers administrators and restricted only.
	// True, if the user is allowed to invite new users to the chat.
	//
	// optional
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// CanPinMessages administrators and restricted only.
	// True, if the user is allowed to pin messages; groups and supergroups only
	//
	// optional
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
	// IsMember is true, if the user is a member of the chat at the moment of
	// the request
	IsMember bool `json:"is_member"`
	// CanSendMessages
	//
	// optional
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// CanSendMediaMessages restricted only.
	// True, if the user is allowed to send text messages, contacts, locations and venues
	//
	// optional
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// CanSendPolls restricted only.
	// True, if the user is allowed to send polls
	//
	// optional
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// CanSendOtherMessages restricted only.
	// True, if the user is allowed to send audios, documents,
	// photos, videos, video notes and voice notes.
	//
	// optional
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// CanAddWebPagePreviews restricted only.
	// True, if the user is allowed to add web page previews to their messages.
	//
	// optional
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
}

// IsCreator returns if the ChatMember was the creator of the chat.
func (chat ChatMember) IsCreator() bool { return chat.Status == "creator" }

// IsAdministrator returns if the ChatMember is a chat administrator.
func (chat ChatMember) IsAdministrator() bool { return chat.Status == "administrator" }

// HasLeft returns if the ChatMember left the chat.
func (chat ChatMember) HasLeft() bool { return chat.Status == "left" }

// WasKicked returns if the ChatMember was kicked from the chat.
func (chat ChatMember) WasKicked() bool { return chat.Status == "kicked" }

// ChatMemberUpdated represents changes in the status of a chat member.
type ChatMemberUpdated struct {
	// Chat the user belongs to.
	Chat *Chat `json:"chat,omitempty"`
	// From is the performer of the action, which resulted in the change.
	From *User `json:"from,omitempty"`
	// Date the change was done in Unix time.
	Date int `json:"date"`
	// Previous information about the chat member.
	OldChatMember *ChatMember `json:"old_chat_member,omitempty"`
	// New information about the chat member.
	NewChatMember *ChatMember `json:"new_chat_member,omitempty"`
	// InviteLink is the link which was used by the user to join the chat;
	// for joining by invite link events only.
	//
	// optional
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`
	// True, if the user joined the chat after sending a direct join request without using an invite link and being approved by an administrator
	ViaJoinRequest bool `json:"via_join_request,omitempty"`
	// True, if the user joined the chat via a chat folder invite link
	ViaChatFolderInviteLink bool `json:"via_chat_folder_invite_link,omitempty"`
}

// ChatJoinRequest represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat to which the request was sent.
	Chat *Chat `json:"chat,omitempty"`
	// User that sent the join request.
	From *User `json:"from,omitempty"`
	// Date the request was sent in Unix time.
	Date int `json:"date"`
	// Bio of the user.
	//
	// optional
	Bio string `json:"bio,omitempty"`
	// InviteLink is the link that was used by the user to send the join request.
	//
	// optional
	InviteLink *ChatInviteLink `json:"invite_link,omitempty"`

	//user_chat_id	Integer	Identifier of a private chat with the user who sent the join request. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier. The bot can use this identifier for 5 minutes to send messages until the join request is processed, assuming no other administrator contacted the user.
	UserChatID int `json:"user_chat_id,omitempty"`
}

// ChatPermissions describes actions that a non-administrator user is
// allowed to take in a chat. All fields are optional.
type ChatPermissions struct {
	// CanSendMessages is true, if the user is allowed to send text messages,
	// contacts, locations and venues
	//
	// optional
	CanSendMessages bool `json:"can_send_messages,omitempty"`
	// CanSendMediaMessages is true, if the user is allowed to send audios,
	// documents, photos, videos, video notes and voice notes, implies
	// can_send_messages
	//
	// optional
	CanSendMediaMessages bool `json:"can_send_media_messages,omitempty"`
	// CanSendPolls is true, if the user is allowed to send polls, implies
	// can_send_messages
	//
	// optional
	CanSendPolls bool `json:"can_send_polls,omitempty"`
	// CanSendOtherMessages is true, if the user is allowed to send animations,
	// games, stickers and use inline bots, implies can_send_media_messages
	//
	// optional
	CanSendOtherMessages bool `json:"can_send_other_messages,omitempty"`
	// CanAddWebPagePreviews is true, if the user is allowed to add web page
	// previews to their messages, implies can_send_media_messages
	//
	// optional
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews,omitempty"`
	// CanChangeInfo is true, if the user is allowed to change the chat title,
	// photo and other settings. Ignored in public supergroups
	//
	// optional
	CanChangeInfo bool `json:"can_change_info,omitempty"`
	// CanInviteUsers is true, if the user is allowed to invite new users to the
	// chat
	//
	// optional
	CanInviteUsers bool `json:"can_invite_users,omitempty"`
	// CanPinMessages is true, if the user is allowed to pin messages. Ignored
	// in public supergroups
	//
	// optional
	CanPinMessages bool `json:"can_pin_messages,omitempty"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	// Location is the location to which the supergroup is connected. Can't be a
	// live location.
	Location Location `json:"location"`
	// Address is the location address; 1-64 characters, as defined by the chat
	// owner
	Address string `json:"address"`
}

// BotCommand represents a bot command.
type BotCommand struct {
	// Command text of the command, 1-32 characters.
	// Can contain only lowercase English letters, digits and underscores.
	Command string `json:"command"`
	// Description of the command, 3-256 characters.
	Description string `json:"description"`
}

/*
StarAmount
Describes an amount of Telegram Stars.

Field	Type	Description
amount	Integer	Integer amount of Telegram Stars, rounded to 0; can be negative
nanostar_amount	Integer	Optional. The number of 1/1000000000 shares of Telegram Stars; from -999999999 to 999999999; can be negative if and only if amount is non-positive
*/

type StarAmount struct {
	Amount         int `json:"amount"`
	NanostarAmount int `json:"nanostar_amount,omitempty"`
}

// BotCommandScope represents the scope to which bot commands are applied.
//
// It contains the fields for all types of scopes, different types only support
// specific (or no) fields.
type BotCommandScope struct {
	Type   string `json:"type"`
	ChatID int64  `json:"chat_id,omitempty"`
	UserID int64  `json:"user_id,omitempty"`
}

// MenuButton describes the bot's menu button in a private chat.
type MenuButton struct {
	// Type is the type of menu button, must be one of:
	// - `commands`
	// - `web_app`
	// - `default`
	Type string `json:"type"`
	// Text is the text on the button, for `web_app` type.
	Text string `json:"text,omitempty"`
	// WebApp is the description of the Web App that will be launched when the
	// user presses the button for the `web_app` type.
	WebApp *WebAppInfo `json:"web_app,omitempty"`
}

// ResponseParameters are various errors that can be returned in APIResponse.
type ResponseParameters struct {
	// The group has been migrated to a supergroup with the specified identifier.
	//
	// optional
	MigrateToChatID int64 `json:"migrate_to_chat_id,omitempty"`
	// In case of exceeding flood control, the number of seconds left to wait
	// before the request can be repeated.
	//
	// optional
	RetryAfter int `json:"retry_after,omitempty"`
}

// BaseInputMedia is a base type for the InputMedia types.
type BaseInputMedia struct {
	// Type of the result.
	Type string `json:"type"`
	// Media file to send. Pass a file_id to send a file
	// that exists on the Telegram servers (recommended),
	// pass an HTTP URL for Telegram to get a file from the Internet,
	// or pass “attach://<file_attach_name>” to upload a new one
	// using multipart/form-data under <file_attach_name> name.
	Media RequestFileData `json:"media"`
	// thumb intentionally missing as it is not currently compatible

	// Caption of the video to be sent, 0-1024 characters after entities parsing.
	//
	// optional
	Caption string `json:"caption,omitempty"`
	// ParseMode mode for parsing entities in the video caption.
	// See formatting options for more details
	// (https://core.telegram.org/bots/api#formatting-options).
	//
	// optional
	ParseMode string `json:"parse_mode,omitempty"`
	// CaptionEntities is a list of special entities that appear in the caption,
	// which can be specified instead of parse_mode
	//
	// optional
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

// InputMediaPhoto is a photo to send as part of a media group.
type InputMediaPhoto struct {
	BaseInputMedia
}

// InputMediaVideo is a video to send as part of a media group.
type InputMediaVideo struct {
	BaseInputMedia
	// Thumbnail of the file sent; can be ignored if thumbnail generation for
	// the file is supported server-side.
	//
	// optional
	Thumb RequestFileData `json:"thumb,omitempty"`
	// Width video width
	//
	// optional
	Width int `json:"width,omitempty"`
	// Height video height
	//
	// optional
	Height int `json:"height,omitempty"`
	// Duration video duration
	//
	// optional
	Duration int `json:"duration,omitempty"`
	// SupportsStreaming pass True, if the uploaded video is suitable for streaming.
	//
	// optional
	SupportsStreaming bool `json:"supports_streaming,omitempty"`
}

// InputMediaAnimation is an animation to send as part of a media group.
type InputMediaAnimation struct {
	BaseInputMedia
	// Thumbnail of the file sent; can be ignored if thumbnail generation for
	// the file is supported server-side.
	//
	// optional
	Thumb RequestFileData `json:"thumb,omitempty"`
	// Width video width
	//
	// optional
	Width int `json:"width,omitempty"`
	// Height video height
	//
	// optional
	Height int `json:"height,omitempty"`
	// Duration video duration
	//
	// optional
	Duration int `json:"duration,omitempty"`
}

// InputMediaAudio is an audio to send as part of a media group.
type InputMediaAudio struct {
	BaseInputMedia
	// Thumbnail of the file sent; can be ignored if thumbnail generation for
	// the file is supported server-side.
	//
	// optional
	Thumb RequestFileData `json:"thumb,omitempty"`
	// Duration of the audio in seconds
	//
	// optional
	Duration int `json:"duration,omitempty"`
	// Performer of the audio
	//
	// optional
	Performer string `json:"performer,omitempty"`
	// Title of the audio
	//
	// optional
	Title string `json:"title,omitempty"`
}

// InputMediaDocument is a general file to send as part of a media group.
type InputMediaDocument struct {
	BaseInputMedia
	// Thumbnail of the file sent; can be ignored if thumbnail generation for
	// the file is supported server-side.
	//
	// optional
	Thumb RequestFileData `json:"thumb,omitempty"`
	// DisableContentTypeDetection disables automatic server-side content type
	// detection for files uploaded using multipart/form-data. Always true, if
	// the document is sent as part of an album
	//
	// optional
	DisableContentTypeDetection bool `json:"disable_content_type_detection,omitempty"`
}

// Sticker represents a sticker.
type Sticker struct {
	// FileID is an identifier for this file, which can be used to download or
	// reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID is a unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`

	//type	String	Type of the sticker, currently one of “regular”, “mask”, “custom_emoji”. The type of the sticker is independent from its format, which is determined by the fields is_animated and is_video.
	Type string `json:"type"`

	// Width sticker width
	Width int `json:"width"`
	// Height sticker height
	Height int `json:"height"`
	// IsAnimated true, if the sticker is animated
	//
	// optional
	IsAnimated bool `json:"is_animated,omitempty"`
	// IsVideo true, if the sticker is a video sticker
	//
	// optional
	IsVideo bool `json:"is_video,omitempty"`
	// Thumbnail sticker thumbnail in the .WEBP or .JPG format
	//
	// optional
	Thumbnail *PhotoSize `json:"thumbnail,omitempty"`
	// Emoji associated with the sticker
	//
	// optional
	Emoji string `json:"emoji,omitempty"`
	// SetName of the sticker set to which the sticker belongs
	//
	// optional
	SetName string `json:"set_name,omitempty"`
	// PremiumAnimation for premium regular stickers, premium animation for the sticker
	//
	// optional
	PremiumAnimation *File `json:"premium_animation,omitempty"`
	// MaskPosition is for mask stickers, the position where the mask should be
	// placed
	//
	// optional
	MaskPosition *MaskPosition `json:"mask_position,omitempty"`
	// CustomEmojiID for custom emoji stickers, unique identifier of the custom emoji
	//
	// optional
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`

	//needs_repainting	True	Optional. True, if the sticker must be repainted to a text color in messages, the color of the Telegram Premium badge in emoji status, white color on chat photos, or another appropriate color in other places
	NeedsRepainting bool `json:"needs_repainting,omitempty"`

	// FileSize
	//
	// optional
	FileSize int `json:"file_size,omitempty"`
}

// StickerSet represents a sticker set.
type StickerSet struct {
	// Name sticker set name
	Name string `json:"name"`
	// Title sticker set title
	Title string `json:"title"`
	// StickerType of stickers in the set, currently one of “regular”, “mask”, “custom_emoji”
	StickerType string `json:"sticker_type"`
	// IsAnimated true, if the sticker set contains animated stickers
	IsAnimated bool `json:"is_animated"`
	// IsVideo true, if the sticker set contains video stickers
	IsVideo bool `json:"is_video"`
	// ContainsMasks true, if the sticker set contains masks
	ContainsMasks bool `json:"contains_masks"`
	// Stickers list of all set stickers
	Stickers []Sticker `json:"stickers"`
	// Thumb is the sticker set thumbnail in the .WEBP or .TGS format
	Thumbnail *PhotoSize `json:"thumb"`
}

// MaskPosition describes the position on faces where a mask should be placed
// by default.
type MaskPosition struct {
	// The part of the face relative to which the mask should be placed.
	// One of “forehead”, “eyes”, “mouth”, or “chin”.
	Point string `json:"point"`
	// Shift by X-axis measured in widths of the mask scaled to the face size,
	// from left to right. For example, choosing -1.0 will place mask just to
	// the left of the default mask position.
	XShift float64 `json:"x_shift"`
	// Shift by Y-axis measured in heights of the mask scaled to the face size,
	// from top to bottom. For example, 1.0 will place the mask just below the
	// default mask position.
	YShift float64 `json:"y_shift"`
	// Mask scaling coefficient. For example, 2.0 means double size.
	Scale float64 `json:"scale"`
}

// Game represents a game. Use BotFather to create and edit games, their short
// names will act as unique identifiers.
type Game struct {
	// Title of the game
	Title string `json:"title"`
	// Description of the game
	Description string `json:"description"`
	// Photo that will be displayed in the game message in chats.
	Photo []PhotoSize `json:"photo"`
	// Text a brief description of the game or high scores included in the game message.
	// Can be automatically edited to include current high scores for the game
	// when the bot calls setGameScore, or manually edited using editMessageText. 0-4096 characters.
	//
	// optional
	Text string `json:"text,omitempty"`
	// TextEntities special entities that appear in text, such as usernames, URLs, bot commands, etc.
	//
	// optional
	TextEntities []MessageEntity `json:"text_entities,omitempty"`
	// Animation is an animation that will be displayed in the game message in chats.
	// Upload via BotFather (https://t.me/botfather).
	//
	// optional
	Animation Animation `json:"animation,omitempty"`
}

// GameHighScore is a user's score and position on the leaderboard.
type GameHighScore struct {
	// Position in high score table for the game
	Position int `json:"position"`
	// User user
	User User `json:"user"`
	// Score score
	Score int `json:"score"`
}

// CallbackGame is for starting a game in an inline keyboard button.
type CallbackGame struct{}

// WebhookInfo is information about a currently set webhook.
type WebhookInfo struct {
	// URL webhook URL, may be empty if webhook is not set up.
	URL string `json:"url"`
	// HasCustomCertificate true, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// PendingUpdateCount number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`
	// IPAddress is the currently used webhook IP address
	//
	// optional
	IPAddress string `json:"ip_address,omitempty"`
	// LastErrorDate unix time for the most recent error
	// that happened when trying to deliver an update via webhook.
	//
	// optional
	LastErrorDate int `json:"last_error_date,omitempty"`
	// LastErrorMessage error message in human-readable format for the most recent error
	// that happened when trying to deliver an update via webhook.
	//
	// optional
	LastErrorMessage string `json:"last_error_message,omitempty"`
	// LastSynchronizationErrorDate is the unix time of the most recent error that
	// happened when trying to synchronize available updates with Telegram datacenters.
	LastSynchronizationErrorDate int `json:"last_synchronization_error_date,omitempty"`
	// MaxConnections maximum allowed number of simultaneous
	// HTTPS connections to the webhook for update delivery.
	//
	// optional
	MaxConnections int `json:"max_connections,omitempty"`
	// AllowedUpdates is a list of update types the bot is subscribed to.
	// Defaults to all update types
	//
	// optional
	AllowedUpdates []string `json:"allowed_updates,omitempty"`
}

// IsSet returns true if a webhook is currently set.
func (info WebhookInfo) IsSet() bool {
	return info.URL != ""
}

// InlineQuery is a Query from Telegram for an inline request.
type InlineQuery struct {
	// ID unique identifier for this query
	ID string `json:"id"`
	// From sender
	From *User `json:"from"`
	// Query text of the query (up to 256 characters).
	Query string `json:"query"`
	// Offset of the results to be returned, can be controlled by the bot.
	Offset string `json:"offset"`
	// Type of the chat, from which the inline query was sent. Can be either
	// “sender” for a private chat with the inline query sender, “private”,
	// “group”, “supergroup”, or “channel”. The chat type should be always known
	// for requests sent from official clients and most third-party clients,
	// unless the request was sent from a secret chat
	//
	// optional
	ChatType string `json:"chat_type,omitempty"`
	// Location sender location, only for bots that request user location.
	//
	// optional
	Location *Location `json:"location,omitempty"`
}

// ChosenInlineResult is an inline query result chosen by a User
type ChosenInlineResult struct {
	// ResultID the unique identifier for the result that was chosen
	ResultID string `json:"result_id"`
	// From the user that chose the result
	From *User `json:"from"`
	// Location sender location, only for bots that require user location
	//
	// optional
	Location *Location `json:"location,omitempty"`
	// InlineMessageID identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	//
	// optional
	InlineMessageID string `json:"inline_message_id,omitempty"`
	// Query the query that was used to obtain the result
	Query string `json:"query"`
}

// SentWebAppMessage contains information about an inline message sent by a Web App
// on behalf of a user.
type SentWebAppMessage struct {
	// Identifier of the sent inline message. Available only if there is an inline
	// keyboard attached to the message.
	//
	// optional
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// LabeledPrice represents a portion of the price for goods or services.
type LabeledPrice struct {
	// Label portion label
	Label string `json:"label"`
	// Amount price of the product in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json),
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	Amount int `json:"amount"`
}

// Invoice contains basic information about an invoice.
type Invoice struct {
	// Title product name
	Title string `json:"title"`
	// Description product description
	Description string `json:"description"`
	// StartParameter unique bot deep-linking parameter that can be used to generate this invoice
	StartParameter string `json:"start_parameter"`
	// Currency three-letter ISO 4217 currency code
	// (see https://core.telegram.org/bots/payments#supported-currencies)
	Currency string `json:"currency"`
	// TotalAmount total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json
	// (https://core.telegram.org/bots/payments/currencies.json),
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}

// ShippingAddress represents a shipping address.
type ShippingAddress struct {
	// CountryCode ISO 3166-1 alpha-2 country code
	CountryCode string `json:"country_code"`
	// State if applicable
	State string `json:"state"`
	// City city
	City string `json:"city"`
	// StreetLine1 first line for the address
	StreetLine1 string `json:"street_line1"`
	// StreetLine2 second line for the address
	StreetLine2 string `json:"street_line2"`
	// PostCode address post code
	PostCode string `json:"post_code"`
}

// OrderInfo represents information about an order.
type OrderInfo struct {
	// Name user name
	//
	// optional
	Name string `json:"name,omitempty"`
	// PhoneNumber user's phone number
	//
	// optional
	PhoneNumber string `json:"phone_number,omitempty"`
	// Email user email
	//
	// optional
	Email string `json:"email,omitempty"`
	// ShippingAddress user shipping address
	//
	// optional
	ShippingAddress *ShippingAddress `json:"shipping_address,omitempty"`
}

// ShippingOption represents one shipping option.
type ShippingOption struct {
	// ID shipping option identifier
	ID string `json:"id"`
	// Title option title
	Title string `json:"title"`
	// Prices list of price portions
	Prices []LabeledPrice `json:"prices"`
}

/**
SuccessfulPayment
This object contains basic information about a successful payment. Note that if the buyer initiates a chargeback with the relevant payment provider following this transaction, the funds may be debited from your balance. This is outside of Telegram's control.

Field	Type	Description
currency	String	Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars
total_amount	Integer	Total price in the smallest units of the currency (integer, not float/double). For example, for a price of US$ 1.45 pass amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
invoice_payload	String	Bot-specified invoice payload
subscription_expiration_date	Integer	Optional. Expiration date of the subscription, in Unix time; for recurring payments only
is_recurring	True	Optional. True, if the payment is a recurring payment for a subscription
is_first_recurring	True	Optional. True, if the payment is the first payment for a subscription
shipping_option_id	String	Optional. Identifier of the shipping option chosen by the user
order_info	OrderInfo	Optional. Order information provided by the user
telegram_payment_charge_id	String	Telegram payment identifier
provider_payment_charge_id	String	Provider payment identifier

*/
// SuccessfulPayment contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Currency three-letter ISO 4217 currency code
	// (see https://core.telegram.org/bots/payments#supported-currencies)
	Currency string `json:"currency"`
	// TotalAmount total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45 pass amount = 145.
	// See the exp parameter in currencies.json,
	// (https://core.telegram.org/bots/payments/currencies.json)
	// it shows the number of digits past the decimal point
	// for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// InvoicePayload bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// ShippingOptionID identifier of the shipping option chosen by the user
	//
	// optional
	ShippingOptionID string `json:"shipping_option_id,omitempty"`
	// OrderInfo order info provided by the user
	//
	// optional
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
	// TelegramPaymentChargeID telegram payment identifier
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	// ProviderPaymentChargeID provider payment identifier
	ProviderPaymentChargeID string `json:"provider_payment_charge_id"`

	// SubscriptionExpirationDate expiration date of the subscription, in Unix time; for recurring payments only
	SubscriptionExpirationDate int `json:"subscription_expiration_date,omitempty"`
	// IsRecurring true, if the payment is a recurring payment for a subscription
	IsRecurring bool `json:"is_recurring,omitempty"`
	// IsFirstRecurring true, if the payment is the first payment for a subscription
	IsFirstRecurring bool `json:"is_first_recurring,omitempty"`
}

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	// ID unique query identifier
	ID string `json:"id"`
	// From user who sent the query
	From *User `json:"from"`
	// InvoicePayload bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// ShippingAddress user specified shipping address
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}

// PreCheckoutQuery contains information about an incoming pre-checkout query.
type PreCheckoutQuery struct {
	// ID unique query identifier
	ID string `json:"id"`
	// From user who sent the query
	From *User `json:"from"`
	// Currency three-letter ISO 4217 currency code
	//	// (see https://core.telegram.org/bots/payments#supported-currencies)
	Currency string `json:"currency"`
	// TotalAmount total price in the smallest units of the currency (integer, not float/double).
	//	// For example, for a price of US$ 1.45 pass amount = 145.
	//	// See the exp parameter in currencies.json,
	//	// (https://core.telegram.org/bots/payments/currencies.json)
	//	// it shows the number of digits past the decimal point
	//	// for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// InvoicePayload bot specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// ShippingOptionID identifier of the shipping option chosen by the user
	//
	// optional
	ShippingOptionID string `json:"shipping_option_id,omitempty"`
	// OrderInfo order info provided by the user
	//
	// optional
	OrderInfo *OrderInfo `json:"order_info,omitempty"`
}

/*
Describes the connection of the bot with a business account.

Field	Type	Description
id	String	Unique identifier of the business connection
user	User	Business account user that created the business connection
user_chat_id	Integer	Identifier of a private chat with the user who created the business connection. This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
date	Integer	Date the connection was established in Unix time
rights	BusinessBotRights	Optional. Rights of the business bot
is_enabled	Boolean	True, if the connection is active
*/

type BusinessConnection struct {
	// Unique identifier of the business connection
	Id string `json:"id"`
	// Business account user that created the business connection
	User *User `json:"user,omitempty"`
	// Identifier of a private chat with the user who created the business connection.
	// This number may have more than 32 significant bits and some programming languages may
	// have difficulty/silent defects in interpreting it. But it has at most 52 significant bits,
	// so a 64-bit integer or double-precision float type are safe for storing this identifier.
	UserChatId int `json:"user_chat_id"`
	// Date the connection was established in Unix time
	Date int `json:"date"`
	// True, if the bot can act on behalf of the business account in chats that were active in the last 24 hours
	CanReply bool `json:"can_reply"`
	// True, if the connection is active
	IsEnabled bool `json:"is_enabled"`
	// Rights of the business bot
	Rights *BusinessBotRights `json:"rights,omitempty"`
}

type BusinessMessagesDeleted struct {
	// Unique identifier of the business connection
	BusinessConnectionId string `json:"business_connection_id"`
	// Information about a chat in the business account. The bot may not have access to the chat or the corresponding user.
	Chat *Chat `json:"chat,omitempty"`
	// The list of identifiers of deleted messages in the chat of the business account
	MessageIds []int64 `json:"message_ids"`
}

type MessageReactionUpdated struct {
	// The chat containing the message the user reacted to
	Chat *Chat `json:"chat,omitempty"`
	// Unique identifier of the message inside the chat
	MessageId int64 `json:"message_id"`
	// The user that changed the reaction, if the user isn't anonymous
	//
	// optional
	User *User `json:"user,omitempty"`
	// The chat on behalf of which the reaction was changed, if the user is anonymous
	//
	// optional
	ActorChat *Chat `json:"actor_chat,omitempty"`
	// Date of the change in Unix time
	Date int `json:"date"`
	// Previous list of reaction types that were set by the user
	OldReaction []ReactionType `json:"old_reaction"`
	// New list of reaction types that have been set by the user
	NewReaction []ReactionType `json:"new_reaction"`
}

type MessageReactionCountUpdated struct {
	// The chat containing the message
	Chat *Chat `json:"chat,omitempty"`
	// Unique message identifier inside the chat
	MessageId int64 `json:"message_id"`
	// Date of the change in Unix time
	Date int `json:"date"`
	// List of reactions that are present on the message
	Reactions []ReactionCount `json:"reactions"`
}

type ChatBoostSource struct {
	// Source of the boost,  “premium” or “gift_code” or “giveaway”
	Source string `json:"source"`
	// User that boosted the chat
	User *User `json:"user,omitempty"`
	// Identifier of a message in the chat with the giveaway; the message could have been deleted already.
	// May be 0 if the message isn't sent yet.
	GiveawayMessageId int64 `json:"giveaway_message_id"`
	// True, if the giveaway was completed, but there was no user to win the prize
	IsUnclaimed bool `json:"is_unclaimed"`
}

type ChatBoost struct {
	// Unique identifier of the boost
	BoostId string `json:"boost_id"`
	// Point in time (Unix timestamp) when the chat was boosted
	AddDate int `json:"add_date"`
	// Point in time (Unix timestamp) when the boost will automatically expire,
	// unless the booster's Telegram Premium subscription is prolonged
	ExpirationDate int `json:"expiration_date"`
	// Source of the added boost
	Source *ChatBoostSource `json:"source,omitempty"`
}

type ChatBoostUpdated struct {
	// Chat which was boosted
	Chat *Chat `json:"chat,omitempty"`
	// Information about the chat boost
	Boost *ChatBoost `json:"boost,omitempty"`
}

type ChatBoostRemoved struct {
	// Chat which was boosted
	Chat *Chat `json:"chat,omitempty"`
	// Unique identifier of the boost
	BoostId string `json:"boost_id"`
	// Point in time (Unix timestamp) when the boost was removed
	RemoveDate int `json:"remove_date"`
	// Source of the removed boost
	Source ChatBoostSource `json:"source"`
}

type MessageOrigin struct {
	// Type of the message origin, “user” or “hidden_user” or “chat” or “channel”
	Type string `json:"type"`
	// Date the message was sent originally in Unix time
	Date int `json:"date"`
	// User that sent the message originally
	SenderUser *User `json:"sender_user"`
	// Chat that sent the message originally
	SenderChat *Chat `json:"sender_chat"`

	//Name of the user that sent the message originally
	SenderUserName string `json:"sender_user_name,omitempty"`

	// For messages originally sent by an anonymous chat administrator, original message author signature
	AuthorSignature string `json:"author_signature"`
	// Channel chat to which the message was originally sent
	Chat *Chat `json:"chat"`
	// Unique message identifier inside the chat
	MessageId int64 `json:"message_id"`
}

type ExternalReplyInfo struct {
	// Origin of the message replied to by the given message
	//
	// optional
	Origin *MessageOrigin `json:"origin"`
	// Chat the original message belongs to. Available only if the chat is a supergroup or a channel.
	//
	// optional
	Chat *Chat `json:"chat"`
	// Unique message identifier inside the original chat. Available only if the original chat is a supergroup or a channel.
	//
	// optional
	MessageId int64 `json:"message_id"`
	// Options used for link preview generation for the original message, if it is a text message
	//
	// optional
	LinkPreviewOptions *LinkPreviewOptions `json:"link_preview_options"`
	// Message is an animation, information about the animation
	//
	// optional
	Animation *Animation `json:"animation"`
	// Message is an audio file, information about the file
	//
	// optional
	Audio *Audio `json:"audio"`
	// Message is a general file, information about the file
	//
	// optional
	Document *Document `json:"document"`
	// Message contains paid media; information about the paid media
	//
	// optional
	PaidMedia *PaidMediaInfo `json:"paid_media"`
	// Message is a photo, available sizes of the photo
	//
	// optional
	Photo []PhotoSize `json:"photo"`
	// Message is a sticker, information about the sticker
	//
	// optional
	Sticker *Sticker `json:"sticker"`
	// Message is a forwarded story
	//
	// optional
	Story *Story `json:"story"`
	// Message is a video, information about the video
	//
	// optional
	Video *Video `json:"video"`
	// Message is a video note, information about the video message
	//
	// optional
	VideoNote *VideoNote `json:"video_note"`
	// Message is a voice message, information about the file
	//
	// optional
	Voice *Voice `json:"voice"`
	// True, if the message media is covered by a spoiler animation
	//
	// optional
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Message is a checklist
	//
	// optional
	Checklist *Checklist `json:"checklist"`

	// optional
	Contact *Contact `json:"contact"`
	// Message is a dice with random value
	//
	// optional
	Dice *Dice `json:"dice"`
	// Message is a game, information about the game.
	//
	// optional
	Game *Game `json:"game"`
	// Message is a scheduled giveaway, information about the giveaway
	//
	// optional
	Giveaway *Giveaway `json:"giveaway"`
	// A giveaway with public winners was completed
	//
	// optional
	GiveawayWinners *GiveawayWinners `json:"giveaway_winners"`
	// Message is an invoice for a payment, information about the invoice.
	//
	// optional
	Invoice *Invoice `json:"invoice"`
	// Message is a shared location, information about the location
	//
	// optional
	Location *Location `json:"location"`
	// Message is a native poll, information about the poll
	//
	// optional
	Poll *Poll `json:"poll"`
	// Message is a venue, information about the venue
	//
	// optional
	Venue *Venue `json:"venue"`
}

type TextQuote struct {
	// Text of the quoted part of a message that is replied to by the given message
	Text string `json:"text"`
	// Special entities that appear in the quote. Currently, only bold, italic, underline, strikethrough, spoiler,
	// and custom_emoji entities are kept in quotes.
	//
	// optional
	Entities []MessageEntity `json:"entities"`
	// Approximate quote position in the original message in UTF-16 code units as specified by the sender
	Position int `json:"position"`
	// True, if the quote was chosen manually by the message sender. Otherwise, the quote was added automatically by the server.
	//
	// optional
	IsManual bool `json:"is_manual"`
}

type Story struct {
	// Chat that posted the story
	Chat *Chat `json:"chat"`
	// Unique identifier for the story in the chat
	Id int64 `json:"id"`
}

type LinkPreviewOptions struct {
	// True, if the link preview is disabled
	//
	// optional
	IsDisabled bool `json:"is_disabled"`
	// URL to use for the link preview. If empty, then the first URL found in the message text will be used
	//
	// optional
	Url string `json:"url"`
	// True, if the media in the link preview is supposed to be shrunk;
	// ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	//
	// optional
	PreferSmallMedia bool `json:"prefer_small_media"`
	// True, if the media in the link preview is supposed to be enlarged;
	// ignored if the URL isn't explicitly specified or media size change isn't supported for the preview
	//
	// optional
	PreferLargeMedia bool `json:"prefer_large_media"`
	// True, if the link preview must be shown above the message text; otherwise, the link preview will be shown below the message text
	//
	// optional
	ShowAboveText bool `json:"show_above_text"`
}

type PaidMediaInfo struct {
	// The number of Telegram Stars that must be paid to buy access to the media
	StarCount int `json:"star_count"`
	// Information about the paid media
	PaidMedia []PaidMedia `json:"paid_media"`
}

type PaidMedia struct {
	// Type of the paid media,  “preview” or “photo” or “video”
	Type string `json:"type"`
	// Media width as defined by the sender
	//
	// optional
	Width int `json:"width"`
	// Media height as defined by the sender
	//
	// optional
	Height int `json:"height"`
	// Duration of the media in seconds as defined by the sender
	//
	// optional
	Duration int `json:"duration"`
	// The photo
	Photo []PhotoSize `json:"photo"`
	// The Video
	Video *Video `json:"video"`
}

type RefundedPayment struct {
	// Three-letter ISO 4217 currency code, or “XTR” for payments in Telegram Stars. Currently, always “XTR”
	Currency string `json:"currency"`
	// Total refunded price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.45, total_amount = 145. See the exp parameter in currencies.json,
	// it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot-specified invoice payload
	InvoicePayload string `json:"invoice_payload"`
	// Telegram payment identifier
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Provider payment identifier
	//
	// optional
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
}

type SharedUser struct {
	// Identifier of the shared user. This number may have more than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits,
	// so 64-bit integers or double-precision float types are safe for storing these identifiers.
	// The bot may not have access to the user and could be unable to use this identifier,
	// unless the user is already known to the bot by some other means.
	UserId int64 `json:"user_id"`
	// First name of the user, if the name was requested by the bot
	//
	// optional
	FirstName string `json:"first_name"`
	// Last name of the user, if the name was requested by the bot
	//
	// optional
	LastName string `json:"last_name"`
	// Username of the user, if the username was requested by the bot
	//
	// optional
	UserName string `json:"username"`
	// Available sizes of the chat photo, if the photo was requested by the bot
	//
	// optional
	Photo []PhotoSize `json:"photo"`
}

type UsersShared struct {
	// Identifier of the request
	RequestId int `json:"request_id"`
	// Information about users shared with the bot.
	Users []SharedUser `json:"users"`
}

type ChatShared struct {
	// Identifier of the request
	RequestId int `json:"request_id"`
	// Identifier of the shared chat. This number may have more than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it. But it has at most 52 significant bits,
	// so a 64-bit integer or double-precision float type are safe for storing this identifier.
	// The bot may not have access to the chat and could be unable to use this identifier,
	// unless the chat is already known to the bot by some other means.
	ChatId int64 `json:"chat_id"`
	// Title of the chat, if the title was requested by the bot.
	//
	// optional
	Title string `json:"title"`
	// Username of the chat, if the username was requested by the bot and available.
	//
	// optional
	UserName string `json:"username"`
	// Available sizes of the chat photo, if the photo was requested by the bot
	//
	// optional
	Photo []PhotoSize `json:"photo"`
}

type WriteAccessAllowed struct {
	// True, if the access was granted after the user accepted an explicit request from a Web App sent by the method requestWriteAccess
	//
	// optional
	FromRequest bool `json:"from_request"`
	// Name of the Web App, if the access was granted when the Web App was launched from a link
	//
	// optional
	WebAppName string `json:"web_app_name"`
	// True, if the access was granted when the bot was added to the attachment or side menu
	//
	// optional
	FromAttachmentMenu bool `json:"from_attachment_menu"`
}

type ChatBoostAdded struct {
	// Number of boosts added by the user
	BoostCount int `json:"boost_count"`
}

type BackgroundFill struct {
	// Type of the background fill, “solid” or "gradient" or "freeform_gradient"
	Type string `json:"type"`
	// The color of the background fill in the RGB24 format
	Color int `json:"color"`
	// Top color of the gradient in the RGB24 format
	TopColor int `json:"top_color"`
	// Bottom color of the gradient in the RGB24 format
	BottomColor int `json:"bottom_color"`
	// Clockwise rotation angle of the background fill in degrees; 0-359
	RotationAngle int `json:"rotation_angle"`
	// A list of the 3 or 4 base colors that are used to generate the freeform gradient in the RGB24 format
	Colors []int `json:"colors"`
}

type BackgroundType struct {
	// Type of the background, “fill” or "wallpaper" or "pattern" or "chat_theme"
	Type string `json:"type"`
	// The background fill
	Fill *BackgroundFill `json:"fill"`
	// Dimming of the background in dark themes, as a percentage; 0-100
	DarkThemeDimming int `json:"dark_theme_dimming"`
	// Document with the wallpaper
	Document *Document `json:"document"`
	// True, if the wallpaper is downscaled to fit in a 450x450 square and then box-blurred with radius 12
	//
	// optional
	IsBlurred bool `json:"is_blurred"`
	// True, if the background moves slightly when the device is tilted
	//
	// optional
	IsMoving bool `json:"is_moving"`
	// Intensity of the pattern when it is shown above the filled background; 0-100
	Intensity int `json:"intensity"`
	// True, if the background fill must be applied only to the pattern itself. All other pixels are black in this case. For dark themes only
	//
	// optional
	IsInverted bool `json:"is_inverted"`
}

type ChatBackground struct {
	// Type of the background
	Type *BackgroundType `json:"type,omitempty"`
}

type ForumTopicCreated struct {
	// Name of the topic
	Name string `json:"name"`
	// Color of the topic icon in RGB format
	IconColor int `json:"icon_color"`
	// Unique identifier of the custom emoji shown as the topic icon
	//
	// optional
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type ForumTopicEdited struct {
	// Name of the topic
	Name string `json:"name"`
	// Unique identifier of the custom emoji shown as the topic icon
	//
	// optional
	IconCustomEmojiId string `json:"icon_custom_emoji_id"`
}

type ForumTopicClosed struct {
	// This object represents a service message about a forum topic closed in the chat. Currently holds no information.
}

type ForumTopicReopened struct {
	// This object represents a service message about a forum topic reopened in the chat. Currently holds no information.
}

type GeneralForumTopicHidden struct {
	// This object represents a service message about General forum topic hidden in the chat. Currently holds no information.
}

type GeneralForumTopicUnhidden struct {
	// This object represents a service message about General forum topic unhidden in the chat. Currently holds no information.
	//prize_star_count
	//Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count"`
}

// This object represents a service message about the creation of a scheduled giveaway.
type GiveawayCreated struct {
	//prize_star_count	Integer	Optional. The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	PrizeStarCount int `json:"prize_star_count"`
}

type Giveaway struct {
	// The list of chats which the user must join to participate in the giveaway
	Chats []Chat `json:"chats"`
	// Point in time (Unix timestamp) when winners of the giveaway will be selected
	WinnersSelectionDate int `json:"winners_selection_date"`
	// The number of users which are supposed to be selected as winners of the giveaway
	WinnerCount int `json:"winner_count"`
	// True, if only users who join the chats after the giveaway started should be eligible to win
	//
	// optional
	OnlyNewMembers bool `json:"only_new_members"`
	// True, if the list of giveaway winners will be visible to everyone
	//
	// optional
	HasPublicWinners bool `json:"has_public_winners"`
	// Description of additional giveaway prize
	//
	// optional
	PrizeDescription string `json:"prize_description"`
	// A list of two-letter ISO 3166-1 alpha-2 country codes indicating the countries from
	// which eligible users for the giveaway must come. If empty, then all users can participate in the giveaway.
	// Users with a phone number that was bought on Fragment can always participate in giveaways.
	//
	// optional
	CountryCodes []string `json:"country_codes"`
	// The number of months the Telegram Premium subscription won from the giveaway will be active for
	//
	// optional
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count"`

	//The number of Telegram Stars to be split between giveaway winners; for Telegram Star giveaways only
	//
	// optional
	PrizeStarCount int `json:"prize_star_count"`
}

type GiveawayWinners struct {
	// The chat that created the giveaway
	Chat *Chat `json:"chat"`
	// Identifier of the message with the giveaway in the chat
	GiveawayMessageId int64 `json:"giveaway_message_id"`
	// Point in time (Unix timestamp) when winners of the giveaway were selected
	WinnersSelectionDate int `json:"winners_selection_date"`
	// Total number of winners in the giveaway
	WinnerCount int `json:"winner_count"`
	// List of up to 100 winners of the giveaway
	Winners []User `json:"winners"`
	// The number of other chats the user had to join in order to be eligible for the giveaway
	//
	// optional
	AdditionalChatCount int `json:"additional_chat_count"`
	// The number of months the Telegram Premium subscription won from the giveaway will be active for
	//
	// optional
	PremiumSubscriptionMonthCount int `json:"premium_subscription_month_count"`
	// Number of undistributed prizes
	//
	// optional
	UnclaimedPrizeCount int `json:"unclaimed_prize_count"`
	// True, if only users who had joined the chats after the giveaway started were eligible to win
	//
	// optional
	OnlyNewMembers bool `json:"only_new_members"`
	// True, if the giveaway was canceled because the payment for it was refunded
	//

	// optional
	WasRefunded bool `json:"was_refunded"`
	// Description of additional giveaway prize
	//
	// optional
	PrizeDescription string `json:"prize_description"`

	//prize_star_count
	//Optional. The number of Telegram Stars that were split between giveaway winners; for Telegram Star giveaways only
	//
	// optional
	PrizeStarCount int `json:"prize_star_count"`
}

type GiveawayCompleted struct {
	// Number of winners in the giveaway
	WinnerCount int `json:"winner_count"`
	// Number of undistributed prizes
	//
	// optional
	UnclaimedPrizeCount int `json:"unclaimed_prize_count"`
	// Message with the giveaway that was completed, if it wasn't deleted
	//
	// optional
	GiveawayMessage *Message `json:"giveaway_message"`
	//
	//is_star_giveaway	True	Optional. True, if the giveaway is a Telegram Star giveaway. Otherwise, currently, the giveaway is a Telegram Premium giveaway.
	IsStarGiveaway bool `json:"is_star_giveaway"`
}

type ReactionType struct {
	// Type of the reaction,  “emoji” or “custom_emoji” "paid"
	Type string `json:"type"`
	// Emoji
	Emoji string `json:"emoji"`
	// Custom emoji identifier
	CustomEmojiId string `json:"custom_emoji_id"`
}

type ReactionCount struct {
	// Type of the reaction
	Type ReactionType `json:"type"`
	// Number of times the reaction was added
	TotalCount int `json:"total_count"`
}

/*
*
PaidMediaPurchased
This object contains information about a paid media purchase.
Field	Type	Description
from	User	User who purchased the media
paid_media_payload	String	Bot-specified paid media payload
*/
type PaidMediaPurchased struct {
	From             *User  `json:"from"`
	PaidMediaPayload string `json:"paid_media_payload"`
}

/*
*
PaidMessagePriceChanged
Describes a service message about a change in the price of paid messages within a chat.
Field	Type	Description
paid_message_star_count	Integer	The new number of Telegram Stars that must be paid by non-administrator users of the supergroup chat for each sent message
*/
type PaidMessagePriceChanged struct {
	PaidMessageStarCount int `json:"paid_message_star_count"`
}

/*
BusinessBotRights
Represents the rights of a business bot.

Field	Type	Description
can_reply	True	Optional. True, if the bot can send and edit messages in the private chats that had incoming messages in the last 24 hours
can_read_messages	True	Optional. True, if the bot can mark incoming private messages as read
can_delete_outgoing_messages	True	Optional. True, if the bot can delete messages sent by the bot
can_delete_all_messages	True	Optional. True, if the bot can delete all private messages in managed chats
can_edit_name	True	Optional. True, if the bot can edit the first and last name of the business account
can_edit_bio	True	Optional. True, if the bot can edit the bio of the business account
can_edit_profile_photo	True	Optional. True, if the bot can edit the profile photo of the business account
can_edit_username	True	Optional. True, if the bot can edit the username of the business account
can_change_gift_settings	True	Optional. True, if the bot can change the privacy settings pertaining to gifts for the business account
can_view_gifts_and_stars	True	Optional. True, if the bot can view gifts and the amount of Telegram Stars owned by the business account
can_convert_gifts_to_stars	True	Optional. True, if the bot can convert regular gifts owned by the business account to Telegram Stars
can_transfer_and_upgrade_gifts	True	Optional. True, if the bot can transfer and upgrade gifts owned by the business account
can_transfer_stars	True	Optional. True, if the bot can transfer Telegram Stars received by the business account to its own account, or use them to upgrade and transfer gifts
can_manage_stories	True	Optional. True, if the bot can post, edit and delete stories on behalf of the business account
*/
type BusinessBotRights struct {
	CanReply                   bool `json:"can_reply"`
	CanReadMessages            bool `json:"can_read_messages"`
	CanDeleteSentMessages      bool `json:"can_delete_sent_messages"`
	CanDeleteAllMessages       bool `json:"can_delete_all_messages"`
	CanEditName                bool `json:"can_edit_name"`
	CanEditBio                 bool `json:"can_edit_bio"`
	CanEditProfilePhoto        bool `json:"can_edit_profile_photo"`
	CanEditUsername            bool `json:"can_edit_username"`
	CanChangeGiftSettings      bool `json:"can_change_gift_settings"`
	CanViewGiftsAndStars       bool `json:"can_view_gifts_and_stars"`
	CanConvertGiftsToStars     bool `json:"can_convert_gifts_to_stars"`
	CanTransferAndUpgradeGifts bool `json:"can_transfer_and_upgrade_gifts"`
	CanTransferStars           bool `json:"can_transfer_stars"`
	CanManageStories           bool `json:"can_manage_stories"`
	CanDeleteOutgoingMessages  bool `json:"can_delete_outgoing_messages"`
}

/*
ReplyParameters
Describes reply parameters for the message that is being sent.

Field	Type	Description
message_id	Integer	Identifier of the message that will be replied to in the current chat, or in the chat chat_id if it is specified
chat_id	Integer or String	Optional. If the message to be replied to is from a different chat, unique identifier for the chat or username of the channel (in the format @channelusername). Not supported for messages sent on behalf of a business account.
allow_sending_without_reply	Boolean	Optional. Pass True if the message should be sent even if the specified message to be replied to is not found. Always False for replies in another chat or forum topic. Always True for messages sent on behalf of a business account.
quote	String	Optional. Quoted part of the message to be replied to; 0-1024 characters after entities parsing. The quote must be an exact substring of the message to be replied to, including bold, italic, underline, strikethrough, spoiler, and custom_emoji entities. The message will fail to send if the quote isn't found in the original message.
quote_parse_mode	String	Optional. Mode for parsing entities in the quote. See formatting options for more details.
quote_entities	Array of MessageEntity	Optional. A JSON-serialized list of special entities that appear in the quote. It can be specified instead of quote_parse_mode.
quote_position	Integer	Optional. Position of the quote in the original message in UTF-16 code units
*/

type ReplyParameters struct {
	MessageId                int             `json:"message_id"`
	ChatId                   string          `json:"chat_id"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply"`
	Quote                    string          `json:"quote"`
	QuoteParseMode           string          `json:"quote_parse_mode"`
	QuoteEntities            []MessageEntity `json:"quote_entities"`
	QuotePosition            int             `json:"quote_position"`
}

/*
DirectMessagePriceChanged
Describes a service message about a change in the price of direct messages sent to a channel chat.

Field	Type	Description
are_direct_messages_enabled	Boolean	True, if direct messages are enabled for the channel chat; false otherwise
direct_message_star_count	Integer	Optional. The new number of Telegram Stars that must be paid by users for each direct message sent to the channel. Does not apply to users who have been exempted by administrators. Defaults to 0.
*/
type DirectMessagePriceChanged struct {
	AreDirectMessagesEnabled bool `json:"are_direct_messages_enabled"`
	DirectMessageStarCount   int  `json:"direct_message_star_count"`
}
