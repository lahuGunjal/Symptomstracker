package model

import (
	"github.com/gocraft/dbr"
)

type AuthResponse struct {
	UserId string `json:"userId"`
	Tocken string `json:"tocken"`
}
type LoginSQL struct {
	LoginId  dbr.NullString `json:"loginId" db:"loginId"`
	PassWord dbr.NullString `json:"pwd" db:"pwd"`
}
