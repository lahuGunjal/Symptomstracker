package resisterUser

import (
	"Lahu/symptomsTracker/api/model"
	"log"
	"net/http"

	"github.com/google/uuid"
	"github.com/labstack/echo"
)

func Init(o *echo.Group, r *echo.Group) {
	o.POST("/resisterUser", resisterUserRoute)

}

func resisterUserRoute(c echo.Context) error {
	userDetails := model.UserDetails{}
	RequestId, _ := uuid.NewUUID()

	bindErr := c.Bind(&userDetails)
	if bindErr != nil {
		log.Println("resisterUserRoute() Error while parameter binding", RequestId, bindErr)
	}
	userDetails.UserId = RequestId.String()
	status, createUserErr := createUserService(userDetails)
	if createUserErr != nil {
		log.Println("resisterUserRoute createUserService() ", createUserErr)
		return c.JSON(http.StatusInternalServerError, createUserErr)
	}
	return c.JSON(http.StatusOK, status)
}
