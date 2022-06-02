package main

import (
	"github.com/gorilla/mux"
	"internshipApplicationTemplate/pkg/api"
	"net/http"
)

func routes(chartaHandler *api.ChartaHandler) *mux.Router {
	mux := mux.NewRouter()

	mux.HandleFunc("/chartas/", chartaHandler.AddCharta).Methods(http.MethodPost)
	mux.HandleFunc("/chartas/{id}/", chartaHandler.AddPartCharta).Methods(http.MethodPost)
	mux.HandleFunc("/chartas/{id}/", chartaHandler.GetPartCharta).Methods(http.MethodGet)
	mux.HandleFunc("/chartas/{id}/", chartaHandler.DeleteCharta).Methods(http.MethodDelete)

	return mux
}
