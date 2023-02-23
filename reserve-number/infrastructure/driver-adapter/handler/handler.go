package handler

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	port "github.com/devpablocristo/growuphr/reserve-number/application/port"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type Handler struct {
	numberService port.Service
}

func NewHandler(ps port.Service) *Handler {
	return &Handler{
		numberService: ps,
	}
}

func (h *Handler) ReserveNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close()

	// userName := r.URL.Query().Get("id")
	// if idQuery == "" {
	// 	return errors.New(idQuery)
	// }

	userName := chi.URLParam(r, "username")

	fmt.Println(userName)

	// userName := r.h.URLParam(r, "username") // 👈 getting path param
	// _, err := writer.Write([]byte("Hello " + username))
	// if err != nil {
	// 	log.Println(err)
	// }

	var newNumber domain.Number
	err := json.NewDecoder(body).Decode(&newNumber)
	if err != nil {
		responseErr := cmsapi.InvalidJSON("handler.AddNumber", err)
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
	number, err := h.numberService.AddNumber(ctx, &newNumber)
	if err != nil {
		responseErr := cmsapi.InternalServerError("handler.AddNumber", err)
		w.WriteHeader(responseErr.StatusCode)
		err = json.NewEncoder(w).Encode(cmsapi.FailResponse(responseErr))
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(responseErr)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(cmsapi.SuccessResponse(http.StatusCreated, number))
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

	fmt.Println("list")
}
