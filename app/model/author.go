package model

type Author struct {
	Base
	Name string `json:"name" gorm:"unique"`
}

func (Author) All() (authors []Author, err error) {
	if err = databaseConnection.Find(&authors).Error; err != nil {
		return []Author{}, err
	}

	return
}
