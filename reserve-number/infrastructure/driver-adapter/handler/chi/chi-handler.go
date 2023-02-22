package handler

import (
	"encoding/json"
	"log"
	"net/http"

	cmsapi "github.com/devpablocristo/growuphr/internal/commons/api"
	port "github.com/devpablocristo/growuphr/reserve-number/application/port"
	domain "github.com/devpablocristo/growuphr/reserve-number/domain"
)

type ChiHandler struct {
	numberService port.Service
}

func NewChiHandler(ps port.Service) *ChiHandler {
	return &ChiHandler{
		numberService: ps,
	}
}

func (h *ChiHandler) AddNumber(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	body := r.Body
	defer body.Close()

	var newNumber domain.Number
	err := json.NewDecoder(body).Decode(&newNumber)
	if err != nil {
		var errReq *cmsapi.APIError
		errReq.InvalidJSON("handler.AddNumber", err)
		w.WriteHeader(errReq.StatusCode)
		err = json.NewEncoder(w).Encode(cmsapi.FailResponse(errReq))
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(errReq)
		return
	}

	ctx := r.Context()
	number, err := h.numberService.AddNumber(ctx, &newNumber)

	if err != nil {
		var errReq *cmsapi.APIError
		errReq.InternalServerError("handler.AddNumber", err)
		w.WriteHeader(errReq.StatusCode)
		err = json.NewEncoder(w).Encode(cmsapi.FailResponse(errReq))
		if err != nil {
			log.Println(err.Error())
			w.Write([]byte(err.Error()))
			return
		}
		log.Println(errReq)
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
