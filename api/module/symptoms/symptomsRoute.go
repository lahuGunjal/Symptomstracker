package symptoms

import (
	"Lahu/GoExample/echoserve/api/model"
	"log"
	"net/http"

	"github.com/labstack/echo"
)

func Init(o *echo.Group, r *echo.Group) {
	r.POST("/addSymptoms", addSymptomsRoute)
}

//addSymptomsRoute to add dailysymptoms
func addSymptomsRoute(c echo.Context) error {
	symptomDetails := model.Symptoms{}
	// bind Parameter to symptoms struct
	bindErr := c.Bind(&symptomDetails)
	if bindErr != nil {
		log.Println("parameter binding error", bindErr)
		return c.JSON(http.StatusExpectationFailed, bindErr)
	}
	// call addSymptomsService to insert new symptoms in db
	status, addSymptomErr := addSymptomsService(symptomDetails)
	if addSymptomErr != nil {
		log.Println("addSymptomsRoute addSymptomsService", addSymptomErr)
		return c.JSON(http.StatusInternalServerError, addSymptomErr)
	}

	return c.JSON(http.StatusOK, status)
}
