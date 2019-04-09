package user

import (
	"github.com/angadthandi/goapp/model/helper"
	"github.com/angadthandi/goapp/model/tbluser"
	"github.com/jinzhu/gorm"
)

type SignupStruct struct {
	Username helper.JsonNullString `json:"Username"`
	Password helper.JsonNullString `json:"Password"`
	Email    helper.JsonNullString `json:"Email"`
}

type LoginStruct struct {
	Username helper.JsonNullString `json:"Username"`
	Password helper.JsonNullString `json:"Password"`
}

// Signup func creates new user,
// if username or email already exist, return error
func Signup(
	db *gorm.DB,
	data SignupStruct,
) (tbluser.TblUser, error) {
	var (
		ret tbluser.TblUser
		err error
	)

	// user exists

	// if not, create

	// hash password

	return ret, err
}
