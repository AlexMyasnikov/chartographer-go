package service

import (
	"github.com/google/uuid"
	"os"
)

func (c *ChartaService) DeleteCharta(id uuid.UUID) error {
	filename, err := c.DB.GetChartaName(id)
	if err != nil {
		return err
	}

	c.DB.DeleteCharta(id)

	if err = os.Remove(filename); err != nil {
		return err
	}

	return nil
}
