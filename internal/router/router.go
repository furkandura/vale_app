package router

import (
	"vale_app/internal/api/company"

	"github.com/labstack/echo"
)

func Set(e *echo.Echo) {

	api := e.Group("/api")

	companyRoutes := api.Group("/company")
	companyRoutes.POST("/register", company.Register)

}
