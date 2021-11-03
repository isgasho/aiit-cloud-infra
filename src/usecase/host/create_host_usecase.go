package host

import (
	"context"

	"github.com/mi-bear/infra-control/domain/model"
	"github.com/mi-bear/infra-control/usecase/repository"
)

type CreateHostInputPort struct {
	Name  string
	Limit int
}

type CreateHostOutputPort struct {
	Host *model.Host
}

type CreateHostUseCase struct {
	hostRepo repository.HostRepository
}

func NewCreateHostUseCase(r repository.HostRepository) *CreateHostUseCase {
	return &CreateHostUseCase{r}
}

func (u *CreateHostUseCase) Execute(ctx context.Context, in *CreateHostInputPort) (*CreateHostOutputPort, error) {
	host := &model.Host{
		Name:  in.Name,
		Limit: in.Limit,
	}

	host, err := u.hostRepo.Store(ctx, host)
	if err != nil {
		return nil, err
	}

	return &CreateHostOutputPort{host}, nil
}
