package main

import (
	"github.com/go-playground/validator"
	"github.com/renaldid/go-sisko/app"
	"github.com/renaldid/go-sisko/controller"
	"github.com/renaldid/go-sisko/helper"
	"github.com/renaldid/go-sisko/middleware"
	"github.com/renaldid/go-sisko/repository"
	"github.com/renaldid/go-sisko/service"
	"net/http"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	siswaRepo := repository.NewSiswaRepo()
	siswaService := service.NewSiswaService(siswaRepo, db, validate)
	siswaController := controller.NewSiswaController(siswaService)
	router := app.NewRouter(siswaController)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
