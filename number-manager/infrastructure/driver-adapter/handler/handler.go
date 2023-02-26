package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	port "github.com/devpablocristo/growuphr/number-manager/application/port"
	domain "github.com/devpablocristo/growuphr/number-manager/domain"
)

type Handler struct {
	numberManager port.NumberManager
}

func NewHandler(nrs port.NumberManager) *Handler {
	return &Handler{
		numberManager: nrs,
	}
}

func (h *Handler) ReserveNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close()

	// stdlib query param
	// localhost:8080/api/v1/number-service/reserve?unername={client1}
	// userName := r.URL.Query().Get("id")
	// if userName == "" {
	// 	errMsg := "empty query param"
	// 	w.Write([]byte(errMsg))
	// 	log.Println(errMsg)
	// 	return
	// }

	// chi path param
	// localhost:8080/api/v1/number-service/reserve/client1
	newUserName := chi.URLParam(r, "username")
	if newUserName == "" {
		// aplicar api error
		errMsg := "empty path param"
		w.Write([]byte(errMsg))
		log.Println(errMsg)
		return
	}
	newUser := &domain.User{
		Username: newUserName,
	}

	newNumber := &domain.Number{} //domain.Number
	err := json.NewDecoder(body).Decode(newNumber)
	if err != nil {
		responseErr := cmsapi.InvalidJSON("AddNumber", "handler", err)
		w.WriteHeader(responseErr.StatusCode)
		err = json.NewEncoder(w).Encode(cmsapi.FailResponse(responseErr))
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(responseErr.Error())
		return
	}

	ctx := r.Context()
	newResNum := &domain.ReservedNumber{
		User:   newUser,
		Number: newNumber,
	}

	err = h.numberManager.AddReserveNumber(ctx, newResNum)
	if err != nil {
		var responseErr *cmsapi.APIError
		if errors.As(err, &responseErr) {
			w.WriteHeader(responseErr.StatusCode)
			err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(responseErr.StatusCode, responseErr.Message))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			return
		}
	}

	w.WriteHeader(http.StatusCreated)
	msg := fmt.Sprintf("number '%v' reserve for client '%v'", strconv.Itoa(newNumber.Number), newUser.Username)
	err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(http.StatusCreated, msg))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}
}

/*
lista de numeros y clientes asociados al numero
*/
func (h *Handler) ReservedNumbers(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	list, err := h.numberManager.ReservedNumbers(ctx)
	if err != nil {
		var responseErr *cmsapi.APIError
		if errors.As(err, &responseErr) {
			w.WriteHeader(responseErr.StatusCode)
			err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(responseErr.StatusCode, responseErr.Message))
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				log.Println(err.Error())
				w.Write([]byte(err.Error()))
				return
			}
			return
		}

		log.Println(err.Error())

		return
	}

	if list == nil {
		w.WriteHeader(http.StatusCreated)
		err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(http.StatusCreated, "empty list"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}

	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(http.StatusCreated, list))
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err.Error())
		w.Write([]byte(err.Error()))
		return
	}

	fmt.Println(list)
}
