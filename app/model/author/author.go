package author

import (
	"errors"

	"github.com/gentildpinto/olist-api/app/model"
	"gorm.io/gorm"
)

var db *gorm.DB

type Author struct {
	model.Base
	Name string ` json:"name" gorm:"unique"`
}

func Initialize(dbConn *gorm.DB) (err error) {
	if dbConn == nil {
		return errors.New("db connection is nil")
	}
	db = dbConn

	return
}

func GetAll() (authors []Author, err error) {
	err = db.Find(&authors).Error
	return
}
