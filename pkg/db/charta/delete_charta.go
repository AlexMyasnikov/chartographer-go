package charta

import (
	"github.com/google/uuid"
)

func (c *Charta) DeleteCharta(id uuid.UUID) {
	delete(c.db, id)
}
