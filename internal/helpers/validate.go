package helpers

import (
	"errors"
	"fmt"

	"github.com/go-playground/validator"
)

type CustomValidator struct {
	Validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {

	var message string

	// cv.Validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
	// 	name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
	// 	if name == "-" {
	// 		return ""
	// 	}
	// 	return name
	// })

	err := cv.Validator.Struct(i)

	if err != nil {
		for inx, item := range err.(validator.ValidationErrors) {
			// fmt.Println(item.Tag(), item.Param(), item.Kind(), item.StructNamespace(), item.Field(), item.Value())
			if inx != 0 {
				message += " "
			}

			if item.Tag() == "required" {
				message += fmt.Sprintf("%s alanı gereklidir.", item.Field())
			} else if item.Tag() == "email" {
				message += fmt.Sprintf("%s alanına girilen e-posta adresi geçersiz.", item.Field())
			} else if item.Tag() == "min" {
				message += fmt.Sprintf("%s alanı %s değerinden büyük olmalıdır.", item.Field(), item.Param())
			} else if item.Tag() == "max" {
				message += fmt.Sprintf("%s alanı %s değerinden küçük olmalıdır.", item.Field(), item.Param())
			}
		}

		return errors.New(message)
	}

	return nil

}
