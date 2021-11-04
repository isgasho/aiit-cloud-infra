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
	keyRepo      repository.KeyRepository
}

func NewDeleteInstanceUseCase(
	ir repository.InstanceRepository,
	kr repository.KeyRepository) *DeleteInstanceUseCase {
	return &DeleteInstanceUseCase{ir, kr}
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

	// TODO: Key を削除する
	// TODO: Instance モデルの Key を nil にする

	return &DeleteInstanceOutputPort{instance}, nil
}
