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
	query := `INSERT INTO "hosts" ("name", "limit") VALUES ($1, $2) RETURNING "id"`
	if err := r.executor.QueryRowContext(ctx, query, host.Name, host.Limit).Scan(&host.ID); err != nil {
		return nil, err
	}
	return host, nil
}

func (r *hostRepository) GetAvailableHost(ctx context.Context, size int) (*model.Host, error) {
	host := &model.Host{}

	// TODO: クエリではなくロジックで解決する
	query := `SELECT total.id FROM
(SELECT
  h.id,
  h.limit,
  COALESCE(SUM(i.size), 0)::INTEGER + $1 AS sum
FROM
  hosts h LEFT JOIN instances i ON (h.id = i.host_id)
GROUP BY h.id, h.limit) total
WHERE
  total.limit > total.sum
ORDER BY (total.limit - total.sum) DESC LIMIT 1
`
	if err := r.executor.QueryRowContext(ctx, query, size).Scan(&host.ID); err != nil {
		return nil, err
	}
	return host, nil
}
