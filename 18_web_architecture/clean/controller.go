package main

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"net/http"
)

type Controller interface {
	Get(w http.ResponseWriter, r *http.Request)
}

type userController struct {
	service Service
}

func (c *userController) Get(w http.ResponseWriter, r *http.Request) {
	param := chi.URLParam(r, "uuid")

	user, err := c.service.Get(param)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

	payload, err := json.Marshal(&user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(payload)
}

func NewController(service Service) Controller {
	return &userController{
		service: service,
	}
}
