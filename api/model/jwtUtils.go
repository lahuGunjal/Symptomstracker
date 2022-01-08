package model

import (
	jwt "github.com/dgrijalva/jwt-go"
)

const (
	//JwtKey for key of EMA
	JwtKey = "secret"
)

type JwtCustomClaims struct {
	LoginID string `json:"loginId,omitempty"`
	jwt.StandardClaims
}
