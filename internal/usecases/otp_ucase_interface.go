package ucases

import (
	"context"

	"github.com/ichsansaid/multi-artha-otp/internal/entity"
)

type OtpUcaseInterface interface {
	RequestOtp(ctx context.Context, userId string) (entity.OtpEntity, error)
	ValidateOtp(ctx context.Context, userId string, otp string) (bool, error)
}
