package service

import (
	"github.com/google/uuid"
	"golang.org/x/image/bmp"
	"image"
	"internshipApplicationTemplate/pkg/models"
	"os"
	"time"
)

func (c *ChartaService) AddCharta(width, height int) (uuid.UUID, error) {
	filename := "bmp_" + time.Now().Format(time.RFC3339) + ".bmp"

	fCharta, err := os.Create(filename)
	if err != nil {
		return uuid.UUID{}, err
	}

	rgba64 := image.NewRGBA64(image.Rect(width, height, 0, 0))
	err = bmp.Encode(fCharta, rgba64)
	if err != nil {
		return uuid.UUID{}, err
	}

	charta := &models.Charta{
		Id:   uuid.New(),
		Name: filename,
	}
	c.DB.AddCharta(charta)
	if err != nil {
		return uuid.UUID{}, err
	}
	return charta.Id, nil
}
