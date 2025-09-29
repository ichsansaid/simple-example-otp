package dto

type ReqOtpDto struct {
	UserId string `validate:"required" json:"user_id"`
}
