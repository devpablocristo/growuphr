package application

import (
	"context"
	"time"

	uuid "github.com/google/uuid"

	port "github.com/devpablocristo/growuphr/reserve-number/application/port"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type UserService struct {
	storage port.Storage
}

func NewUserService(s port.Storage) *UserService {
	return &UserService{
		storage: s,
	}
}

func (ns *UserService) AddUser(ctx context.Context, u *domain.User) error {
	u.UUID = uuid.New().String()
	u.CreatedAt = time.Now()

	// err := ns.storage.AddNumber(ctx, u)
	// if err != nil {
	// 	return err
	// }
	return nil

}

func (ns *UserService) GetUser(ctx context.Context, userName string) (*domain.User, error) {
	u, err := ns.storage.GetUser(ctx, userName)
	if err != nil {
		return nil, err
	}
	return u, nil
}
