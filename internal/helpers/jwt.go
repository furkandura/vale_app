package helpers

import (
	"errors"
	"strings"
	"time"
	"vale_app/configs"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo"
)

type JwtCustomClaims struct {
	ID int
	jwt.StandardClaims
}

// Oturum açmış kullanıcının idsini döner.
func GetAuthID(c echo.Context) int {
	claims, _ := ParseToken(GetHeaderToken(c))
	return claims.ID
}

// Gelen token parçalanıp eğer ki hata varsa hatayı yoksa claims'i döner.
func ParseToken(tokenStr string) (*JwtCustomClaims, error) {

	tkn, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(configs.SecretKeyJwt), nil
	})

	if err != nil {
		return nil, errors.New("Oturum anahtarı geçersiz veya süresi dolmuş.")
	}

	if !tkn.Valid {
		return nil, errors.New("Oturum anahtarı doğrulanamadı.")
	}

	claims := tkn.Claims.(*JwtCustomClaims)

	return claims, nil
}

// Gelen request header'ından tokenı alır.
func GetHeaderToken(c echo.Context) string {
	headerToken := strings.Replace(c.Request().Header.Get("Authorization"), "Bearer ", "", 1)
	return headerToken
}

func GenerateToken(companyId int) (string, error) {

	claims := &JwtCustomClaims{
		ID: companyId,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 12).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(configs.SecretKeyJwt))

	return tokenString, err
}
