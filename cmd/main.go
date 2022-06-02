package main

import (
	"github.com/rs/zerolog/log"
	"internshipApplicationTemplate/pkg/api"
	"internshipApplicationTemplate/pkg/db/charta"
	"internshipApplicationTemplate/pkg/service"
	"net/http"
)

var (
	addr = ":8080"
)

func main() {
	chartaDb := charta.NewCharta()
	chartaServ := service.NewChartaService(chartaDb)
	chartaHandler := &api.ChartaHandler{
		Service: chartaServ,
	}

	setupDir()

	mux := routes(chartaHandler)
	handler := recovery(mux)

	log.Print("Listening on port=", addr)
	err := http.ListenAndServe(addr, handler)
	if err != nil {
		log.Fatal().Err(err)
	}
}
