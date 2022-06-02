package charta

import (
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"internshipApplicationTemplate/pkg/db"
	"internshipApplicationTemplate/pkg/models"
	"testing"
)

func TestCharta_AddCharta_DeleteCharta_GetFilename(t *testing.T) {
	c := NewCharta()

	expected := &models.Charta{
		Id:   uuid.New(),
		Name: "test",
	}

	// ADD CHARTA
	c.AddCharta(expected)
	_, ok := c.db[expected.Id]
	require.Equal(t, true, ok)

	// ADD GET CHARTA NAME
	actual, err := c.GetChartaName(expected.Id)
	require.NoError(t, err)
	require.Equal(t, expected.Name, actual)

	// DELETE CHARTA
	c.DeleteCharta(expected.Id)
	actual, ok = c.db[expected.Id]
	require.Equal(t, false, ok)
	require.Equal(t, "", actual)

	// CHECK IF EXISTS
	actual, err = c.GetChartaName(expected.Id)
	require.Equal(t, db.ErrNotFound, err)
	require.Equal(t, "", actual)
}
