package address

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type GetAddressOutputPort struct {
	Address *model.Address
}

type GetAddressUseCase struct {
	addressRepo repository.AddressRepository
}

func NewGetAddressUseCase(r repository.AddressRepository) *GetAddressUseCase {
	return &GetAddressUseCase{r}
}

func (u *GetAddressUseCase) Execute(ctx context.Context) (*GetAddressOutputPort, error) {
	address, err := u.addressRepo.FindUnassigned(ctx)
	if err != nil {
		return nil, err
	}
	return &GetAddressOutputPort{address}, nil
}
