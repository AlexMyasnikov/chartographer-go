package api

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"internshipApplicationTemplate/pkg/db"
	"net/http"
)

func (c *ChartaHandler) DeleteCharta(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = c.Service.DeleteCharta(id)
	if errors.Is(db.ErrNotFound, err) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}
