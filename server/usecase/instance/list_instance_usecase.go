package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type ListInstancesInputPort struct {
	HostID int
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
	instances, err = u.instanceRepo.FindByHostID(ctx, in.HostID)
	if err != nil {
		return nil, err
	}

	return &ListInstancesOutputPort{Instances: instances}, nil
}
