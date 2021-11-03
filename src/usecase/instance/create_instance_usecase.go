package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type CreateInstanceInputPort struct {
	HostID  int
	Name    string
	Size    int
	Key     *model.Key
	Address *model.Address
}

type CreateInstanceOutputPort struct {
	Instance *model.Instance
}

type CreateInstanceUseCase struct {
	instanceRepo repository.InstanceRepository
}

func NewCreateInstanceUseCase(r repository.InstanceRepository) *CreateInstanceUseCase {
	return &CreateInstanceUseCase{r}
}

func (u *CreateInstanceUseCase) Execute(ctx context.Context, in *CreateInstanceInputPort) (*CreateInstanceOutputPort, error) {
	instance := &model.Instance{
		HostID:  in.HostID,
		Name:    in.Name,
		State:   model.Starting, // 仮想マシン起動中
		Size:    in.Size,
		Key:     in.Key,
		Address: in.Address,
	}

	instance, err := u.instanceRepo.Store(ctx, instance)
	if err != nil {
		return nil, err
	}

	return &CreateInstanceOutputPort{instance}, nil
}
