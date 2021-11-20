package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type ListInstancesInputPort struct {
	State model.State
}

type ListInstancesOutputPort struct {
	Instances []*model.Instance
}

type ListInstancesUseCase struct {
	instanceRepo repository.InstanceRepository
}

func NewListInstancesUseCase(r repository.InstanceRepository) *ListInstancesUseCase {
	return &ListInstancesUseCase{r}
}

func (u *ListInstancesUseCase) Execute(ctx context.Context, in *ListInstancesInputPort) (*ListInstancesOutputPort, error) {
	var instances []*model.Instance
	var err error
	instances, err = u.instanceRepo.FindByState(ctx, in.State)
	if err != nil {
		return nil, err
	}

	return &ListInstancesOutputPort{Instances: instances}, nil
}
