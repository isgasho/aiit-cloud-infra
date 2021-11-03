package repository

import (
	"context"
	"database/sql"
	"fmt"

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
	query := `INSERT INTO addresses (ip_address, mac_address) VALUES ($1, $2) RETURNING id`
	if err := r.executor.QueryRowContext(ctx, query, address.IPAddress, address.MacAddress).Scan(&address.ID); err != nil {
		return nil, err
	}
	return address, nil
}

func (r *addressRepository) Update(ctx context.Context, address *model.Address) (*model.Address, error) {
	query := `UPDATE addresses SET instance_id = $1 WHERE id = $2`
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

// TODO: Address の払出し
// 割り当てのない Address を探して ID を返却する
