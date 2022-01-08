package symptoms

import (
	"Lahu/GoExample/echoserve/api/model"
	"Lahu/GoExample/echoserve/common/postgressSQLDB"
	"errors"
	"log"
	"time"

	"github.com/gocraft/dbr"
)

//InsertNewSymptomsDB insert new syptoms in mysql table Symptoms
func InsertNewSymptomsDB(symptom model.Symptoms) (bool, error) {
	var session *dbr.Session
	connection, err := postgressSQLDB.GetSQLConnection()
	if err != nil {
		log.Println("error occured while getConnection", err)
		return false, err
	}
	session = connection.NewSession(nil)
	layout := "2006-01-02"
	currentDate := time.Now()
	currentDate.Format(layout)
	result, Sqlerr := session.InsertInto("Symptoms").Columns("userId", "pulseRate", "bloodPressure", "otherSymptoms", "date").Values(
		symptom.UserId,
		symptom.PulseRate,
		symptom.BloodPressure,
		symptom.OtherSymptoms,
		currentDate,
	).Exec()
	if Sqlerr != nil {
		log.Println("ERROR : InsertNewSymptomsDB : Error occured while GetSQLConnection ---> INSERT ", Sqlerr)
		return false, Sqlerr
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(count, err)
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, errors.New("UNABLE_TO_INSERT_DATA")
}

//GetPastWeakSymptomDetailsDB get data of pastweakSymptom details from symtoms table
func GetPastWeakSymptomDetailsDB(userId string) ([]model.Symptoms, error) {
	symptoms := []model.Symptoms{}
	var session *dbr.Session

	connection, err := postgressSQLDB.GetSQLConnection()
	if err != nil {
		log.Println("error occured while getConnection", err)
		return symptoms, err
	}
	session = connection.NewSession(nil)
	TableName := "Symptoms"
	sql := "SELECT * FROM " + TableName + " WHERE userId=? AND date >= DATE_SUB(NOW(), INTERVAL 7 DAY);"
	count, Sqlerr := session.SelectBySql(sql, userId).Load(&symptoms)
	if Sqlerr != nil {
		log.Println("ERROR : GetPastWeakSymptomDetailsDB : Error occured while GetSQLConnection ---> ", err)
		return symptoms, Sqlerr
	}
	if count > 0 {
		log.Println("OUT : GetPastWeakSymptomDetailsDB", count, symptoms)
		return symptoms, nil
	}
	return symptoms, errors.New("NO_DATA_FOUND")
}
