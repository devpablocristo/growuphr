package application

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/google/uuid"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	port "github.com/devpablocristo/growuphr/number-manager/application/port"
	domain "github.com/devpablocristo/growuphr/number-manager/domain"
)

type NumberManager struct {
	storage port.Storage
}

func NewNumberManager(st port.Storage) *NumberManager {
	return &NumberManager{
		storage: st,
	}
}

func (nm *NumberManager) AddReserveNumber(ctx context.Context, newResNum *domain.ReservedNumber) error {

	/*
		cada cliente puede reservar máximo un número.
		Una vez un número es reservado, no se debe poder reservar de nuevo, y el servidor debe
		arrojar un error adecuado indicando el motivo del error.

		caso 1:
			usuario nuevo - numero nuevo - ok
		caso 2:
			usuario existente - numero nuevo - ok - no se da, al crear el usuario se agreaga el numero
		caso 3:
			usuario nuevo - numero tomado - not ok
		caso 4:
			usuario existente - numero tomado -  not ok
		caso 5:
			usuario nuevo - numero 0 - not ok
		caso 6:
			usuario existente - numero 0 - not ok


		if existe usuario {
			if numero no disponible { // caso 4
				return numero no dispoblie
			}
		}

		// no existe el usuario
		if numero tomado { // 3
			return numero no disponible
		}

		retyrn ok // caso 1 y 2
	*/

	// check si numero igual 0
	if newResNum.Number.Number == 0 {
		// el 0 no es numero valido
		eMsg := fmt.Errorf("number '%v' is invalid, try with another one", newResNum.Number.Number)
		err := cmsapi.NewAPIError(200, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		return err
	}

	// check si usuario existe
	// numero diferente de 0, o sea, numero valido
	rn, foundUsr := nm.storage.CheckForUsername(ctx, newResNum.User.Username)
	if foundUsr {
		// si usuario existe, ya le fue añadido un numero,
		// el usuario ya tiene numero
		eMsg := fmt.Errorf("user '%v' has already reserved the number '%v'", rn.User.Username, rn.Number.Number)
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		return err
	}

	// el usuario no existe
	// numero diferente de 0, o sea, numero valido
	rn, foundNum := nm.storage.CheckForNumber(ctx, newResNum.Number.Number)
	if foundNum {
		// el numero existe,
		// numero invalido
		eMsg := fmt.Errorf("number '%v' already reserved, try with another one", rn.Number.Number)
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		return err
	}

	newResNum.UUID = uuid.New().String()
	newResNum.CreatedAt = time.Now()

	newResNum.Number.UUID = uuid.New().String()
	newResNum.Number.CreatedAt = newResNum.CreatedAt

	newResNum.User.UUID = uuid.New().String()
	newResNum.User.CreatedAt = newResNum.CreatedAt

	err := nm.storage.Create(ctx, newResNum)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	return nil
}

func (nm *NumberManager) ReservedNumbers(ctx context.Context) (map[string]*domain.ReservedNumber, error) {
	nm.storage.List(ctx)
	list := nm.storage.List(ctx)
	if len(list) == 0 {
		eMsg := errors.New("empty list")
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "ReservedNumbers", "application", eMsg)
		log.Println(err.Error())
		return nil, err
	}

	// for v := range list {
	// 	fmt.Println(v)
	// }

	// log.Println(list)
	return list, nil
}
