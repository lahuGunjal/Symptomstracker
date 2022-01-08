package login

import (
	"Lahu/symptomsTracker/api/model"
	JWTUtils "Lahu/symptomsTracker/api/module/jwtUtils"
	"errors"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Init(o *echo.Group, r *echo.Group) {
	o.POST("/loginUser", loginRoute)
}

func loginRoute(c echo.Context) error {
	userDetails := model.UserDetails{}
	//bind pramameter to userDetails Struct
	bindErr := c.Bind(&userDetails)
	if bindErr != nil {
		log.Println("PARAMETER_BINDING_ERROR", bindErr)
		return c.JSON(http.StatusExpectationFailed, errors.New("PARAMETER_BINDING_ERROR"))
	}
	if userDetails.LoginId == "" || userDetails.PassWord == "" {
		log.Println("LOGINID_OR_PASSWORD_BLANK")
		return c.JSON(http.StatusExpectationFailed, errors.New("LOGINID_OR_PASSWORD_BLANK"))
	}
	loginData, validateErr := ValidateUser(userDetails)
	if validateErr != nil {
		log.Println("Error occured while validating user", validateErr)
		return c.JSON(http.StatusExpectationFailed, errors.New("USER_VALIDATION_FAILED"))
	}
	if loginData.LoginId != "" {
		t, err := JWTUtils.GenerateTocken(userDetails)
		if err != nil {
			log.Println("Error In geerating tocken ", userDetails.LoginId, err)
			return c.JSON(http.StatusInternalServerError, errors.New("GENERATING_TOCKEN_ERROR"))
		}
		authResponse := model.AuthResponse{}
		authResponse.UserId = loginData.UserId
		authResponse.Tocken = t
		return c.JSON(http.StatusOK, authResponse)
	}
	return c.JSON(http.StatusExpectationFailed, errors.New("LOGIN_DEATAILS_NOT_FOUND"))
}
