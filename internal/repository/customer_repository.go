package repository

import (
	"vale_app/models"
	"vale_app/models/requests"

	"gorm.io/gorm"
)

type CustomerRepository struct {
	DB *gorm.DB
}

func (repo *Repositories) Customer() *CustomerRepository {
	return &CustomerRepository{DB: repo.DB}
}

func (cr *CustomerRepository) New(req requests.CustomerCreateRequest, companyId int) error {
	customer := &models.Customer{
		FullName:  req.FullName,
		CompanyId: companyId,
		Phone:     req.Phone,
		Vehicles:  req.Vehicles,
		Note:      req.Phone,
	}
	return cr.DB.Model(&models.Customer{}).Create(&customer).Error
}

func (cr *CustomerRepository) Update(req requests.CustomerUpdateRequest, companyId int) error {
	customer := &models.Customer{
		FullName: req.FullName,
		Phone:    req.Phone,
		Vehicles: req.Vehicles,
		Note:     req.Phone,
	}
	return cr.DB.Model(&models.Customer{}).Where("id = ?", req.CustomerId).Updates(&customer).Error
}

// Girilen müşteri kaydı oturum açmış kullancının mı kontrolü
func (pr *CustomerRepository) CheckCustomerAuth(customerId int, companyId int) bool {
	var count int64

	pr.DB.Model(&models.Customer{}).Where("company_id = ?", companyId).Where("id = ?", customerId).Count(&count)

	if count == 0 {
		return false
	}

	return true
}

func (pr *CustomerRepository) Delete(customerId int) error {
	var delete models.Parking
	return pr.DB.Model(&models.Customer{}).Where("id = ?", customerId).Delete(&delete).Error
}

func (pr *CustomerRepository) FindById(customerId int) models.Customer {
	var customer models.Customer
	pr.DB.Model(&models.Customer{}).Where("id = ?", customerId).First(&customer)

	return customer
}
