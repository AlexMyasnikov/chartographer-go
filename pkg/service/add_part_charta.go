package service

import (
	"github.com/google/uuid"
	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	"image"
	"os"
)

// dst - исходный папирус, на который будет накладываться изображение

// src - восстановленный фрагмент

// r - границы, в которые будет вставлен фрагмент,
// где Min - координаты (x, y) относительно левого верхнего угла dst

func (c *ChartaService) AddPartCharta(x, y, width, height int, img image.Image, id uuid.UUID) error {
	dstFilename, err := c.DB.GetChartaName(id)
	if err != nil {
		return err
	}

	dstImg, err := os.Open(dstFilename)
	if err != nil {
		return err
	}

	dst, err := bmp.Decode(dstImg)
	if err != nil {
		return err
	}

	r := image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + width, Y: y + height},
	}

	draw.Draw(dst.(*image.RGBA), r, img, image.Point{}, draw.Src)

	out, err := os.Create(dstFilename)
	if err != nil {
		return err
	}

	err = bmp.Encode(out, dst)
	if err != nil {
		return err
	}

	return nil
}
