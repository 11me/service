package validator

import (
	"encoding/json"
	"errors"

	"github.com/go-playground/validator"
)

type Payment struct {
	Transaction  string `json:"transaction" validate:"required"`
	Currency     string `json:"currency" validate:"required"`
	Provider     string `json:"provider" validate:"required"`
	Amount       int    `json:"amount" validate:"required"`
	PaymentDt    int    `json:"payment_dt" validate:"required"`
	Bank         string `json:"bank" validate:"required"`
	DeliveryCost int    `json:"delivery_cost" validate:"gte=0"`
}

type Items struct {
	ChrtID     int    `json:"chrt_id" validate:"required"`
	Price      int    `json:"price" validate:"required"`
	Rid        string `json:"rid" validate:"required"`
	Name       string `json:"name" validate:"required"`
	Sale       int    `json:"sale" validate:"required"`
	Size       string `json:"size"`
	TotalPrice int    `json:"total_price" validate:"required"`
}

type Order struct {
	OrderUID          string  `json:"order_uid" validate:"required"`
	Entry             string  `json:"entry" validate:"required"`
	InternalSignature string  `json:"internal_signature" validate:"required"`
	Payment           Payment `json:"payment" validate:"required"`
	Items             []Items `json:"items" validate:"required"`
	Locale            string  `json:"locale" validate:"required"`
	CustomerID        string  `json:"customer_id" validate:"required"`
	TrackNumber       string  `json:"track_number" validate:"required"`
	DeliveryService   string  `json:"delivery_service" validate:"required"`
}

var order Order

func isJSON(b []byte) bool {
	return json.Unmarshal(b, &order) == nil
}

func ValidateData(data []byte) error {

	if !isJSON(data) {
		return errors.New("not a valid JSON")
	}

	validate := validator.New()
	err := validate.Struct(order)

	if err != nil {
		return err
	}
	return nil
}
