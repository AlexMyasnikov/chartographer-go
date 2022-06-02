package db

import (
	"errors"
	"github.com/google/uuid"
	"internshipApplicationTemplate/pkg/models"
)

type Charta interface {
	AddCharta(charta *models.Charta)

	GetChartaName(id uuid.UUID) (string, error)

	DeleteCharta(id uuid.UUID)
}

var ErrNotFound = errors.New("charta not found")
