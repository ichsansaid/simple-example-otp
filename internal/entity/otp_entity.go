package entity

import (
	"math/rand"
	"strconv"
	"time"
)

// Asumsikan satu User hanya bisa satu OTP secara kardinalitas
type OtpEntity struct {
	UserId    string `json:"user_id" gorm:"primaryKey"`
	OtpCode   string `json:"otp_code"`
	Status    string `json:"status"`
	ExpiredAt int64  `json:"expired_at"`
}

func (e *OtpEntity) CreateOtpCode() string {
	min := 100000
	max := 999999
	e.Status = "pending"
	e.OtpCode = strconv.Itoa(rand.Intn(max-min+1) + min)
	e.ExpiredAt = time.Now().Add(time.Minute * 20).Unix()
	return e.OtpCode
}

func (e *OtpEntity) IsExpired(t time.Time) bool {
	return e.ExpiredAt < t.Unix()
}

func (e *OtpEntity) TableName() string {
	return "otp"
}
