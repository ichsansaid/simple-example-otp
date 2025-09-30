package repo

import (
	"context"

	"github.com/ichsansaid/multi-artha-otp/internal/entity"
	"gorm.io/gorm"
)

type TxKey struct{}

type OtpRepoImpl struct {
	DB *gorm.DB
}

func NewOtpRepoImpl(db *gorm.DB) OtpRepoInterface {
	return &OtpRepoImpl{DB: db}
}

// Do implements OtpRepoInterface.
func (o *OtpRepoImpl) Do(ctx context.Context, fn func(context.Context) error) error {
	tx := o.DB.Begin()

	defer tx.Rollback()
	if err := fn(context.WithValue(ctx, TxKey{}, tx)); err != nil {
		return err
	} else {
		return tx.Commit().Error
	}
}

// FindOtp implements OtpRepoInterface.
func (o *OtpRepoImpl) FindOtp(ctx context.Context, userId string) (entity.OtpEntity, error) {
	dbSelected := o.DB
	if tx, ok := ctx.Value(TxKey{}).(*gorm.DB); ok {
		dbSelected = tx
	}
	var otp entity.OtpEntity
	if err := dbSelected.Where("user_id = ?", userId).Find(&otp).Error; err != nil {
		return entity.OtpEntity{}, err
	}
	return otp, nil

}

// SaveOtp implements OtpRepoInterface.
func (o *OtpRepoImpl) SaveOtp(ctx context.Context, otp entity.OtpEntity) error {
	dbSelected := o.DB
	if tx, ok := ctx.Value(TxKey{}).(*gorm.DB); ok {
		dbSelected = tx
	}
	if err := dbSelected.Save(&otp).Error; err != nil {
		return err
	}
	return nil

}
