package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type CreateInstanceInputPort struct {
	HostID int
	Name   string
	Size   int
	// Key     *model.Key
	// Address *model.Address
}

type CreateInstanceOutputPort struct {
	Instance *model.Instance
}

type CreateInstanceUseCase struct {
	instanceRepo repository.InstanceRepository
	addressRepo  repository.AddressRepository
	keyRepo      repository.KeyRepository
}

func NewCreateInstanceUseCase(
	ir repository.InstanceRepository,
	ar repository.AddressRepository,
	kr repository.KeyRepository) *CreateInstanceUseCase {
	return &CreateInstanceUseCase{ir, ar, kr}
}

func (u *CreateInstanceUseCase) Execute(ctx context.Context, in *CreateInstanceInputPort) (*CreateInstanceOutputPort, error) {
	instance := &model.Instance{
		HostID: in.HostID,
		Name:   in.Name,
		State:  model.Starting, // 仮想マシン起動中
		Size:   in.Size,
	}

	instance, err := u.instanceRepo.Store(ctx, instance)
	if err != nil {
		return nil, err
	}

	// TODO: Key を作成する
	// TODO: Address を払い出す

	// TODO: Instance モデルに Address と Key を設定する

	return &CreateInstanceOutputPort{instance}, nil
}
