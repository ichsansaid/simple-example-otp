package main

import (
	"github.com/go-playground/validator/v10"
	"github.com/ichsansaid/multi-artha-otp/internal/app"
	handler "github.com/ichsansaid/multi-artha-otp/internal/handlers"
	repo "github.com/ichsansaid/multi-artha-otp/internal/repositories"
	ucases "github.com/ichsansaid/multi-artha-otp/internal/usecases"
)

func main() {

	//Asumsikan menggunakan DI sederhana
	gormDb := app.NewGormDB()
	otpRepo := repo.NewOtpRepoImpl(gormDb)
	otpUcase := ucases.NewOtpUcase(otpRepo)
	instanceValidator := validator.New()
	otpHandler := handler.NewOtpHandler(otpUcase, instanceValidator)

	fiberApp := app.NewFiberApp()

	app.InvokeRouter(fiberApp, *otpHandler)

	//Asumsikan port nya statis aja gaperlu dikeluarkan ke ENV
	if err := fiberApp.Listen(":3000"); err != nil {
		panic(err)
	}
}
