package key

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type CreateKeyInputPort struct {
	InstanceID int
	Data       string
}

type CreateKeyOutputPort struct {
	Key *model.Key
}

type CreateKeyUseCase struct {
	keyRepo repository.KeyRepository
}

func NewCreateKeyUseCase(r repository.KeyRepository) *CreateKeyUseCase {
	return &CreateKeyUseCase{r}
}

func (u *CreateKeyUseCase) Execute(ctx context.Context, in *CreateKeyInputPort) (*CreateKeyOutputPort, error) {
	key := &model.Key{
		InstanceID: in.InstanceID,
		Data:       in.Data,
	}

	key, err := u.keyRepo.Store(ctx, key)
	if err != nil {
		return nil, err
	}

	return &CreateKeyOutputPort{key}, nil
}
