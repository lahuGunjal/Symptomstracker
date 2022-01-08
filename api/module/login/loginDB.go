package login

import (
	"Lahu/symptomsTracker/api/model"
	postgressSQLDB "Lahu/symptomsTracker/common/mySQLDB"
	"errors"
	"log"

	"github.com/gocraft/dbr"
)

//GetDataFromSQLDB to Get Login data from LoginDetails SQL Table
func GetDataFromSQLDB(userDetails model.UserDetails) (model.UserDetails, error) {
	var session *dbr.Session
	userDetailsLocal := model.UserDetails{}
	connection, err := postgressSQLDB.GetSQLConnection()
	if err != nil {
		log.Println("error occured while getConnection", err)
		return model.UserDetails{}, err
	}
	log.Println(userDetails.LoginId)
	session = connection.NewSession(nil)
	TableName := "UserDetails"
	sql := "SELECT * FROM " + TableName + " WHERE loginId=? ;"

	count, Sqlerr := session.SelectBySql(sql, userDetails.LoginId).Load(&userDetailsLocal)
	if Sqlerr != nil {
		log.Println("ERROR : GetDataFromSQLDB : Error occured while GetSQLConnection ---> ", err)
		return model.UserDetails{}, Sqlerr
	}
	if count > 0 {
		log.Println("OUT : GetDataFromSQLDB", count, userDetailsLocal)
		return userDetailsLocal, nil
	}

	return model.UserDetails{}, errors.New("DATA_NOT_FOUND")
}
