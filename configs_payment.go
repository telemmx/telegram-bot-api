package tgbotapi

/**
sendInvoice
Use this method to send invoices. On success, the sent Message is returned.

Parameter	Type	Required	Description
chat_id	Integer or String	Yes	Unique identifier for the target chat or username of the target channel (in the format @channelusername)
message_thread_id	Integer	Optional	Unique identifier for the target message thread (topic) of the forum; for forum supergroups only
title	String	Yes	Product name, 1-32 characters
description	String	Yes	Product description, 1-255 characters
payload	String	Yes	Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
provider_token	String	Optional	Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
currency	String	Yes	Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
prices	Array of LabeledPrice	Yes	Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
max_tip_amount	Integer	Optional	The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
suggested_tip_amounts	Array of Integer	Optional	A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
start_parameter	String	Optional	Unique deep-linking parameter. If left empty, forwarded copies of the sent message will have a Pay button, allowing multiple users to pay directly from the forwarded message, using the same invoice. If non-empty, forwarded copies of the sent message will have a URL button with a deep link to the bot (instead of a Pay button), with the value used as the start parameter
provider_data	String	Optional	JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
photo_url	String	Optional	URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service. People like it better when they see what they are paying for.
photo_size	Integer	Optional	Photo size in bytes
photo_width	Integer	Optional	Photo width
photo_height	Integer	Optional	Photo height
need_name	Boolean	Optional	Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
need_phone_number	Boolean	Optional	Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
need_email	Boolean	Optional	Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
need_shipping_address	Boolean	Optional	Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
send_phone_number_to_provider	Boolean	Optional	Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
send_email_to_provider	Boolean	Optional	Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
is_flexible	Boolean	Optional	Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.
disable_notification	Boolean	Optional	Sends the message silently. Users will receive a notification with no sound.
protect_content	Boolean	Optional	Protects the contents of the sent message from forwarding and saving
allow_paid_broadcast	Boolean	Optional	Pass True to allow up to 1000 messages per second, ignoring broadcasting limits for a fee of 0.1 Telegram Stars per message. The relevant Stars will be withdrawn from the bot's balance
message_effect_id	String	Optional	Unique identifier of the message effect to be added to the message; for private chats only
reply_parameters	ReplyParameters	Optional	Description of the message to reply to
reply_markup	InlineKeyboardMarkup	Optional	A JSON-serialized object for an inline keyboard. If empty, one 'Pay total price' button will be shown. If not empty, the first button must be a Pay button.
*/

// InvoiceConfig contains information for sendInvoice request.
type InvoiceConfig struct {
	BaseChat
	Title                     string         // required
	Description               string         // required
	Payload                   string         // required
	ProviderToken             string         // required
	Currency                  string         // required
	Prices                    []LabeledPrice // required
	MaxTipAmount              int
	SuggestedTipAmounts       []int
	StartParameter            string
	ProviderData              string
	PhotoURL                  string
	PhotoSize                 int
	PhotoWidth                int
	PhotoHeight               int
	NeedName                  bool
	NeedPhoneNumber           bool
	NeedEmail                 bool
	NeedShippingAddress       bool
	SendPhoneNumberToProvider bool
	SendEmailToProvider       bool
	IsFlexible                bool
	AllowPaidBroadcast        bool
	MessageEffectID           int
	ReplyParameters           interface{}
}

func (config InvoiceConfig) params() (Params, error) {
	params, err := config.BaseChat.params()
	if err != nil {
		return params, err
	}

	params["title"] = config.Title
	params["description"] = config.Description
	params["payload"] = config.Payload
	params["provider_token"] = config.ProviderToken
	params["currency"] = config.Currency
	if err = params.AddInterface("prices", config.Prices); err != nil {
		return params, err
	}

	params.AddNonZero("max_tip_amount", config.MaxTipAmount)
	params.AddNonEmpty("start_parameter", config.StartParameter)
	params.AddNonEmpty("provider_data", config.ProviderData)
	params.AddNonEmpty("photo_url", config.PhotoURL)
	params.AddNonZero("photo_size", config.PhotoSize)
	params.AddNonZero("photo_width", config.PhotoWidth)
	params.AddNonZero("photo_height", config.PhotoHeight)
	params.AddBool("need_name", config.NeedName)
	params.AddBool("need_phone_number", config.NeedPhoneNumber)
	params.AddBool("need_email", config.NeedEmail)
	params.AddBool("need_shipping_address", config.NeedShippingAddress)
	params.AddBool("is_flexible", config.IsFlexible)
	params.AddBool("send_phone_number_to_provider", config.SendPhoneNumberToProvider)
	params.AddBool("send_email_to_provider", config.SendEmailToProvider)
	params.AddBool("allow_paid_broadcast", config.AllowPaidBroadcast)
	params.AddNonZero("message_effect_id", config.MessageEffectID)
	params.AddInterface("suggested_tip_amounts", config.SuggestedTipAmounts)
	params.AddInterface("reply_parameters", config.ReplyParameters)
	return params, nil
}

func (config InvoiceConfig) method() string {
	return "sendInvoice"
}

/*

createInvoiceLink
Use this method to create a link for an invoice. Returns the created invoice link as String on success.

Parameter	Type	Required	Description
business_connection_id	String	Optional	Unique identifier of the business connection on behalf of which the link will be created. For payments in Telegram Stars only.
title	String	Yes	Product name, 1-32 characters
description	String	Yes	Product description, 1-255 characters
payload	String	Yes	Bot-defined invoice payload, 1-128 bytes. This will not be displayed to the user, use it for your internal processes.
provider_token	String	Optional	Payment provider token, obtained via @BotFather. Pass an empty string for payments in Telegram Stars.
currency	String	Yes	Three-letter ISO 4217 currency code, see more on currencies. Pass “XTR” for payments in Telegram Stars.
prices	Array of LabeledPrice	Yes	Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc.). Must contain exactly one item for payments in Telegram Stars.
subscription_period	Integer	Optional	The number of seconds the subscription will be active for before the next payment. The currency must be set to “XTR” (Telegram Stars) if the parameter is used. Currently, it must always be 2592000 (30 days) if specified. Any number of subscriptions can be active for a given bot at the same time, including multiple concurrent subscriptions from the same user. Subscription price must no exceed 10000 Telegram Stars.
max_tip_amount	Integer	Optional	The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double). For example, for a maximum tip of US$ 1.45 pass max_tip_amount = 145. See the exp parameter in currencies.json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies). Defaults to 0. Not supported for payments in Telegram Stars.
suggested_tip_amounts	Array of Integer	Optional	A JSON-serialized array of suggested amounts of tips in the smallest units of the currency (integer, not float/double). At most 4 suggested tip amounts can be specified. The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
provider_data	String	Optional	JSON-serialized data about the invoice, which will be shared with the payment provider. A detailed description of required fields should be provided by the payment provider.
photo_url	String	Optional	URL of the product photo for the invoice. Can be a photo of the goods or a marketing image for a service.
photo_size	Integer	Optional	Photo size in bytes
photo_width	Integer	Optional	Photo width
photo_height	Integer	Optional	Photo height
need_name	Boolean	Optional	Pass True if you require the user's full name to complete the order. Ignored for payments in Telegram Stars.
need_phone_number	Boolean	Optional	Pass True if you require the user's phone number to complete the order. Ignored for payments in Telegram Stars.
need_email	Boolean	Optional	Pass True if you require the user's email address to complete the order. Ignored for payments in Telegram Stars.
need_shipping_address	Boolean	Optional	Pass True if you require the user's shipping address to complete the order. Ignored for payments in Telegram Stars.
send_phone_number_to_provider	Boolean	Optional	Pass True if the user's phone number should be sent to the provider. Ignored for payments in Telegram Stars.
send_email_to_provider	Boolean	Optional	Pass True if the user's email address should be sent to the provider. Ignored for payments in Telegram Stars.
is_flexible	Boolean	Optional	Pass True if the final price depends on the shipping method. Ignored for payments in Telegram Stars.

*/

type InvoiceLinkConfig struct {
	BusinessConnectionID      string
	Title                     string
	Description               string
	Payload                   string
	ProviderToken             string
	Currency                  string
	Prices                    []LabeledPrice
	SubscriptionPeriod        int
	MaxTipAmount              int
	SuggestedTipAmounts       []int
	ProviderData              string
	PhotoURL                  string
	PhotoSize                 int
	PhotoWidth                int
	PhotoHeight               int
	NeedName                  bool
	NeedPhoneNumber           bool
	NeedEmail                 bool
	NeedShippingAddress       bool
	SendPhoneNumberToProvider bool
	SendEmailToProvider       bool
	IsFlexible                bool
}

func (config InvoiceLinkConfig) params() (Params, error) {
	params := Params{}
	params.AddNonEmpty("business_connection_id", config.BusinessConnectionID)
	params.AddNonEmpty("title", config.Title)
	params.AddNonEmpty("description", config.Description)
	params.AddNonEmpty("payload", config.Payload)
	params.AddNonEmpty("provider_token", config.ProviderToken)
	params.AddNonEmpty("currency", config.Currency)
	params.AddNonZero("subscription_period", config.SubscriptionPeriod)
	params.AddNonZero("max_tip_amount", config.MaxTipAmount)
	params.AddInterface("suggested_tip_amounts", config.SuggestedTipAmounts)
	params.AddNonEmpty("provider_data", config.ProviderData)
	params.AddNonEmpty("photo_url", config.PhotoURL)
	params.AddNonZero("photo_size", config.PhotoSize)
	params.AddNonZero("photo_width", config.PhotoWidth)
	params.AddNonZero("photo_height", config.PhotoHeight)
	params.AddBool("need_name", config.NeedName)
	params.AddBool("need_phone_number", config.NeedPhoneNumber)
	params.AddBool("need_email", config.NeedEmail)
	params.AddBool("need_shipping_address", config.NeedShippingAddress)
	params.AddBool("send_phone_number_to_provider", config.SendPhoneNumberToProvider)
	params.AddBool("send_email_to_provider", config.SendEmailToProvider)
	params.AddBool("is_flexible", config.IsFlexible)
	err := params.AddInterface("prices", config.Prices)
	return params, err
}

func (config InvoiceLinkConfig) method() string {
	return "createInvoiceLink"
}

/*

getStarTransactions
Returns the bot's Telegram Star transactions in chronological order. On success, returns a StarTransactions object.

Parameter	Type	Required	Description
offset	Integer	Optional	Number of transactions to skip in the response
limit	Integer	Optional	The maximum number of transactions to be retrieved. Values between 1-100 are accepted. Defaults to 100.
*/

type GetStarTransactionsConfig struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

func (config GetStarTransactionsConfig) method() string {
	return "getStarTransactions"
}

func (config GetStarTransactionsConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero("offset", config.Offset)
	params.AddNonZero("limit", config.Limit)
	return params, nil
}

/*
refundStarPayment
Refunds a successful payment in Telegram Stars. Returns True on success.

Parameter	Type	Required	Description
user_id	Integer	Yes	Identifier of the user whose payment will be refunded
telegram_payment_charge_id	String	Yes	Telegram payment identifier

*/

type RefundStarPaymentConfig struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
}

func (config RefundStarPaymentConfig) method() string {
	return "refundStarPayment"
}

func (config RefundStarPaymentConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("telegram_payment_charge_id", config.TelegramPaymentChargeID)
	return params, nil
}

/*
editUserStarSubscription
Allows the bot to cancel or re-enable extension of a subscription paid in Telegram Stars. Returns True on success.

Parameter	Type	Required	Description
user_id	Integer	Yes	Identifier of the user whose subscription will be edited
telegram_payment_charge_id	String	Yes	Telegram payment identifier for the subscription
is_canceled	Boolean	Yes	Pass True to cancel extension of the user subscription; the subscription must be active up to the end of the current subscription period. Pass False to allow the user to re-enable a subscription that was previously canceled by the bot.

*/

type EditUserStarSubscriptionConfig struct {
	UserID                  int64  `json:"user_id"`
	TelegramPaymentChargeID string `json:"telegram_payment_charge_id"`
	IsCanceled              bool   `json:"is_canceled"`
}

func (config EditUserStarSubscriptionConfig) method() string {
	return "editUserStarSubscription"
}

func (config EditUserStarSubscriptionConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddNonEmpty("telegram_payment_charge_id", config.TelegramPaymentChargeID)
	params.AddBool("is_canceled", config.IsCanceled)
	return params, nil
}

type GetMyStarBalanceConfig struct{}

func (config GetMyStarBalanceConfig) method() string {
	return "getMyStarBalance"
}

func (config GetMyStarBalanceConfig) params() (Params, error) {
	return Params{}, nil
}
