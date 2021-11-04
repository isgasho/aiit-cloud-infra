package repository

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type instanceRepository struct {
	executor ContextExecutor
}

type TransactionalInstanceRepository interface {
	repository.InstanceRepository
	WithTx(tx *sql.Tx) *instanceRepository
}

func (r *instanceRepository) WithTx(tx *sql.Tx) *instanceRepository {
	newRepo := *r
	newRepo.executor = tx
	return &newRepo
}

func NewInstanceRepository(db *sql.DB) *instanceRepository {
	return &instanceRepository{db}
}

func (r *instanceRepository) Store(ctx context.Context, instance *model.Instance) (*model.Instance, error) {
	query := `INSERT INTO instances (host_id, name, state, size) VALUES ($1, $2, $3, $4) RETURNING id`
	if err := r.executor.QueryRowContext(
		ctx, query,
		instance.HostID,
		instance.Name,
		instance.State,
		instance.Size).Scan(&instance.ID); err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *instanceRepository) Update(ctx context.Context, instance *model.Instance) (*model.Instance, error) {
	query := `UPDATE instances SET state = $1 WHERE id = $2`
	result, err := r.executor.ExecContext(ctx, query, instance.State, instance.ID)
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
	return instance, nil
}

func (r *instanceRepository) Delete(ctx context.Context, instance *model.Instance) (*model.Instance, error) {
	query := `UPDATE instances SET state = $1 WHERE id = $2`
	result, err := r.executor.ExecContext(ctx, query, instance.State, instance.ID)
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
	return instance, nil
}

func (r *instanceRepository) FindByID(ctx context.Context, id int) (*model.Instance, error) {
	instance := &model.Instance{}

	// TODO: address は updated_at を設けて重複した場合は、最も新しいもの (最近紐付けが行われたもの) を採用する
	query := `
SELECT i.id, i.host_id, i.name, i.state, i.size, k.data, a.ip_address, a.mac_address
FROM
  instances i
    LEFT JOIN (
      SELECT * FROM keys WHERE instance_id = $1 ORDER BY created_at desc limit 1) k ON (i.id = k.instance_id)
    LEFT JOIN (
      SELECT * FROM address WHERE instance_id = $1 ORDER BY created_at desc limit 1) a ON (i.id = a.instance_id)
WHERE id = $1
`
	if err := r.executor.QueryRowContext(ctx, query, id).Scan(
		&instance.ID,
		&instance.Name,
		&instance.State,
		&instance.Size,
		&instance.Key.Data,
		&instance.Address.IPAddress,
		&instance.Address.MacAddress); err != nil {
		return nil, err
	}
	return instance, nil
}

func (r *instanceRepository) FindByHostID(ctx context.Context, hostID int) ([]*model.Instance, error) {
	// TODO: implement
	query := ``

	rows, err := r.executor.QueryContext(ctx, query, hostID)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			// TODO: Logging
			return
		}
	}(rows)

	instances := make([]*model.Instance, 0)

	for rows.Next() {
		instance := &model.Instance{}
		if err := rows.Scan(
			&instance.ID,
			&instance.Name,
			&instance.State,
			&instance.Size,
			&instance.Key.Data,
			&instance.Address.IPAddress,
			&instance.Address.MacAddress); err != nil {
			return nil, err
		}
		instances = append(instances, instance)
	}

	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return instances, nil
}
