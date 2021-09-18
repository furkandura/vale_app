package repository

import (
	"vale_app/models"
	"vale_app/models/requests"

	"gorm.io/gorm"
)

type ParkingRepository struct {
	DB *gorm.DB
}

func (repo *Repositories) Parking() *ParkingRepository {
	return &ParkingRepository{DB: repo.DB}
}

func (pr *ParkingRepository) New(req requests.ParkingCreateRequest, companyId int) error {

	parking := models.Parking{
		CompanyId:     companyId,
		Type:          req.Type,
		CustomerId:    req.CustomerId,
		Plate:         "34 ABP 052",
		DateOfReceipt: req.DateOfReceipt,
		Note:          req.Note,
	}

	return pr.DB.Model(&models.Parking{}).Create(&parking).Error
}

func (pr *ParkingRepository) Update(req requests.ParkingUpdateRequest, companyId int) error {

	parking := models.Parking{
		CompanyId:     companyId,
		Type:          req.Type,
		CustomerId:    req.CustomerId,
		Plate:         req.Plate,
		DateOfReceipt: req.DateOfReceipt,
		Note:          req.Note,
	}

	return pr.DB.Model(&models.Parking{}).Where("id = ?", req.ParkingId).Updates(&parking).Error
}

func (pr *ParkingRepository) All(companyId int) []models.Parking {

	var parkings []models.Parking

	pr.DB.Model(&models.Parking{}).Where("company_id = ?", companyId).Find(&parkings)

	return parkings

}

// Girilen park kaydı oturum açmış kullancının mı kontrolü
func (pr *ParkingRepository) CheckParkingAuth(parkingId int, companyId int) bool {
	var count int64

	pr.DB.Model(&models.Parking{}).Where("company_id = ?", companyId).Where("id = ?", parkingId).Count(&count)

	if count == 0 {
		return false
	}

	return true
}

func (pr *ParkingRepository) Delete(parkingId int) error {
	var delete models.Parking
	return pr.DB.Model(&models.Parking{}).Where("id = ?", parkingId).Delete(&delete).Error
}
