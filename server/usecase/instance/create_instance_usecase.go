package instance

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type CreateInstanceInputPort struct {
	Name string
	Size int
	Key  *model.Key
	// Address *model.Address
}

type CreateInstanceOutputPort struct {
	Instance *model.Instance
}

type CreateInstanceUseCase struct {
	hostRepo     repository.HostRepository
	instanceRepo repository.InstanceRepository
	addressRepo  repository.AddressRepository
	keyRepo      repository.KeyRepository
}

func NewCreateInstanceUseCase(
	hr repository.HostRepository,
	ir repository.InstanceRepository,
	ar repository.AddressRepository,
	kr repository.KeyRepository) *CreateInstanceUseCase {
	return &CreateInstanceUseCase{hr, ir, ar, kr}
}

func (u *CreateInstanceUseCase) Execute(ctx context.Context, in *CreateInstanceInputPort) (*CreateInstanceOutputPort, error) {
	// Host を割り出す
	host, err := u.hostRepo.GetAvailableHost(ctx, in.Size)
	if err != nil {
		return nil, err
	}

	// Address を払い出す
	addresses, err := u.addressRepo.FindUnassigned(ctx, host.ID)
	if err != nil {
		return nil, err
	}

	// Instance を追加する
	instance := &model.Instance{
		HostID: host.ID,
		Name:   in.Name,
		State:  model.Starting, // 仮想マシン起動中
		Size:   in.Size,
	}

	instance, err = u.instanceRepo.Store(ctx, instance)
	if err != nil {
		return nil, err
	}

	address := &model.Address{
		ID:         addresses[0].ID,
		InstanceID: instance.ID,
		IPAddress:  addresses[0].IPAddress,
		MacAddress: addresses[0].MacAddress,
	}
	address, err = u.addressRepo.Update(ctx, address)
	if err != nil {
		return nil, err
	}

	// Key を作成する (dummy)
	key := &model.Key{
		InstanceID: instance.ID,
		Data:       "dummy",
	}
	key, err = u.keyRepo.Store(ctx, key)
	if err != nil {
		return nil, err
	}

	// Instance モデルに Address と Key を設定する
	instance.Address = address
	instance.Key = key

	return &CreateInstanceOutputPort{instance}, nil
}
