package service_test

import (
	"context"
	"log"
	"testing"

	"github.com/11me/wb_service/config"
	"github.com/11me/wb_service/service"
	"github.com/11me/wb_service/validator"
)

var path string = "../config/config.json"

func TestConnect(t *testing.T) {
	cfg, err := config.ParseConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.New()

	ctx := context.Background()
	err = svc.Connect(ctx, *cfg)

	if err != nil {
		log.Fatal(err)
	}
}

func TestPopulateJSON(t *testing.T) {
	cfg, _ := config.ParseConfig(path)
	// json data
	data := []byte(`
			'{
			"order_uid": "bfc4e94948fe433f8b55b68a0f169e56",
			"entry": "WBIL",
			"internal_signature": "8062ffbe87d2446dbc98b05a0e4e07bd53565c0b2fd90c6c6e6a96594de61433",
			"payment": {
				"transaction": "afc4e94948fe433f8b55b68a0f169e56",
				"currency": "USD",
				"provider": "WbPay",
				"amount": 3107,
				"payment_dt": 1614540555,
				"bank": "alpha",
				"delivery_cost": 0
			},
			"items": [
				{
					"chrt_id": 11241243,
					"price": 300,
					"rid": "7ed1002bcbc14fa594d48b422ef6980d",
					"name": "Сказки малышам. (библиотека детского сада).",
					"sale": 45,
					"size": "",
					"total_price": 164
				}
			],
			"locale": "ru",
			"customer_id": "5ea488619943420eaefcbcc402eb8ddc",
			"track_number": "WBIL2817015795SL",
			"delivery_service": "meest"
		}'
	`)

	ctx := context.Background()

	svc := service.New()
	svc.Connect(ctx, *cfg)
	svc.PopulateJSON(ctx, data)
}

func TestGetOrder(t *testing.T) {
	id := "afc4e94948fe433f8b55b68a0f169e56"
	cfg, err := config.ParseConfig(path)
	if err != nil {
		log.Fatal(err)
	}

	svc := service.New()

	ctx := context.Background()
	err = svc.Connect(ctx, *cfg)

	if err != nil {
		log.Fatal(err)
	}

	//var res interface{}
	res, err := svc.GetOrder(ctx, id)
	if err != nil {
		log.Fatal(err)
	}
	// extract the order from interface
	i := res.(validator.Order)

	log.Printf("res is: %v", i.Items[0].Price)
}
