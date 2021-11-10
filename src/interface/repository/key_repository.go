package repository

import (
	"context"
	"database/sql"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type keyRepository struct {
	executor ContextExecutor
}

type TransactionalKeyRepository interface {
	repository.KeyRepository
	WithTx(tx *sql.Tx) *keyRepository
}

func (r *keyRepository) WithTx(tx *sql.Tx) *keyRepository {
	newRepo := *r
	newRepo.executor = tx
	return &newRepo
}

func NewKeyRepository(db *sql.DB) *keyRepository {
	return &keyRepository{db}
}

func (r *keyRepository) Store(ctx context.Context, key *model.Key) (*model.Key, error) {
	query := `INSERT INTO "keys" ("instance_id", "data") VALUES ($1, $2) RETURNING "id"`
	if err := r.executor.QueryRowContext(ctx, query, key.InstanceID, key.Data).Scan(&key.ID); err != nil {
		return nil, err
	}
	return key, nil
}
