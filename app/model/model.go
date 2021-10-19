package model

import (
	"time"

	"github.com/gentildpinto/olist-api/config/database"
	"gorm.io/gorm"
)

var db *gorm.DB = database.GetDB()

type Base struct {
	ID        string         `sql:"type:uuid;primary_key;default:uuid_generate_v4()"`
	CreatedAt time.Time      `json:"-"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}
