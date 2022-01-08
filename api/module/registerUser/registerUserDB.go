package resisterUser

import (
	"Lahu/symptomsTracker/api/model"
	mysqldb "Lahu/symptomsTracker/common/mySQLDB"
	"log"

	"github.com/gocraft/dbr"
)

func resisterUserDB(userDetails model.UserDetails) (bool, error) {
	var session *dbr.Session
	connection, err := mysqldb.GetSQLConnection()
	if err != nil {
		log.Println("error occured while getConnection", err)
		return false, err
	}
	session = connection.NewSession(nil)

	result, Sqlerr := session.InsertInto("UserDetails").Columns("userId", "name", "contactNum", "emailId", "address", "loginId", "pwd").Values(
		userDetails.UserId,
		userDetails.Name,
		userDetails.ContactNum,
		userDetails.EmailId,
		userDetails.Address,
		userDetails.LoginId,
		userDetails.PassWord,
	).Exec()
	if Sqlerr != nil {
		log.Println("ERROR : resisterUserDB : Error occured while GetSQLConnection ---> INSERT ", Sqlerr)
		return false, Sqlerr
	}
	count, err := result.RowsAffected()
	if err != nil {
		log.Println(count, err)
		return false, err
	}

	return true, nil
}
