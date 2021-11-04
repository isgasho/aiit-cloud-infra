package address

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type CreateAddressInputPort struct {
	IPAddress  string
	MacAddress string
}

type CreateAddressOutputPort struct {
	Address *model.Address
}

type CreateAddressUseCase struct {
	addressRepo repository.AddressRepository
}

func NewCreateAddressUseCase(r repository.AddressRepository) *CreateAddressUseCase {
	return &CreateAddressUseCase{r}
}

func (u *CreateAddressUseCase) Execute(ctx context.Context, in *CreateAddressInputPort) (*CreateAddressOutputPort, error) {
	address := &model.Address{
		IPAddress:  in.IPAddress,
		MacAddress: in.MacAddress,
	}

	address, err := u.addressRepo.Store(ctx, address)
	if err != nil {
		return nil, err
	}

	return &CreateAddressOutputPort{address}, nil
}
