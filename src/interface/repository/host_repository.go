package repository

import (
	"context"
	"database/sql"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type hostRepository struct {
	executor ContextExecutor
}

type TransactionalHostRepository interface {
	repository.HostRepository
	WithTx(tx *sql.Tx) *hostRepository
}

func (r *hostRepository) WithTx(tx *sql.Tx) *hostRepository {
	newRepo := *r
	newRepo.executor = tx
	return &newRepo
}

func NewHostRepository(db *sql.DB) *hostRepository {
	return &hostRepository{db}
}

func (r *hostRepository) Store(ctx context.Context, host *model.Host) (*model.Host, error) {
	query := `INSERT INTO hosts (name, limit) VALUES ($1, $2) RETURNING id`
	if err := r.executor.QueryRowContext(ctx, query, host.Name, host.Limit).Scan(&host.ID); err != nil {
		return nil, err
	}
	return host, nil
}
