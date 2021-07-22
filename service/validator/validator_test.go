package validator_test

import (
	"log"
	"testing"

	"github.com/11me/wb_service/validator"
)

func TestValidateData(t *testing.T) {
	data := []byte(`{"key": 12}`)

	// test for invalid json data
	err := validator.ValidateData(data)

	if err != nil {
		log.Print(err)
	}
}
