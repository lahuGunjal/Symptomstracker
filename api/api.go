package api

import (
	"Lahu/symptomsTracker/api/middleware"
	Login "Lahu/symptomsTracker/api/module/login"
	resisterUser "Lahu/symptomsTracker/api/module/registerUser"
	"Lahu/symptomsTracker/api/module/symptoms"

	"github.com/labstack/echo"
)

func Init(e *echo.Echo) {
	r := e.Group("/r")
	o := e.Group("/o")
	middleware.Init(e, r, o)
	Login.Init(o, r)
	resisterUser.Init(o, r)
	symptoms.Init(o, r)

}
