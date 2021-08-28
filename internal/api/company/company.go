package company

import (
	"net/http"
	"vale_app/internal/helpers"
	"vale_app/internal/repository"
	"vale_app/models/requests"

	"github.com/labstack/echo"
)

func Login(c echo.Context) error {
	return c.JSON(http.StatusNotFound, "Sea")
}

func Register(c echo.Context) error {
	var req requests.CompanyRegisterRequest

	_ = c.Bind(&req)

	err := c.Validate(&req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(err.Error(), nil, http.StatusNotFound))
	}

	isValidPhone := helpers.IsValidPhone(req.Phone)

	if isValidPhone == false {
		return c.JSON(http.StatusNotFound, helpers.Response("Girilen telefon numarası geçersiz.", nil, http.StatusNotFound))
	}

	isRepeatPhone := repository.Get().Company().RepeatedPhoneCheck(req.Phone)

	if isRepeatPhone == false {
		return c.JSON(http.StatusNotFound, helpers.Response("Girilen telefon numarası ile daha önceden kayıt yapılmış.", nil, http.StatusNotFound))
	}

	err = repository.Get().Company().New(req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt oluşturulurken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla oluşturuldu.", nil, http.StatusOK))

}