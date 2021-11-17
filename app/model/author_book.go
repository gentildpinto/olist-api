package model

import (
	"time"

	"github.com/gentildpinto/olist-api/config/logger"
	"gorm.io/gorm"
)

type AuthorBook struct {
	AuthorID  UUID           `gorm:"primaryKey"`
	BookID    UUID           `gorm:"primaryKey"`
	CreatedAt time.Time      `gorm:"not null"`
	UpdatedAt time.Time      `json:"-"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

func (AuthorBook) Create(authorID, bookID UUID) (err error) {
	authorBook := AuthorBook{
		AuthorID: authorID,
		BookID:   bookID,
	}

	if err = logger.Log(databaseConnection.Create(&authorBook).Error); err != nil {
		return err
	}

	return
}
