package api

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"net/http"
	"strconv"
)

const (
	maxWidth  = 20000
	maxHeight = 50000
)

func (c *ChartaHandler) AddCharta(w http.ResponseWriter, r *http.Request) {
	width, err := strconv.Atoi(r.URL.Query().Get("width"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	height, err := strconv.Atoi(r.URL.Query().Get("height"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	if width > maxWidth || height > maxHeight || width <= 0 || height <= 0 {
		log.Error().Msg("Ширина/длина не удовлетворяют требованиям")
		http.Error(w, http.StatusText(http.StatusPreconditionFailed), http.StatusPreconditionFailed)
		return
	}

	id, err := c.Service.AddCharta(width, height)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(id)
	if err != nil {
		log.Error().Msg("Ошибка маршалинга")
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusCreated)

	w.Write(resp)
}
