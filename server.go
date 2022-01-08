package main

import (
	"Lahu/symptomsTracker/api"
	"net/http"

	"github.com/labstack/echo"
)

func helloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, World!")
}
func main() {
	e := echo.New()
	api.Init(e)
	e.GET("/", helloWorld)
	e.Logger.Fatal(e.Start(":1323"))
}
