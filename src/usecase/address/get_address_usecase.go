package address

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type GetAddressInputPort struct {
	InstanceID int
}

type GetAddressOutputPort struct {
	Address *model.Address
}

type GetAddressUseCase struct {
	addressRepo repository.AddressRepository
}

func NewGetAddressUseCase(r repository.AddressRepository) *GetAddressUseCase {
	return &GetAddressUseCase{r}
}

func (u *GetAddressUseCase) Execute(ctx context.Context, in *GetAddressInputPort) (*GetAddressOutputPort, error) {
	address, err := u.addressRepo.FindByInstanceID(ctx, in.InstanceID)
	if err != nil {
		return nil, err
	}
	return &GetAddressOutputPort{address}, nil
}
