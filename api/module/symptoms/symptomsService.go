package symptoms

import (
	"Lahu/GoExample/echoserve/api/model"
	"log"
)

//addSymptomsService to add new symptom details in database
func addSymptomsService(symptom model.Symptoms) (bool, error) {
	//call db method to store symptom details
	status, insertNewSymptomsErr := InsertNewSymptomsDB(symptom)
	if insertNewSymptomsErr != nil {
		log.Println("addSymptomsService InsertNewSymptomsDB()  ", insertNewSymptomsErr)
		return false, insertNewSymptomsErr
	}
	return status, nil
}
