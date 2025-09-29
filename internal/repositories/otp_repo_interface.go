package repo

import (
	"context"

	"github.com/ichsansaid/multi-artha-otp/internal/entity"
)

type OtpRepoInterface interface {
	SaveOtp(ctx context.Context, otp entity.OtpEntity) error
	FindOtp(ctx context.Context, userId string) (entity.OtpEntity, error)
	DeleteOtp(ctx context.Context, userId string) error
	Do(ctx context.Context, fn func(context.Context) error) error
}
