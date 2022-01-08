package login

import (
	"Lahu/symptomsTracker/api/model"
	"log"
)

//ValidateUser to check user credential with database
func ValidateUser(userDetails model.UserDetails) (model.UserDetails, error) {
	// call GetDataFromSQLDB to get login data from SQL
	loginData, dbErr := GetDataFromSQLDB(userDetails)
	if dbErr != nil {
		log.Println("ValidateUser GetDataFromSQLDB() Error in Get Data From SQl", dbErr)
		return model.UserDetails{}, dbErr
	}

	if loginData.LoginId != "" {
		if loginData.LoginId == userDetails.LoginId && loginData.PassWord == userDetails.PassWord {
			return loginData, nil
		}
	}
	return model.UserDetails{}, nil
}
