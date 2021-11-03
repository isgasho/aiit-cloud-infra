package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type UpdateInstanceInputPort struct {
	ID    int
	State model.State
}

type UpdateInstanceOutputPort struct {
	Instance *model.Instance
}

type UpdateInstanceUseCase struct {
	instanceRepo repository.InstanceRepository
}

func NewUpdateInstanceUseCase(r repository.InstanceRepository) *UpdateInstanceUseCase {
	return &UpdateInstanceUseCase{r}
}

func (u *UpdateInstanceUseCase) Execute(ctx context.Context, in *UpdateInstanceInputPort) (*UpdateInstanceOutputPort, error) {
	instance := &model.Instance{
		ID:    in.ID,
		State: in.State,
	}

	instance, err := u.instanceRepo.Update(ctx, instance)
	if err != nil {
		return nil, err
	}

	return &UpdateInstanceOutputPort{instance}, nil
}
