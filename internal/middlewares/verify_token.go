package middlewares

import (
	"net/http"
	"vale_app/internal/helpers"

	"github.com/labstack/echo"
)

func VerifyToken(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		_, err := helpers.ParseToken(helpers.GetHeaderToken(c))

		if err != nil {
			return c.JSON(http.StatusUnauthorized, helpers.Response(err.Error(), nil, http.StatusUnauthorized))
		}

		return next(c)

	}
}
