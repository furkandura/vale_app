package repository

import (
	"vale_app/models"
	"vale_app/models/requests"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func (repo *Repositories) Company() *CompanyRepository {
	return &CompanyRepository{DB: repo.DB}
}

func (cr *CompanyRepository) New(req requests.CompanyRegisterRequest) error {

	password, _ := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)

	company := models.Company{
		FullName:    req.FullName,
		Phone:       req.Phone,
		CompanyName: req.CompanyName,
		Password:    string(password),
		Email:       req.Email,
	}

	return cr.DB.Model(models.Company{}).Create(&company).Error
}

// Kayıtlı telefon var mı diye kontröl eder.
func (cr *CompanyRepository) RepeatedPhoneCheck(phone string) bool {

	var count int64

	cr.DB.Model(models.Company{}).Where("phone = ?", phone).Count(&count)

	if count > 0 {
		return false
	}

	return true
}

func (cr *CompanyRepository) Login(phone string, password string) models.Company {
	var user models.Company

	cr.DB.Model(&models.Company{}).Where("phone = ?", phone).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return user
	}

	return user
}
