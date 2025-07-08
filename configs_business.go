package tgbotapi

/*

removeChatVerification
Removes verification from a chat that is currently verified on behalf of the organization represented by the bot. Returns True on success.

Parameter	Type	Required	Description
chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
readBusinessMessage
Marks incoming message as read on behalf of a business account. Requires the can_read_messages business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection on behalf of which to read the message
chat_id	Integer	Yes	Unique identifier of the chat in which the message was received. The chat must have been active in the last 24 hours.
message_id	Integer	Yes	Unique identifier of the message to mark as read
deleteBusinessMessages
Delete messages on behalf of a business account. Requires the can_delete_sent_messages business bot right to delete messages sent by the bot itself, or the can_delete_all_messages business bot right to delete any message. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection on behalf of which to delete the messages
message_ids	Array of Integer	Yes	A JSON-serialized list of 1-100 identifiers of messages to delete. All messages must be from the same chat. See deleteMessage for limitations on which messages can be deleted
setBusinessAccountName
Changes the first and last name of a managed business account. Requires the can_change_name business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
first_name	String	Yes	The new value of the first name for the business account; 1-64 characters
last_name	String	Optional	The new value of the last name for the business account; 0-64 characters
setBusinessAccountUsername
Changes the username of a managed business account. Requires the can_change_username business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
username	String	Optional	The new value of the username for the business account; 0-32 characters
setBusinessAccountBio
Changes the bio of a managed business account. Requires the can_change_bio business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
bio	String	Optional	The new value of the bio for the business account; 0-140 characters
setBusinessAccountProfilePhoto
Changes the profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
photo	InputProfilePhoto	Yes	The new profile photo to set
is_public	Boolean	Optional	Pass True to set the public photo, which will be visible even if the main photo is hidden by the business account's privacy settings. An account can have only one public photo.
removeBusinessAccountProfilePhoto
Removes the current profile photo of a managed business account. Requires the can_edit_profile_photo business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
is_public	Boolean	Optional	Pass True to remove the public photo, which is visible even if the main photo is hidden by the business account's privacy settings. After the main photo is removed, the previous profile photo (if present) becomes the main photo.
setBusinessAccountGiftSettings
Changes the privacy settings pertaining to incoming gifts in a managed business account. Requires the can_change_gift_settings business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
show_gift_button	Boolean	Yes	Pass True, if a button for sending a gift to the user or by the business account must always be shown in the input field
accepted_gift_types	AcceptedGiftTypes	Yes	Types of gifts accepted by the business account
getBusinessAccountStarBalance
Returns the amount of Telegram Stars owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns StarAmount on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
transferBusinessAccountStars
Transfers Telegram Stars from the business account balance to the bot's balance. Requires the can_transfer_stars business bot right. Returns True on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
star_count	Integer	Yes	Number of Telegram Stars to transfer; 1-10000
getBusinessAccountGifts
Returns the gifts received and owned by a managed business account. Requires the can_view_gifts_and_stars business bot right. Returns OwnedGifts on success.

Parameter	Type	Required	Description
business_connection_id	String	Yes	Unique identifier of the business connection
exclude_unsaved	Boolean	Optional	Pass True to exclude gifts that aren't saved to the account's profile page
exclude_saved	Boolean	Optional	Pass True to exclude gifts that are saved to the account's profile page
exclude_unlimited	Boolean	Optional	Pass True to exclude gifts that can be purchased an unlimited number of times
exclude_limited	Boolean	Optional	Pass True to exclude gifts that can be purchased a limited number of times
exclude_unique	Boolean	Optional	Pass True to exclude unique gifts
sort_by_price	Boolean	Optional	Pass True to sort results by gift price instead of send date. Sorting is applied before pagination.
offset	String	Optional	Offset of the first entry to return as received from the previous request; use empty string to get the first chunk of results
limit	Integer	Optional	The maximum number of gifts to be returned; 1-100. Defaults to 100

*/

type InputProfilePhoto struct {
	Type               string
	Photo              string
	Animation          string
	MainFrameTimestamp float32
}

/*
*
unlimited_gifts	Boolean	True, if unlimited regular gifts are accepted
limited_gifts	Boolean	True, if limited regular gifts are accepted
unique_gifts	Boolean	True, if unique gifts or gifts that can be upgraded to unique for free are accepted
premium_subscription	Boolean	True, if a Telegram Premium subscription is accepted
*/
type AcceptedGiftTypes struct {
	UnlimitedGifts      bool
	LimitedGifts        bool
	UniqueGifts         bool
	PremiumSubscription bool
}

type RemoveChatVerificationConfig struct {
	ChatID string
}

func (config RemoveChatVerificationConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("chat_id", config.ChatID)
	return params, nil
}

func (config RemoveChatVerificationConfig) method() string {
	return "removeChatVerification"
}

type ReadBusinessMessageConfig struct {
	BusinessConnectionID string
	ChatID               string
	MessageID            int
}

func (config ReadBusinessMessageConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("chat_id", config.ChatID)
	params.AddNonZero("message_id", config.MessageID)
	return params, nil
}

func (config ReadBusinessMessageConfig) method() string {
	return "readBusinessMessage"
}

type DeleteBusinessMessagesConfig struct {
	BusinessConnectionID string
	MessageIDs           []int
}

func (config DeleteBusinessMessagesConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddInterface("message_ids", config.MessageIDs)
	return params, nil
}

func (config DeleteBusinessMessagesConfig) method() string {
	return "deleteBusinessMessages"
}

type SetBusinessAccountNameConfig struct {
	BusinessConnectionID string
	FirstName            string
	LastName             string
}

func (config SetBusinessAccountNameConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("first_name", config.FirstName)
	params.AddNonEmpty("last_name", config.LastName)
	return params, nil
}

func (config SetBusinessAccountNameConfig) method() string {
	return "setBusinessAccountName"
}

type SetBusinessAccountUsernameConfig struct {
	BusinessConnectionID string
	Username             string
}

func (config SetBusinessAccountUsernameConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("username", config.Username)
	return params, nil
}

func (config SetBusinessAccountUsernameConfig) method() string {
	return "setBusinessAccountUsername"
}

type SetBusinessAccountBioConfig struct {
	BusinessConnectionID string
	Bio                  string
}

func (config SetBusinessAccountBioConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("bio", config.Bio)
	return params, nil
}

func (config SetBusinessAccountBioConfig) method() string {
	return "setBusinessAccountBio"
}

type SetBusinessAccountProfilePhotoConfig struct {
	BusinessConnectionID string
	Photo                InputProfilePhoto
	IsPublic             bool
}

func (config SetBusinessAccountProfilePhotoConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddInterface("photo", config.Photo)
	params.AddBool("is_public", config.IsPublic)
	return params, nil
}

type RemoveBusinessAccountProfilePhotoConfig struct {
	BusinessConnectionID string
	IsPublic             bool
}

func (config RemoveBusinessAccountProfilePhotoConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddBool("is_public", config.IsPublic)
	return params, nil
}

func (config RemoveBusinessAccountProfilePhotoConfig) method() string {
	return "removeBusinessAccountProfilePhoto"
}

type SetBusinessAccountGiftSettingsConfig struct {
	BusinessConnectionID string
	ShowGiftButton       bool
	AcceptedGiftTypes    AcceptedGiftTypes
}

func (config SetBusinessAccountGiftSettingsConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddBool("show_gift_button", config.ShowGiftButton)
	params.AddInterface("accepted_gift_types", config.AcceptedGiftTypes)
	return params, nil
}

func (config SetBusinessAccountGiftSettingsConfig) method() string {
	return "setBusinessAccountGiftSettings"
}

type GetBusinessAccountStarBalanceConfig struct {
	BusinessConnectionID string
}

func (config GetBusinessAccountStarBalanceConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	return params, nil
}

func (config GetBusinessAccountStarBalanceConfig) method() string {
	return "getBusinessAccountStarBalance"
}

type GetBusinessAccountGiftsConfig struct {
	BusinessConnectionID string
	ExcludeUnsaved       bool
	ExcludeSaved         bool
	ExcludeUnlimited     bool
	ExcludeLimited       bool
	ExcludeUnique        bool
	SortByPrice          bool
	Offset               string
	Limit                int
}

func (config GetBusinessAccountGiftsConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddBool("exclude_unsaved", config.ExcludeUnsaved)
	params.AddBool("exclude_saved", config.ExcludeSaved)
	params.AddBool("exclude_unlimited", config.ExcludeUnlimited)
	params.AddBool("exclude_limited", config.ExcludeLimited)
	params.AddBool("exclude_unique", config.ExcludeUnique)
	params.AddBool("sort_by_price", config.SortByPrice)
	params.AddNonEmpty("offset", config.Offset)
	params.AddNonZero("limit", config.Limit)
	return params, nil
}

func (config GetBusinessAccountGiftsConfig) method() string {
	return "getBusinessAccountGifts"
}

type TransferBusinessAccountStarsConfig struct {
	BusinessConnectionID string
	StarCount            int
}

func (config TransferBusinessAccountStarsConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonZero("star_count", config.StarCount)
	return params, nil
}

func (config TransferBusinessAccountStarsConfig) method() string {
	return "transferBusinessAccountStars"
}
