package repository

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
)

type InstanceRepository interface {
	Store(ctx context.Context, instance *model.Instance) (*model.Instance, error)
	Update(ctx context.Context, instance *model.Instance) (*model.Instance, error)
	Delete(ctx context.Context, instance *model.Instance) (*model.Instance, error)
	FindByID(ctx context.Context, id int) (*model.Instance, error)
	FindByHostID(ctx context.Context, hostID int) ([]*model.Instance, error)
}
