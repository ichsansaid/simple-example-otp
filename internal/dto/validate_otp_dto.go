package dto

type ValidateOtpDto struct {
	UserId string `validate:"required" json:"user_id"`
	Otp    string `validate:"required,min=6,max=6" json:"otp"`
}
