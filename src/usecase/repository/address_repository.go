package repository

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
)

type AddressRepository interface {
	Update(ctx context.Context, instance *model.Address) (*model.Address, error)
	FindByInstanceID(ctx context.Context, instanceID int) (*model.Address, error)
}
