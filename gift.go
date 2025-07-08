package tgbotapi

type Gift struct {
	// Unique identifier of the gift
	Id string `json:"id"`
	// The sticker that represents the gift
	Sticker Sticker `json:"sticker"`
	// The number of Telegram Stars that must be paid to send the sticker
	StarCount int `json:"star_count"`
	// Optional. The number of Telegram Stars that must be paid to upgrade the gift to a unique one
	UpgradeStarCount int `json:"upgrade_star_count"`
	// Optional. The total number of the gifts of this type that can be sent; for limited gifts only
	TotalCount int `json:"total_count"`
	// Optional. The number of remaining gifts of this type that can be sent; for limited gifts only
	RemainingCount int `json:"remaining_count"`
}

/**
UniqueGiftModel
This object describes the model of a unique gift.
Field	Type	Description
name	String	Name of the model
sticker	Sticker	The sticker that represents the unique gift
rarity_per_mille	Integer	The number of unique gifts that receive this model for every 1000 gifts upgraded
*/

type UniqueGiftModel struct {
	// Name of the model
	Name string `json:"name"`
	// The sticker that represents the unique gift
	Sticker Sticker `json:"sticker"`
	// The number of unique gifts that receive this model for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

type UniqueGiftBackdropColors struct {
	// The color in the center of the backdrop in RGB format
	CenterColor int `json:"center_color"`
	// The color on the edges of the backdrop in RGB format
	EdgeColor int `json:"edge_color"`
	// The color to be applied to the symbol in RGB format
	SymbolColor int `json:"symbol_color"`
	// The color for the text on the backdrop in RGB format
	TextColor int `json:"text_color"`
}

type UniqueGiftInfo struct {
	// Information about the gift
	Gift UniqueGift `json:"gift"`
	// Origin of the gift. Currently, either “upgrade” or “transfer”
	Origin string `json:"origin"`
	// Optional. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts
	OwnedGiftId string `json:"owned_gift_id"`
	// Optional. Number of Telegram Stars that must be paid to transfer the gift; omitted if the bot cannot transfer the gift
	TransferStarCount int `json:"transfer_star_count"`
}

/*
*
UniqueGift
This object describes a unique gift that was upgraded from a regular gift.

Field	Type	Description
base_name	String	Human-readable name of the regular gift from which this unique gift was upgraded
name	String	Unique name of the gift. This name can be used in https://t.me/nft/... links and story areas
number	Integer	Unique number of the upgraded gift among gifts upgraded from the same regular gift
model	UniqueGiftModel	Model of the gift
symbol	UniqueGiftSymbol	Symbol of the gift
backdrop	UniqueGiftBackdrop	Backdrop of the gift
*/
type UniqueGift struct {
	// Human-readable name of the regular gift from which this unique gift was upgraded
	BaseName string `json:"base_name"`
	// Unique name of the gift. This name can be used in https://t.me/nft/... links and story areas
	Name string `json:"name"`
	// Unique number of the upgraded gift among gifts upgraded from the same regular gift
	Number int `json:"number"`
	// Model of the gift
	Model UniqueGiftModel `json:"model"`
	// Symbol of the gift
	Symbol UniqueGiftSymbol `json:"symbol"`
	// Backdrop of the gift
	Backdrop UniqueGiftBackdrop `json:"backdrop"`
}

/*
UniqueGiftSymbol
This object describes the symbol shown on the pattern of a unique gift.

Field	Type	Description
name	String	Name of the symbol
sticker	Sticker	The sticker that represents the unique gift
rarity_per_mille	Integer	The number of unique gifts that receive this model for every 1000 gifts upgraded
*/
type UniqueGiftSymbol struct {
	// Name of the symbol
	Name string `json:"name"`
	// The sticker that represents the unique gift
	Sticker Sticker `json:"sticker"`
	// The number of unique gifts that receive this model for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

/*
UniqueGiftBackdrop
This object describes the backdrop of a unique gift.

Field	Type	Description
name	String	Name of the backdrop
colors	UniqueGiftBackdropColors	Colors of the backdrop
rarity_per_mille	Integer	The number of unique gifts that receive this backdrop for every 1000 gifts upgraded
*/
type UniqueGiftBackdrop struct {
	// Name of the backdrop
	Name string `json:"name"`
	// Colors of the backdrop
	Colors UniqueGiftBackdropColors `json:"colors"`
	// The number of unique gifts that receive this backdrop for every 1000 gifts upgraded
	RarityPerMille int `json:"rarity_per_mille"`
}

/*
GiftInfo
Describes a service message about a regular gift that was sent or received.

Field	Type	Description
gift	Gift	Information about the gift
owned_gift_id	String	Optional. Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts
convert_star_count	Integer	Optional. Number of Telegram Stars that can be claimed by the receiver by converting the gift; omitted if conversion to Telegram Stars is impossible
prepaid_upgrade_star_count	Integer	Optional. Number of Telegram Stars that were prepaid by the sender for the ability to upgrade the gift
can_be_upgraded	True	Optional. True, if the gift can be upgraded to a unique gift
text	String	Optional. Text of the message that was added to the gift
entities	Array of MessageEntity	Optional. Special entities that appear in the text
is_private	True	Optional. True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them
*/
type GiftInfo struct {
	// Information about the gift
	Gift Gift `json:"gift"`
	// Unique identifier of the received gift for the bot; only present for gifts received on behalf of business accounts
	OwnedGiftId string `json:"owned_gift_id,omitempty"`
	// Number of Telegram Stars that can be claimed by the receiver by converting the gift; omitted if conversion to Telegram Stars is impossible
	ConvertStarCount int `json:"convert_star_count,omitempty"`
	// Number of Telegram Stars that were prepaid by the sender for the ability to upgrade the gift
	PrepaidUpgradeStarCount int `json:"prepaid_upgrade_star_count,omitempty"`
	// True, if the gift can be upgraded to a unique gift
	CanBeUpgraded bool `json:"can_be_upgraded,omitempty"`
	// Text of the message that was added to the gift
	Text string `json:"text,omitempty"`
	// Special entities that appear in the text
	Entities []MessageEntity `json:"entities,omitempty"`
	// True, if the sender and gift text are shown only to the gift receiver; otherwise, everyone will be able to see them
	IsPrivate bool `json:"is_private,omitempty"`
}
