package JWTUtils

import (
	"Lahu/symptomsTracker/api/model"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateTocken(login model.UserDetails) (string, error) {
	claims := model.JwtCustomClaims{
		login.LoginId,

		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
		},
	}
	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte(model.JwtKey))
	if err != nil {
		// logginghelper.LogError("GenerateToken SignedString() Error: ", err)
		return "", err
	}
	return t, nil

}
