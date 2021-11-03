package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type DeleteInstanceInputPort struct {
	ID int
}

type DeleteInstanceOutputPort struct {
	Instance *model.Instance
}

type DeleteInstanceUseCase struct {
	instanceRepo repository.InstanceRepository
}

func NewDeleteInstanceUseCase(r repository.InstanceRepository) *DeleteInstanceUseCase {
	return &DeleteInstanceUseCase{r}
}

func (u *DeleteInstanceUseCase) Execute(ctx context.Context, in *DeleteInstanceInputPort) (*DeleteInstanceOutputPort, error) {
	instance := &model.Instance{
		ID:    in.ID,
		State: model.Terminating, // 仮想マシン削除中
	}

	instance, err := u.instanceRepo.Delete(ctx, instance)
	if err != nil {
		return nil, err
	}

	return &DeleteInstanceOutputPort{instance}, nil
}
