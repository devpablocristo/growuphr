package port

import (
	"context"

	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type User interface {
	AddUser(context.Context, *domain.User) error
	GetUser(context.Context, string) (*domain.User, error)
}
