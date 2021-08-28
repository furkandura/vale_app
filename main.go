package main

import (
	"vale_app/internal/helpers"
	"vale_app/internal/router"

	"github.com/go-playground/validator"
	"github.com/labstack/echo"
)

func main() {

	// Auto migrate

	e := echo.New()
	e.Validator = &helpers.CustomValidator{Validator: validator.New()}

	router.Set(e)

	e.Logger.Fatal(e.Start(":8889"))

}
