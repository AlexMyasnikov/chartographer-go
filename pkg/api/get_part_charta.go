package api

import (
	"errors"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"internshipApplicationTemplate/pkg/db"
	"net/http"
	"os"
	"strconv"
)

const (
	maxPartWidth  = 5000
	maxPartHeight = 5000
)

func (c *ChartaHandler) GetPartCharta(w http.ResponseWriter, r *http.Request) {
	x, err := strconv.Atoi(r.URL.Query().Get("x"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	y, err := strconv.Atoi(r.URL.Query().Get("y"))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

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

	if height > maxPartHeight || width > maxPartWidth {
		http.Error(w, http.StatusText(http.StatusPreconditionFailed), http.StatusPreconditionFailed)
		return
	}

	id, err := uuid.Parse(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}
	buffer, err := c.Service.GetPartCharta(x, y, width, height, id)
	if errors.Is(db.ErrNotFound, err) || errors.Is(os.ErrNotExist, err) {
		http.NotFound(w, r)
		return
	}
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "image/bmp")
	w.Header().Set("Content-Length", strconv.Itoa(len(buffer.Bytes())))
	w.WriteHeader(http.StatusOK)

	w.Write(buffer.Bytes())
}
