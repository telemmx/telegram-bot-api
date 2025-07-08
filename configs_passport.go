package tgbotapi

/*
setPassportDataErrors
Informs a user that some of the Telegram Passport elements they provided contains errors. The user will not be able to re-submit their Passport to you until the errors are fixed (the contents of the field for which you returned the error must change). Returns True on success.

Use this if the data submitted by the user doesn't satisfy the standards your service requires for any reason. For example, if a birthday date seems invalid, a submitted document is blurry, a scan shows evidence of tampering, etc. Supply some details in the error message to make sure the user knows how to correct the issues.

Parameter	Type	Required	Description
user_id	Integer	Yes	User identifier
errors	Array of PassportElementError	Yes	A JSON-serialized array describing the errors
*/

type SetPassportDataErrorsConfig struct {
	UserID int64                  `json:"user_id"`
	Errors []PassportElementError `json:"errors"`
}

func (config SetPassportDataErrorsConfig) method() string {
	return "setPassportDataErrors"
}

func (config SetPassportDataErrorsConfig) params() (Params, error) {
	params := Params{}
	params.AddNonZero64("user_id", config.UserID)
	params.AddInterface("errors", config.Errors)
	return params, nil
}
