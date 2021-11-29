package repository

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
)

type HostRepository interface {
	Store(ctx context.Context, instance *model.Host) (*model.Host, error)
	GetAvailableHost(ctx context.Context, size int) (*model.Host, error)
}
