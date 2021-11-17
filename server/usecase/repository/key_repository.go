package repository

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
)

type KeyRepository interface {
	Store(ctx context.Context, instance *model.Key) (*model.Key, error)
	Delete(ctx context.Context, instanceID int) error
}
