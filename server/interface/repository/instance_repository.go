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
	key := &model.Key{}
	address := &model.Address{}

	// address は updated_at を設けて重複した場合は、最も新しいもの (最近紐付けが行われたもの) を採用する
	query := `
SELECT i.id, i.host_id, i.name, i.state, i.size, k.id, k.data, a.id, a.ip_address, a.mac_address
FROM
  "instances" i
    LEFT JOIN (
      SELECT * FROM "keys" WHERE "instance_id" = $1 ORDER BY "created_at" DESC LIMIT 1) k ON (i.id = k.instance_id)
    LEFT JOIN (
      SELECT * FROM "addresses" WHERE "instance_id" = $1 ORDER BY "created_at" DESC LIMIT 1) a ON (i.id = a.instance_id)
WHERE "id" = $1
`
	if err := r.executor.QueryRowContext(ctx, query, id).Scan(
		&instance.ID,
		&instance.HostID,
		&instance.Name,
		&instance.State,
		&instance.Size,
		&key.ID,
		&key.Data,
		&address.ID,
		&address.IPAddress,
		&address.MacAddress); err != nil {
		return nil, err
	}
	key.InstanceID = instance.ID
	address.InstanceID = instance.ID

	instance.Key = key
	instance.Address = address

	return instance, nil
}

func (r *instanceRepository) FindByState(ctx context.Context, state model.State) ([]*model.Instance, error) {
	query := `
SELECT i.id, i.host_id, i.name, i.state, i.size, k.id, k.data, a.id, a.ip_address, a.mac_address
FROM
  "instances" i
    LEFT JOIN "keys" k ON (i.id = k.instance_id)
    LEFT JOIN "addresses" a ON (i.id = a.instance_id)
WHERE "state" = $1
`

	rows, err := r.executor.QueryContext(ctx, query, state)
	if err != nil {
		return nil, err
	}
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			return
		}
	}(rows)

	instances := make([]*model.Instance, 0)

	for rows.Next() {
		instance := &model.Instance{}
		key := &model.Key{}
		address := &model.Address{}
		var keyID, keyData interface{}
		if err := rows.Scan(
			&instance.ID,
			&instance.HostID,
			&instance.Name,
			&instance.State,
			&instance.Size,
			&keyID,
			&keyData,
			&address.ID,
			&address.IPAddress,
			&address.MacAddress); err != nil {
			return nil, err
		}

		if keyID != nil {
			key.ID = int(keyID.(int64))
		}
		if keyData != nil {
			key.Data = keyData.(string)
		}

		key.InstanceID = instance.ID
		address.InstanceID = instance.ID

		instance.Key = key
		instance.Address = address
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
