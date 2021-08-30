package parking

import (
	"net/http"
	"strconv"
	"vale_app/internal/helpers"
	"vale_app/internal/repository"
	"vale_app/models/requests"

	"github.com/labstack/echo"
)

// Park kaydı oluşturur.
func Create(c echo.Context) error {
	var req requests.ParkingCreateRequest

	_ = c.Bind(&req)

	err := c.Validate(&req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(err.Error(), nil, http.StatusNotFound))
	}

	isCustomerAuth := repository.Get().Customer().CheckCustomerAuth(req.CustomerId, helpers.GetAuthID(c))

	if !isCustomerAuth {
		return c.JSON(http.StatusNotFound, helpers.Response("Müşteri kaydı bulunamadı.", nil, http.StatusNotFound))
	}

	err = repository.Get().Parking().New(req, helpers.GetAuthID(c))

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt oluşturulurken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla oluşturuldu.", nil, http.StatusOK))

}

// Park kaydını günceller.
func Update(c echo.Context) error {
	var req requests.ParkingUpdateRequest

	_ = c.Bind(&req)

	err := c.Validate(&req)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response(err.Error(), nil, http.StatusNotFound))
	}

	checkParkingAuth := repository.Get().Parking().CheckParkingAuth(req.ParkingId, helpers.GetAuthID(c))

	if !checkParkingAuth {
		return c.JSON(http.StatusNotFound, helpers.Response("Park kaydı bulunamadı.", nil, http.StatusNotFound))
	}

	if req.Type == 2 && req.DateOfDelivery == nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Park kaydını 'Müşteriye teslim edildi.' olarak değiştirmek için lütfen müşterinin teslim aldığı tarihi girin.", nil, http.StatusNotFound))
	}

	err = repository.Get().Parking().Update(req, helpers.GetAuthID(c))

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt güncellenirken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla güncellendi.", nil, http.StatusOK))

}

// Park kaydını siler.
func Delete(c echo.Context) error {

	id := c.Param("id")
	idInt, err := strconv.Atoi(id)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Geçerli bir id girilmedi.", nil, http.StatusNotFound))
	}

	checkParkingAuth := repository.Get().Parking().CheckParkingAuth(idInt, helpers.GetAuthID(c))

	if !checkParkingAuth {
		return c.JSON(http.StatusNotFound, helpers.Response("Park kaydı bulunamadı.", nil, http.StatusNotFound))
	}

	err = repository.Get().Parking().Delete(idInt)

	if err != nil {
		return c.JSON(http.StatusNotFound, helpers.Response("Kayıt silinirken bir hata oluştu.", nil, http.StatusNotFound))
	}

	return c.JSON(http.StatusOK, helpers.Response("Kayıt başarıyla silindi.", nil, http.StatusOK))

}
