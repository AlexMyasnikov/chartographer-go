package charta

import (
	"github.com/google/uuid"
	"internshipApplicationTemplate/pkg/db"
)

func (c *Charta) GetChartaName(id uuid.UUID) (string, error) {
	filename, ok := c.db[id]
	if !ok {
		return "", db.ErrNotFound
	}

	return filename, nil
}
