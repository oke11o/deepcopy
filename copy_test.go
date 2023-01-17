package deepcopy

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestCopy(t *testing.T) {
	t.Run("not pointer", func(t *testing.T) {
		source := uint32(1)
		res := Copy(source)
		assert.Equal(t, uint32(1), res)
	})
	t.Run("nil pointer", func(t *testing.T) {
		var source *uint32
		res := Copy(source)
		resV, ok := res.(*uint32)
		assert.True(t, ok)
		assert.Nil(t, resV)
	})
	t.Run("pointer", func(t *testing.T) {
		source := uint32(1)
		sourceP := &source
		assert.Equal(t, uint32(1), *sourceP)
		res := Copy(sourceP)
		*sourceP = 3
		resP, ok := res.(*uint32)
		assert.True(t, ok)
		assert.Equal(t, uint32(1), *resP)
		assert.Equal(t, uint32(3), *sourceP)
	})
	t.Run("pointer to time", func(t *testing.T) {
		source, err := time.Parse(time.RFC3339, "2023-01-01T13:18:00Z")
		assert.NoError(t, err)
		sourceP := &source

		result := Copy(sourceP)

		source2, err := time.Parse(time.RFC3339, "2023-01-02T13:18:00Z")
		*sourceP = source2

		resultV, ok := result.(*time.Time)
		assert.True(t, ok)
		assert.Equal(t, "2023-01-01T13:18:00Z", (*resultV).Format(time.RFC3339))
	})
}
