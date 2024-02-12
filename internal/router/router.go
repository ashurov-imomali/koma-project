package router

import (
	"github.com/gorilla/mux"
	"main/internal/handler"
	"net/http"
)

func NewRouter(handler *handler.Handler) *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/ping", handler.GetPings).Methods(http.MethodGet)
	
	return mux

}
