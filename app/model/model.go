package model

import (
	"database/sql/driver"
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type (
	UUID uuid.UUID

	Base struct {
		ID        UUID           `json:"id" gorm:"primary_key;default:(UUID_TO_BIN(UUID()))"`
		CreatedAt time.Time      `json:"-"`
		UpdatedAt time.Time      `json:"-"`
		DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	}
)

func (field UUID) String() string {
	return uuid.UUID(field).String()
}

func (field UUID) GormDataType() string {
	return "binary(16)"
}

func (field UUID) MarshalJSON() ([]byte, error) {
	s := uuid.UUID(field)
	str := "\"" + s.String() + "\""

	return []byte(str), nil
}

func (field *UUID) UnmarshalJSON(bytes []byte) (err error) {
	s, err := uuid.ParseBytes(bytes)
	*field = UUID(s)

	return
}

func (field *UUID) Scan(value interface{}) (err error) {
	bytes, _ := value.([]byte)
	parseByte, err := uuid.FromBytes(bytes)
	*field = UUID(parseByte)

	return
}

func (field UUID) Value() (driver.Value, error) {
	return uuid.UUID(field).MarshalBinary()
}
