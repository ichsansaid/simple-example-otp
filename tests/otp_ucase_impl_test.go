package tests

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/ichsansaid/multi-artha-otp/internal/entity"
	ucases "github.com/ichsansaid/multi-artha-otp/internal/usecases"
)

type MockOtpRepo struct {
	mock.Mock
}

func (m *MockOtpRepo) Do(ctx context.Context, fn func(txCtx context.Context) error) error {
	return fn(ctx)
}

func (m *MockOtpRepo) FindOtp(ctx context.Context, userId string) (entity.OtpEntity, error) {
	args := m.Called(ctx, userId)
	return args.Get(0).(entity.OtpEntity), args.Error(1)
}

func (m *MockOtpRepo) SaveOtp(ctx context.Context, otp entity.OtpEntity) error {
	args := m.Called(ctx, otp)
	return args.Error(0)
}

const testUserID = "user-test-01"
const correctOtpCode = "123456"

func expiredTimeUnix() int64 {
	return time.Now().Add(-5 * time.Minute).Unix()
}

func validTimeUnix() int64 {
	return time.Now().Add(5 * time.Minute).Unix()
}

func TestOtpUcaseImpl_RequestOtp(t *testing.T) {
	mockRepo := new(MockOtpRepo)
	ucase := ucases.NewOtpUcase(mockRepo)
	ctx := context.Background()

	t.Run("Create_New_Otp_When_NotFound", func(t *testing.T) {
		mockRepo.
			On("FindOtp", ctx, testUserID).
			Return(entity.OtpEntity{}, nil).
			Once()

		mockRepo.
			On("SaveOtp", ctx, mock.AnythingOfType("entity.OtpEntity")).
			Return(nil).
			Run(func(args mock.Arguments) {
				savedOtp := args.Get(1).(entity.OtpEntity)
				assert.Equal(t, testUserID, savedOtp.UserId, "UserId harus diisi")
				assert.Len(t, savedOtp.OtpCode, 6, "OtpCode harus 6 digit")
				assert.Equal(t, "pending", savedOtp.Status, "Status harus 'pending' setelah dibuat")
				assert.Greater(t, savedOtp.ExpiredAt, time.Now().Unix(), "ExpiredAt harus di masa depan")
			}).
			Once()

		result, err := ucase.RequestOtp(ctx, testUserID)

		assert.NoError(t, err)
		assert.Equal(t, testUserID, result.UserId)
		assert.Equal(t, "pending", result.Status)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Recreate_When_Otp_Expired", func(t *testing.T) {
		expiredOtp := entity.OtpEntity{
			UserId:    testUserID,
			OtpCode:   "000000",
			Status:    "pending",
			ExpiredAt: expiredTimeUnix(),
		}
		mockRepo.
			On("FindOtp", ctx, testUserID).
			Return(expiredOtp, nil).
			Once()

		mockRepo.
			On("SaveOtp", ctx, mock.AnythingOfType("entity.OtpEntity")).
			Return(nil).
			Run(func(args mock.Arguments) {
				savedOtp := args.Get(1).(entity.OtpEntity)
				assert.NotEqual(t, expiredOtp.OtpCode, savedOtp.OtpCode)
				assert.Equal(t, "pending", savedOtp.Status)
			}).
			Once()

		result, err := ucase.RequestOtp(ctx, testUserID)

		assert.NoError(t, err)
		assert.Equal(t, testUserID, result.UserId)
		assert.Equal(t, "pending", result.Status)
		assert.NotEqual(t, expiredOtp.OtpCode, result.OtpCode)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Return_Existing_Otp_When_Valid", func(t *testing.T) {
		validOtp := entity.OtpEntity{
			UserId:    testUserID,
			OtpCode:   correctOtpCode,
			Status:    "pending",
			ExpiredAt: validTimeUnix(),
		}
		mockRepo.
			On("FindOtp", ctx, testUserID).
			Return(validOtp, nil).
			Once()

		mockRepo.
			On("SaveOtp", ctx, mock.Anything).
			Return(nil).
			Maybe()

		result, err := ucase.RequestOtp(ctx, testUserID)

		assert.NoError(t, err)
		assert.Equal(t, validOtp, result)
	})
}
