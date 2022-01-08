package resisterUser

import (
	"Lahu/symptomsTracker/api/model"
	"log"
)

func createUserService(userDetails model.UserDetails) (bool, error) {
	status, resisterUserErr := resisterUserDB(userDetails)
	if resisterUserErr != nil {
		log.Println("createUserService() resisterUserDB", resisterUserErr)
		return false, resisterUserErr
	}
	return status, nil
}
