package address

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type UpdateAddressInputPort struct {
	ID         int
	InstanceID int
}

type UpdateAddressOutputPort struct {
	Address *model.Address
}

type UpdateAddressUseCase struct {
	addressRepo repository.AddressRepository
}

// func NewUpdateAddressUseCase(r repository.AddressRepository) *UpdateAddressUseCase {
// 	return &UpdateAddressUseCase{r}
// }

func (u *UpdateAddressUseCase) Execute(ctx context.Context, in *UpdateAddressInputPort) (*UpdateAddressOutputPort, error) {
	address := &model.Address{
		ID:         in.ID,
		InstanceID: in.InstanceID,
	}

	address, err := u.addressRepo.Update(ctx, address)
	if err != nil {
		return nil, err
	}

	return &UpdateAddressOutputPort{address}, nil
}
