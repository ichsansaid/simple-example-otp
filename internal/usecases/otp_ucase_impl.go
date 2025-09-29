package ucases

import (
	"context"
	"fmt"
	"time"

	"github.com/ichsansaid/multi-artha-otp/internal/entity"
	repo "github.com/ichsansaid/multi-artha-otp/internal/repositories"
)

type OtpUcaseImpl struct {
	repo repo.OtpRepoInterface
}

func NewOtpUcase(repo repo.OtpRepoInterface) OtpUcaseInterface {
	return &OtpUcaseImpl{
		repo: repo,
	}
}

// RequestOtp implements OtpUcaseInterface.
func (o *OtpUcaseImpl) RequestOtp(ctx context.Context, userId string) (entity.OtpEntity, error) {
	var result entity.OtpEntity
	if err := o.repo.Do(ctx, func(txCtx context.Context) error {
		if otp, err := o.repo.FindOtp(txCtx, userId); err != nil {
			return err
		} else {
			if otp.UserId == "" || otp.IsExpired(time.Now()) {

				otp.UserId = userId
				otp.CreateOtpCode()
				if err := o.repo.SaveOtp(txCtx, otp); err != nil {
					return err
				}

			}
			result = otp
			return nil
		}
	}); err != nil {
		return result, err
	}
	return result, nil

}

// ValidateOtp implements OtpUcaseInterface.
func (o *OtpUcaseImpl) ValidateOtp(ctx context.Context, userId string, otpCode string) (bool, error) {
	var result bool = false
	if err := o.repo.Do(ctx, func(ctx context.Context) error {
		if otp, err := o.repo.FindOtp(ctx, userId); err != nil {
			return err
		} else if otp.OtpCode != otpCode {
			return nil
		} else {
			if otp.IsExpired(time.Now()) {
				return fmt.Errorf("Expired")
			}
			otp.Status = "validated"
			if err := o.repo.SaveOtp(ctx, otp); err != nil {
				return err
			}
		}
		result = true
		return nil
	}); err != nil {
		return result, err
	}

	return result, nil
}
