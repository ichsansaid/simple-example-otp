package repo

import (
	"context"

	"github.com/ichsansaid/multi-artha-otp/internal/entity"
)

type OtpRepoInterface interface {
	SaveOtp(ctx context.Context, otp entity.OtpEntity) error
	FindOtp(ctx context.Context, userId string) (entity.OtpEntity, error)
	Do(ctx context.Context, fn func(context.Context) error) error
}
