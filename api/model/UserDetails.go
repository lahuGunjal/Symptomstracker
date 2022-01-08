package model

type UserDetails struct {
	Name       string `json:"name" db:"name"`
	ContactNum int    `json:"contactNum" db:"contactNum"`
	EmailId    string `json:"emailId" db:"emailId"`
	Address    string `json:"address" db:"address"`
	LoginId    string `json:"loginId" db:"loginId"`
	PassWord   string `json:"password" db:"pwd"`
	UserId     string `json:"userId" db:"userId"`
}
