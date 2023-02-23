package application

import (
	"context"

	port "github.com/devpablocristo/growuphr/reserve-number/application/port"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type ReserveNumberService struct {
	userService   port.User
	numberService port.Number
}

func NewReserveNumberService(us port.User, ns port.Number) *ReserveNumberService {
	return &ReserveNumberService{
		userService:   us,
		numberService: ns,
	}
}

func (ns *ReserveNumberService) ReserveNumber(ctx context.Context, n *domain.Number, userName string) error {
	// 1.1 verificar si el usuario existe
	// u, err := ns.GetUser(ctx, n.UserName)
	// if err != nil {
	// 	// 1.2 verificar si numero fue ya reservado
	// 	fmt.Println(u)

	// 	return nil
	// }

	// 2.1 si no existe crear el usuario
	// err = ns.AddUser(ctx, n)
	// if err != nil {
	// 	return nil
	// }

	// return nil
}

func (rns *ReserveNumberService) ReservedNumbers(ctx context.Context) error {
	return nil
}
