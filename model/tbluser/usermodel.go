package tbluser

import (
	"github.com/angadthandi/goapp/model/helper"
	"github.com/jinzhu/gorm"
)

type TblUser struct {
	gorm.Model
	UserID   helper.JsonNullInt64  `gorm:"column:UserID;primary_key" json:"UserID"`
	Username helper.JsonNullString `gorm:"column:Username" json:"Username"`
	Password helper.JsonNullString `gorm:"column:Password" json:"Password"`
	Email    helper.JsonNullString `gorm:"column:Email" json:"Email"`
}

// TableName sets the insert table name for this struct type
func (TblUser) TableName() string {
	return "tblUser"
}

// GetUserByID returns user struct by userID
func GetUserByID(
	db *gorm.DB,
	userID string,
) (TblUser, error) {
	var ret TblUser

	err := db.
		Where("UserID = ?", userID).
		Find(&ret).
		Error

	return ret, err
}

// ValidateLogin returns UserID based on username/password
func ValidateLogin(
	db *gorm.DB,
	username string,
	password string,
) (TblUser, error) {
	var ret TblUser

	// hash password

	err := db.
		Where("Username = ?", username).
		Where("Password = ?", password).
		Find(&ret).
		Error

	return ret, err
}
