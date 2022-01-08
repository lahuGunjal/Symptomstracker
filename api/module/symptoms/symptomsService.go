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

//GetPastWeakSymptomDetailsService Get past weak symptoms data from sql
func GetPastWeakSymptomDetailsService(userId string) ([]model.Symptoms, error) {
	symptoms := []model.Symptoms{}
	symptoms, getSymptomFromDBErr := GetPastWeakSymptomDetailsDB(userId)
	if getSymptomFromDBErr != nil {
		log.Println("GetPastWeakSymptomDetailsService GetPastWeakSymptomDetailsDB", getSymptomFromDBErr)
		return symptoms, getSymptomFromDBErr
	}

	return symptoms, nil
}
