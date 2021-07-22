package service

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/11me/wb_service/config"
	"github.com/11me/wb_service/validator"
	"github.com/jackc/pgx/v4/pgxpool"
)

type service struct {
	pool *pgxpool.Pool
}

func New() *service {
	return new(service)
}

func (svc *service) Connect(ctx context.Context, cfg config.Config) error {
	connStr := fmt.Sprintf("postgres://%v:%v@%v:5432/%v",
		cfg.DBUser,
		cfg.DBPass,
		cfg.DBUri,
		cfg.DBName,
	)

	pool, err := pgxpool.Connect(ctx, connStr)

	if err != nil {
		return err
	}

	// assign a pool
	svc.pool = pool

	return nil
}

// NOTE: maybe good idea to not hardcode the table name
func (db *service) PopulateJSON(ctx context.Context, data []byte) error {

	queryStr := fmt.Sprintf(`
		INSERT INTO public.order
		SELECT *
		FROM json_populate_record(NULL::public.order, %v)
	`, string(data))

	_, err := db.pool.Exec(ctx, queryStr)

	if err != nil {
		return err
	}

	return nil
}

func (svc *service) GetOrder(ctx context.Context, id string) (interface{}, error) {

	var order validator.Order
	// unmarshal JSON returned by DB into the string
	var payment string
	// save json array
	var items string

	query := fmt.Sprintf(`
		SELECT order_uid, entry, internal_signature, to_json(payment),
					 to_json(items), locale, customer_id, track_number, delivery_service
		FROM "order"
		WHERE order_uid='%v'
	`, id)

	rows, err := svc.pool.Query(ctx, query)

	for rows.Next() {
		//NOTE: how to deal with items and payment?
		defer rows.Close()

		err := rows.Scan(
			&order.OrderUID,
			&order.Entry,
			&order.InternalSignature,
			&payment,
			&items,
			&order.Locale,
			&order.CustomerID,
			&order.TrackNumber,
			&order.DeliveryService,
		)

		// unmarshal payment and items to struct
		json.Unmarshal([]byte(payment), &order.Payment)
		json.Unmarshal([]byte(items), &order.Items)

		if err != nil {
			return nil, err
		}
	}

	if err != nil {
		return nil, err
	}

	return order, nil
}

func (svc *service) Close() {
	svc.pool.Close()
}
