package service

import (
	"github.com/stretchr/testify/require"
	"golang.org/x/image/bmp"
	"image"
	"internshipApplicationTemplate/pkg/db"
	"internshipApplicationTemplate/pkg/db/charta"
	"os"
	"testing"
	"time"
)

type testChartaService struct {
	*ChartaService
	DB db.Charta
}

func newTestService() *testChartaService {
	db := charta.NewCharta()
	return &testChartaService{DB: db, ChartaService: NewChartaService(db)}
}

func TestChartaService_AddCharta(t *testing.T) {
	c := newTestService()

	_, err := c.AddCharta(500, 500)
	require.NoError(t, err)

	f, _ := os.Open("bmp_" + time.Now().Format(time.RFC3339) + ".bmp")
	defer os.Remove(f.Name())
	require.NoError(t, err)
	created, _ := bmp.Decode(f)

	err = os.Chdir("./expected")
	require.NoError(t, err)
	f, err = os.Open("add_charta.bmp")
	require.NoError(t, err)
	expected, err := bmp.Decode(f)
	require.NoError(t, err)

	res := compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	os.Chdir("../")
}

func TestChartaService_DeleteCharta(t *testing.T) {
	c := newTestService()

	id, err := c.AddCharta(500, 500)
	require.NoError(t, err)

	err = c.DeleteCharta(id)
	require.NoError(t, err)
}

func TestChartaService_AddPartCharta_GetPartCharta(t *testing.T) {
	c := newTestService()

	err := os.Chdir("./expected")
	require.NoError(t, err)

	// ожидаемые фрагменты
	full, err := os.Open("full.bmp")
	upperLeft, _ := os.Open("upper-left.bmp")
	centerLeft, _ := os.Open("center-left.bmp")
	lowerLeft, _ := os.Open("lower-left.bmp")
	upperRight, _ := os.Open("upper-right.bmp")
	centerRight, _ := os.Open("center-right.bmp")
	lowerRight, _ := os.Open("lower-right.bmp")
	require.NoError(t, err)

	os.Chdir("../")

	id, err := c.AddCharta(500, 500)
	require.NoError(t, err)

	// Добавляю восстановленный фрагмент, который полностью покрывает начальное изображение
	expected, _ := bmp.Decode(full)

	err = c.AddPartCharta(0, 0, 500, 500, expected, id)
	require.NoError(t, err)

	// FULL
	buf, err := c.GetPartCharta(0, 0, 500, 500, id)
	require.NoError(t, err)
	created, err := bmp.Decode(buf)
	require.NoError(t, err)
	res := compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// UPPER LEFT
	buf, err = c.GetPartCharta(-100, -100, 200, 200, id)
	require.NoError(t, err)
	created, err = bmp.Decode(buf)
	require.NoError(t, err)
	expected, _ = bmp.Decode(upperLeft)
	res = compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// CENTER LEFT
	buf, err = c.GetPartCharta(-100, 250, 250, 50, id)
	require.NoError(t, err)
	created, err = bmp.Decode(buf)
	require.NoError(t, err)
	expected, _ = bmp.Decode(centerLeft)
	res = compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// LOWER LEFT
	buf, err = c.GetPartCharta(-100, 400, 200, 50, id)
	require.NoError(t, err)
	created, err = bmp.Decode(buf)
	require.NoError(t, err)
	expected, _ = bmp.Decode(lowerLeft)
	res = compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// UPPER RIGHT
	buf, err = c.GetPartCharta(400, -100, 150, 200, id)
	require.NoError(t, err)
	created, err = bmp.Decode(buf)
	require.NoError(t, err)
	expected, _ = bmp.Decode(upperRight)
	res = compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// CENTER RIGHT
	buf, err = c.GetPartCharta(400, 200, 200, 50, id)
	require.NoError(t, err)
	created, err = bmp.Decode(buf)
	require.NoError(t, err)
	expected, _ = bmp.Decode(centerRight)
	res = compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// LOWER LEFT
	buf, err = c.GetPartCharta(450, 450, 100, 100, id)
	require.NoError(t, err)
	created, err = bmp.Decode(buf)
	require.NoError(t, err)
	expected, _ = bmp.Decode(lowerRight)
	res = compare(expected.(*image.RGBA), created.(*image.RGBA))
	require.Equal(t, true, res)

	// удаляю изображение после всех тестов
	f, err := os.Open("bmp_" + time.Now().Format(time.RFC3339) + ".bmp")
	os.Remove(f.Name())
}

// compare сравнивает два изображения
func compare(expected, actual *image.RGBA) bool {
	if expected.Bounds() != actual.Bounds() {
		return false
	}

	accumError := int32(0)

	for i := 0; i < len(expected.Pix); i++ {
		accumError += int32(sqDiffUInt8(expected.Pix[i], actual.Pix[i]))
	}

	if accumError != 0 {
		return false
	} else {
		return true
	}
}

func sqDiffUInt8(x, y uint8) uint64 {
	d := uint64(x) - uint64(y)
	return d * d
}
