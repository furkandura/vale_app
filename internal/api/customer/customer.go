package customer

import (
	"net/http"
	"strconv"
	"vale_app/internal/helpers"
	"vale_app/internal/repository"
	"vale_app/models/requests"

	"github.com/labstack/echo"
)

func Create(c echo.Context) error {
	var req requests.CustomerCreateRequest

	_ = c.Bind(&req)

	err := c.Validate(&req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(err.Error(), nil, http.StatusNotFound))
	}

	if req.Phone != nil && !helpers.IsValidPhone(*req.Phone) {
		return c.JSON(http.StatusNotFound, helpers.Response("Girdiğiniz telefon numarası geçersiz.", nil, http.StatusNotFound))
	}

	err = repository.Get().Customer().New(req, helpers.GetAuthID(c))

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt oluşturulurken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla oluşturuldu.", nil, http.StatusOK))
}

func Update(c echo.Context) error {
	var req requests.CustomerUpdateRequest

	_ = c.Bind(&req)

	err := c.Validate(&req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(err.Error(), nil, http.StatusNotFound))
	}

	if req.Phone != nil && !helpers.IsValidPhone(*req.Phone) {
		return c.JSON(http.StatusNotFound, helpers.Response("Girdiğiniz telefon numarası geçersiz.", nil, http.StatusNotFound))
	}

	isCustomerAuth := repository.Get().Customer().CheckCustomerAuth(req.CustomerId, helpers.GetAuthID(c))

	if !isCustomerAuth {
		return c.JSON(http.StatusNotFound, helpers.Response("Müşteri kaydı bulunamadı.", nil, http.StatusNotFound))
	}

	err = repository.Get().Customer().Update(req, helpers.GetAuthID(c))

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt güncellenirken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla güncellendi.", nil, http.StatusOK))
}

func Delete(c echo.Context) error {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Geçerli bir id girilmedi.", nil, http.StatusNotFound))
	}

	checkCustomerAuth := repository.Get().Customer().CheckCustomerAuth(idInt, helpers.GetAuthID(c))

	if !checkCustomerAuth {
		return c.JSON(http.StatusNotFound, helpers.Response("Park kaydı bulunamadı.", nil, http.StatusNotFound))
	}

	err = repository.Get().Customer().Delete(idInt)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt silinirken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla silindi.", nil, http.StatusOK))

}
