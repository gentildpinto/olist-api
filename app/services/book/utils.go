package book

import (
	"encoding/json"
	"reflect"

	"github.com/gentildpinto/olist-api/app/dto"
	"github.com/gentildpinto/olist-api/app/model"
)

func setPayloadValues(book *model.Book, payload dto.UpdateBook) (err error) {
	reflectedPayload := reflect.ValueOf(payload)
	typeOfStruct := reflectedPayload.Type()

	values := make(map[string]interface{})

	for index := 0; index < reflectedPayload.NumField(); index++ {
		if reflectedPayload.Field(index).Interface() == "" || reflectedPayload.Field(index).Interface() == 0 {
			continue
		}

		fieldName := typeOfStruct.Field(index).Name
		fieldValue := reflectedPayload.Field(index).Interface()

		values[fieldName] = fieldValue
	}

	jsonString, err := json.Marshal(values)
	json.Unmarshal(jsonString, &book)

	return
}
