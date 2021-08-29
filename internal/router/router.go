package router

import (
	"fmt"
	"vale_app/internal/api/company"
	"vale_app/internal/helpers"
	"vale_app/internal/middlewares"

	"github.com/labstack/echo"
)

func Set(e *echo.Echo) {

	api := e.Group("/api")

	companyRoutes := api.Group("/company")
	companyRoutes.POST("/register", company.Register)
	companyRoutes.POST("/login", company.Login)
	companyRoutes.POST("/test", func(c echo.Context) error {

		fmt.Println(helpers.GetAuthID(c))
		return nil

	}, middlewares.VerifyToken)

}
