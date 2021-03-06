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
	FindByState(ctx context.Context, state model.State) ([]*model.Instance, error)
}
