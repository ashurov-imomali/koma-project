package handler

import (
	"main/internal/service"
	"net/http"
)

type Handler struct {
	Srv *service.Service
}

func NewHandler(srv *service.Service) *Handler {
	return &Handler{Srv: srv}
}

func (h *Handler) GetPings(w http.ResponseWriter, r *http.Request) {

}
