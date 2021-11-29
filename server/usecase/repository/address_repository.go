package repository

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
)

type AddressRepository interface {
	Store(ctx context.Context, instance *model.Address) (*model.Address, error)
	Update(ctx context.Context, instance *model.Address) (*model.Address, error)
	FindUnassigned(ctx context.Context, hostID int) ([]*model.Address, error)
}
