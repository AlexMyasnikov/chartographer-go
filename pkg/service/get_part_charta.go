package service

import (
	"bytes"
	"github.com/google/uuid"
	"golang.org/x/image/bmp"
	"golang.org/x/image/draw"
	"image"
	"os"
)

func (c *ChartaService) GetPartCharta(x, y, width, height int, id uuid.UUID) (*bytes.Buffer, error) {
	filename, err := c.DB.GetChartaName(id)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	img, err := bmp.Decode(f)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	rect := image.Rectangle{
		Min: image.Point{X: 0, Y: 0},
		Max: image.Point{X: width, Y: height},
	}
	dst := image.NewRGBA64(rect)

	// Данная конструкция позволяет правильно расположить изображение,
	// в случае, когда координаты или размеры запрашиваемого фрагмента находятся за границами
	// исходного изображения
	dr := image.Rectangle{}
	if x < 0 && y < 0 {
		dr = image.Rectangle{
			Min: image.Point{X: -x, Y: -y},
			Max: image.Point{X: width, Y: height},
		}

	} else if x < 0 {
		dr = image.Rectangle{
			Min: image.Point{X: -x, Y: 0},
			Max: image.Point{X: width, Y: height},
		}
	} else if y < 0 {
		dr = image.Rectangle{
			Min: image.Point{X: 0, Y: -y},
			Max: image.Point{X: width, Y: height},
		}
	} else {
		dr = image.Rectangle{
			Min: image.Point{X: 0, Y: 0},
			Max: image.Point{X: width, Y: height},
		}
	}

	r := image.Rectangle{
		Min: image.Point{X: x, Y: y},
		Max: image.Point{X: x + width, Y: y + height},
	}
	src := img.(*image.RGBA).SubImage(r)

	// Записываю вырезанный фрагмент в буфер, так как
	// по какой-то причине без этого фрагмент не вставляется в dst
	buffer := new(bytes.Buffer)
	err = bmp.Encode(buffer, src)
	if err != nil {
		return &bytes.Buffer{}, err
	}
	src, err = bmp.Decode(buffer)
	if err != nil {
		return &bytes.Buffer{}, err
	}

	draw.Draw(dst, dr, src.(image.Image), image.Point{}, draw.Src)

	buffer = new(bytes.Buffer)
	if err = bmp.Encode(buffer, dst); err != nil {
		return &bytes.Buffer{}, err
	}

	return buffer, nil
}
