package router

import (
	"vale_app/internal/api/company"
	"vale_app/internal/api/customer"
	"vale_app/internal/api/parking"
	"vale_app/internal/middlewares"

	"github.com/labstack/echo"
)

func Set(e *echo.Echo) {

	api := e.Group("/api")

	companyRoutes := api.Group("/company")
	companyRoutes.POST("/register", company.Register)
	companyRoutes.POST("/login", company.Login)
	companyRoutes.POST("/update", company.Update, middlewares.VerifyToken)

	parkingRoutes := api.Group("/parking", middlewares.VerifyToken)
	parkingRoutes.POST("/create", parking.Create)
	parkingRoutes.POST("/update", parking.Update)
	parkingRoutes.GET("/delete/:id", parking.Delete)

	customerRoutes := api.Group("/customer", middlewares.VerifyToken)
	customerRoutes.POST("/create", customer.Create)
	customerRoutes.POST("/update", customer.Update)
	customerRoutes.GET("/delete/:id", customer.Delete)

}
