package company

import (
	"net/http"
	"vale_app/internal/helpers"
	"vale_app/internal/repository"
	"vale_app/models/requests"

	"github.com/labstack/echo"
)

// Vale firmasi giriş.
func Login(c echo.Context) error {
	var req requests.CompanyLoginRequest

	_ = c.Bind(&req)

	err := c.Validate(&req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(err.Error(), nil, http.StatusNotFound))
	}

	company := repository.Get().Company().Login(req.Phone, req.Password)

	if company.ID == 0 {
		return c.JSON(http.StatusNotFound, helpers.Response("Girdiğiniz telefon numarası veya şifre yanlış.", nil, http.StatusNotFound))
	}

	token, err := helpers.GenerateToken(company.ID)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Oturum anahtarı oluşturulurken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Giriş başarılı.", map[string]interface{}{
		"token": token,
		"user":  company,
	}, http.StatusOK))

}

// Vale firmasi kayıt ol.
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

// Vale firmasi bilgileri güncele.
func Update(c echo.Context) error {
	var req requests.CompanyUpdateRequest

	_ = c.Bind(&req)

	isValidPhone := helpers.IsValidPhone(req.Phone)

	if isValidPhone == false {
		return c.JSON(http.StatusNotFound, helpers.Response("Girilen telefon numarası geçersiz.", nil, http.StatusNotFound))
	}

	isRepeatPhone := repository.Get().Company().RepeatedPhoneCheck(req.Phone)

	if isRepeatPhone == false {
		return c.JSON(http.StatusNotFound, helpers.Response("Girilen telefon numarasını başka bir hesap zaten kullanılıyor.", nil, http.StatusNotFound))
	}

	err := repository.Get().Company().Update(req, helpers.GetAuthID(c))

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt güncellenirken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla güncellendi.", nil, http.StatusOK))

}
