package charta

import (
	"github.com/google/uuid"
	"internshipApplicationTemplate/pkg/models"
)

type Charta struct {
	db map[uuid.UUID]string
}

func NewCharta() *Charta {
	m := make(map[uuid.UUID]string)
	return &Charta{db: m}
}

func (c *Charta) AddCharta(charta *models.Charta) {
	c.db[charta.Id] = charta.Name
}
