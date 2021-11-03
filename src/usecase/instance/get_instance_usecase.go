package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type GetInstanceInputPort struct {
	ID     int
	HostID int
}

type GetInstanceOutputPort struct {
	Instance *model.Instance
}

type GetInstanceUseCase struct {
	instanceRepo repository.InstanceRepository
}

func NewGetInstanceUseCase(r repository.InstanceRepository) *GetInstanceUseCase {
	return &GetInstanceUseCase{r}
}

func (u *GetInstanceUseCase) Execute(ctx context.Context, in *GetInstanceInputPort) (*GetInstanceOutputPort, error) {
	instance, err := u.instanceRepo.FindByID(ctx, in.ID)
	if err != nil {
		return nil, err
	}
	return &GetInstanceOutputPort{instance}, nil
}
