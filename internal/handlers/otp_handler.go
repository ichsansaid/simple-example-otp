package handler

import (
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/ichsansaid/multi-artha-otp/internal/dto"
	ucases "github.com/ichsansaid/multi-artha-otp/internal/usecases"
)

type OtpHandler struct {
	ucase     ucases.OtpUcaseInterface
	validator *validator.Validate
}

func NewOtpHandler(ucase ucases.OtpUcaseInterface, validator *validator.Validate) *OtpHandler {
	return &OtpHandler{
		ucase:     ucase,
		validator: validator,
	}
}

func (h *OtpHandler) RequestOtp(c *fiber.Ctx) error {
	reqDto := dto.ReqOtpDto{}
	if err := c.BodyParser(&reqDto); err != nil {
		return c.JSON(fiber.Map{
			"error":             "invalid request",
			"error_description": err.Error(),
		})
	}
	if err := h.validator.Struct(reqDto); err != nil {
		return c.JSON(fiber.Map{
			"error":             "invalid request",
			"error_description": err.Error(),
		})
	}
	if otp, err := h.ucase.RequestOtp(c.Context(), reqDto.UserId); err != nil {
		return c.JSON(fiber.Map{
			"error":             "internal server error",
			"error_description": err.Error(),
		})
	} else {
		return c.JSON(fiber.Map{
			"user_id": reqDto.UserId,
			"otp":     otp.OtpCode,
		})
	}
}

func (h *OtpHandler) ValidateOtp(c *fiber.Ctx) error {
	reqDto := dto.ValidateOtpDto{}
	if err := c.BodyParser(&reqDto); err != nil {
		return c.JSON(fiber.Map{
			"error":             "invalid_request",
			"error_description": err.Error(),
		})
	}
	if err := h.validator.Struct(reqDto); err != nil {
		return c.JSON(fiber.Map{
			"error":             "invalid_request",
			"error_description": err.Error(),
		})
	}
	if isFound, err := h.ucase.ValidateOtp(c.Context(), reqDto.UserId, reqDto.Otp); err != nil {
		return c.JSON(fiber.Map{
			"error":             "internal_server_error",
			"error_description": err.Error(),
		})
	} else {
		if !isFound {
			return c.JSON(fiber.Map{
				"error":             "otp_not_found",
				"error_description": "OTP Not Found",
			})
		}
		return c.JSON(fiber.Map{
			"user_id": reqDto.UserId,
			"message": "OTP validated successfully",
		})
	}
}
