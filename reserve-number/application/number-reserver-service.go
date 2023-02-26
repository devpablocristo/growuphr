package application

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	uuid "github.com/google/uuid"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	port "github.com/devpablocristo/growuphr/reserve-number/application/port"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type ReservedNumberService struct {
	storage port.Storage
}

func NewReserveNumberService(st port.Storage) *ReservedNumberService {
	return &ReservedNumberService{
		storage: st,
	}
}

func (nrs *ReservedNumberService) AddReserveNumber(ctx context.Context, newResNum *domain.ReservedNumber) error {

	/*
		cada cliente puede reservar máximo un número.
		Una vez un número es reservado, no se debe poder reservar de nuevo, y el servidor debe
		arrojar un error adecuado indicando el motivo del error.

		caso 1:
			usuario nuevo - numero nuevo - ok
		caso 2:
			usuario existente - numero nuevo - ok
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
		log.Println(nrs.storage.List(ctx))

		return err
	}

	// check si usuario existe
	// numero diferente de 0, o sea, numero valido
	rn, foundUsr := nrs.storage.CheckForUsername(newResNum.User.Username)
	if foundUsr {
		// si usuario existe, ya le fue añadido un numero,
		// el usuario ya tiene numero
		eMsg := fmt.Errorf("user '%v' has already reserved the number '%v'", rn.User.Username, rn.Number.Number)
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		log.Println(nrs.storage.List(ctx))

		return err
	}

	// el usuario no existe
	// numero diferente de 0, o sea, numero valido
	rn, foundNum := nrs.storage.CheckForNumber(newResNum.Number.Number)
	if foundNum {
		fmt.Println("entr 2")
		// el numero existe,
		// numero invalido
		eMsg := fmt.Errorf("number '%v' already reserved, try with another one", rn.Number.Number)
		err := cmsapi.NewAPIError(http.StatusOK, "taken-number", "AddReserveNumber", "application", eMsg)
		log.Println(err.Error())
		log.Println(nrs.storage.List(ctx))
		return err
	}

	newResNum.UUID = uuid.New().String()
	newResNum.CreatedAt = time.Now()
	err := nrs.storage.Create(ctx, newResNum)
	if err != nil {
		log.Println(err.Error())
		log.Println(nrs.storage.List(ctx))
		return err
	}

	log.Println(nrs.storage.List(ctx))
	return nil
}

func (nrs *ReservedNumberService) ReservedNumbers(ctx context.Context) error {
	return nil
}
