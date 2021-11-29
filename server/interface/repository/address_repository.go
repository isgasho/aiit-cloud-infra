package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type addressRepository struct {
	executor ContextExecutor
}

type TransactionalAddressRepository interface {
	repository.AddressRepository
	WithTx(tx *sql.Tx) *addressRepository
}

func (r *addressRepository) WithTx(tx *sql.Tx) *addressRepository {
	newRepo := *r
	newRepo.executor = tx
	return &newRepo
}

func NewAddressRepository(db *sql.DB) *addressRepository {
	return &addressRepository{db}
}

func (r *addressRepository) Store(ctx context.Context, address *model.Address) (*model.Address, error) {
	query := `INSERT INTO "addresses" ("host_id", "ip_address", "mac_address") VALUES ($1, $2, $3) RETURNING "id"`
	log.Println(address)
	if err := r.executor.QueryRowContext(ctx, query, address.HostID, address.IPAddress, address.MacAddress).Scan(&address.ID); err != nil {
		return nil, err
	}
	return address, nil
}

func (r *addressRepository) Update(ctx context.Context, address *model.Address) (*model.Address, error) {
	query := `UPDATE "addresses" SET "instance_id" = $1 WHERE "id" = $2`
	result, err := r.executor.ExecContext(ctx, query, address.InstanceID, address.ID)
	if err != nil {
		return nil, err
	}
	rows, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rows != 1 {
		return nil, fmt.Errorf("expected to affect 1 row, affected %d", rows)
	}
	return address, nil
}

func (r *addressRepository) FindUnassigned(ctx context.Context) ([]*model.Address, error) {
	query := `SELECT "id", "host_id", ip_address", "mac_address" FROM "addresses" WHERE "instance_id" is null order by "created_at" limit 1`

	rows, err := r.executor.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			log.Println("FindUnassigned query rows close failed.")
			return
		}
	}(rows)

	addresses := make([]*model.Address, 0)

	for rows.Next() {
		address := &model.Address{}
		if err := rows.Scan(
			&address.ID,
			&address.HostID,
			&address.IPAddress,
			&address.MacAddress); err != nil {
			return nil, err
		}
		addresses = append(addresses, address)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return addresses, nil
}
