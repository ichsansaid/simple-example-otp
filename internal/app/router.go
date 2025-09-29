package app

import (
	"github.com/gofiber/fiber/v2"
	handler "github.com/ichsansaid/multi-artha-otp/internal/handlers"
)

func InvokeRouter(app *fiber.App, h handler.OtpHandler) {
	app.Post("/request-otp", h.RequestOtp)
	app.Post("/validate-otp", h.ValidateOtp)
}
